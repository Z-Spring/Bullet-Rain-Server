package netmanager

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"reflect"
	"time"

	"BulletRain_server/database"
	"BulletRain_server/databaseModel"
	"BulletRain_server/msg"
	"BulletRain_server/receiveMsgHandler"
)

var (
	clients = make(map[net.Conn]*receiveMsgHandler.ClientInfo)
	methods = make(map[string]reflect.Value)
	m       reflect.Value
)

const pingInterval = 30 * time.Second
const DEFAULT_SIZE = 1024

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	// todo
	handler := &receiveMsgHandler.MsgHandler{}
	typ := reflect.TypeOf(handler)
	val := reflect.ValueOf(handler)
	for i := 0; i < typ.NumMethod(); i++ {
		method := typ.Method(i)
		m = val.Method(i)

		methods[method.Name] = m
	}
}

func NetManager(port int) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Println(err)
		return
	}

	defer listener.Close()
	fmt.Println("[服务器]启动成功")

	newConnection := make(chan net.Conn, 10)

	go func(listener2 net.Listener) {
		for {
			client, err := listener.Accept()
			if err != nil {
				log.Println("accept error: " + err.Error())
				continue
			}
			log.Printf("accept client: %s\n", client.RemoteAddr().String())
			newConnection <- client
		}
	}(listener)

	for {
		select {
		case client := <-newConnection:
			go handleConnection(client)
		case <-time.After(time.Second):
			CheckResultAndPing()
		}

	}
}

func handleConnection(client net.Conn) {
	clientInfo := &receiveMsgHandler.ClientInfo{
		Conn: client,
		ByteArray: msg.ByteArray{
			Bytes:      make([]byte, DEFAULT_SIZE),
			Capacity:   DEFAULT_SIZE,
			InitSize:   DEFAULT_SIZE,
			ReadIndex:  0,
			WriteIndex: 0,
		},
		LastPingTime: time.Now(),
	}
	clients[client] = clientInfo
	for {
		err := Receive(client, clientInfo)
		if err != nil {
			break
		}
	}

}

func Receive(client net.Conn, clientInfo *receiveMsgHandler.ClientInfo) error {
	receive := clientInfo.Bytes
	readBuff := &clientInfo.ByteArray

	if readBuff.Remain() <= 0 {
		OnReceiveData(clientInfo)
		readBuff.MoveBytes()
	}
	// 扩展之后还不够
	if readBuff.Remain() <= 0 {
		CloseConnection(clientInfo)
		log.Fatal("receive fail, maybe msg length > buff capacity")
	}
	count, err := client.Read(receive[readBuff.WriteIndex:readBuff.Remain()])
	// todo: 可能有问题？
	if err != nil {
		log.Printf("client %s is closed", client.RemoteAddr().String())
		//receiveMsgHandler.MsgLeaveBattle(clientInfo, &msgProto.MsgLeaveBattle{})
		go ReadAndUpdateRecordFromLocalFile(clientInfo.Player)
		delete(clients, client)
		if receiveMsgHandler.RM.Rooms != nil {
			delete(receiveMsgHandler.RM.Rooms, clientInfo.Player.RoomId)
		}
		receiveMsgHandler.Pm.RemovePlayer(clientInfo.Player.Id)
		return err
	}

	// close client socket
	// todo
	if count <= 0 {
		CloseConnection(clientInfo)
		delete(clients, client)
		//log.Println("Socket CloseConnection")
		return errors.New("socket CloseConnection")
	}

	readBuff.WriteIndex += count
	OnReceiveData(clientInfo)
	readBuff.CheckAndMoveBytes()
	return nil
}

func OnReceiveData(clientInfo *receiveMsgHandler.ClientInfo) {
	for {
		readBuff := &clientInfo.ByteArray
		if readBuff.Length() < 2 {
			return
		}
		bodyLength, err := readBuff.ReadInt16()
		if err != nil {
			log.Fatal(err)
		}

		if readBuff.Length() < int(bodyLength) {
			return
		}

		nameCount := 0
		//msgBase := msg.MsgBase{}
		protoName := msg.DecodeName(readBuff.Bytes, readBuff.ReadIndex, &nameCount)
		if protoName == "" {
			log.Fatal("decode name fail")
		}
		readBuff.ReadIndex += nameCount

		bodyCount := int(bodyLength) - nameCount
		base := msg.Decode(protoName, readBuff.Bytes, readBuff.ReadIndex, bodyCount)
		readBuff.ReadIndex += bodyCount

		readBuff.CheckAndMoveBytes()

		if method, ok := methods[protoName]; ok {
			clientInfo := reflect.ValueOf(clientInfo)
			//base := msg.ConvertMsgBase(base)
			base := reflect.ValueOf(base)
			args := []reflect.Value{
				clientInfo,
				base,
			}
			method.Call(args)
		}

		if readBuff.Length() <= 2 {
			//OnReceiveData(clientInfo)
			break
		}
	}
}

func OnDisConnect(info *receiveMsgHandler.ClientInfo) {
	fmt.Println("CloseConnection")

}

func CheckResultAndPing() {
	receiveMsgHandler.RM.UpdateGameTime()
	CheckPing()
	receiveMsgHandler.RM.UpdateBattleResult()
}

func CheckPing() {
	for _, info := range clients {
		if time.Since(info.LastPingTime) > pingInterval*4 {
			log.Printf("Ping CloseConnection: %s\n", info.Conn.RemoteAddr().String())
			CloseConnection(info)
			return
		}
	}
}

// CloseConnection todo:
func CloseConnection(info *receiveMsgHandler.ClientInfo) {
	OnDisConnect(info)
	info.Conn.Close()
	delete(clients, info.Conn)
}

func ReadAndUpdateRecordFromLocalFile(player *receiveMsgHandler.Player) {
	if player == nil {
		return
	}
	filePath := fmt.Sprintf("C:\\ProgramData/battleResult_%s.json", player.Id)
	file, err := os.Open(filePath)
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()

	buff := make([]byte, 1024)
	readCount, err := file.Read(buff)
	if err != nil {
		return
	}

	playerDate := &databaseModel.PlayerData{}
	json.Unmarshal(buff[:readCount], playerDate)

	if database.UpdatePlayerData(player.Id, playerDate) {
		log.Println("update player info success")
	} else {
		log.Println("update player info fail")
	}
}
