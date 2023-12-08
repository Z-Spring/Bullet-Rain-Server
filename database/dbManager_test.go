package database

import (
	"testing"

	"BulletRain_server/databaseModel"
)

func TestIsSafeString(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "test1", args: args{str: "admin"}, want: true},
		{name: "test2", args: args{str: "admin;"}, want: false},
		{name: "test3", args: args{str: "admin%"}, want: false},
		{name: "test4", args: args{str: "admin@"}, want: false},
		{name: "test5", args: args{str: "admin!"}, want: false},
		{name: "test6", args: args{str: "admin*"}, want: false},
		{name: "test7", args: args{str: "admin'"}, want: false},
		{name: "test8", args: args{str: "admin\""}, want: true},
		{name: "test9", args: args{str: "admin\\"}, want: false},
		{name: "test10", args: args{str: "admin/"}, want: false},
		{name: "test11", args: args{str: "admin("}, want: false},
		{name: "test12", args: args{str: "admin)"}, want: false},
		{name: "test13", args: args{str: "admin["}, want: false},
		{name: "test14", args: args{str: "admin]"}, want: false},
		{name: "test15", args: args{str: "admin{"}, want: false},
		{name: "test16", args: args{str: "admin}"}, want: false},
		{name: "test17", args: args{str: "admin-"}, want: false},
		{name: "test18", args: args{str: "admin,"}, want: false},
		{name: "test19", args: args{str: "admin."}, want: true},
		{name: "test20", args: args{str: "admin?"}, want: true},
		{name: "test21", args: args{str: "admin:"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsSafeString(tt.args.str); got != tt.want {
				t.Errorf("IsSafeString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRegister(t *testing.T) {
	Connect()
	if ok := Register("admin", "admin"); ok {
		t.Log("register success")
	} else {
		t.Error("register failed")
	}
}

func TestCreatePlayerData(t *testing.T) {
	Connect()
	if ok := CreatePlayerData("admin"); ok {
		t.Log("create player data success")
	} else {
		t.Error("create player data failed")
	}
}

func TestCheckPassword(t *testing.T) {
	Connect()
	if ok := CheckPassword("1", "1"); ok {
		t.Log("check password success")
	} else {
		t.Error("check password failed")
	}
}

func TestGetPlayerData(t *testing.T) {
	Connect()
	playerData := GetPlayerData("1")
	if playerData == (databaseModel.PlayerData{}) {
		t.Error("get player data failed")
	}
	t.Log(playerData.Coin)
}

func TestUpdatePlayerInfo(t *testing.T) {
	Connect()
	playerInfo := &databaseModel.PlayerData{
		Win:  4,
		Text: "loss",
	}

	names := []string{"1", "admin", "123", "roots"}
	for _, name := range names {
		if ok := UpdatePlayerData(name, playerInfo); ok {
			t.Logf("update  %s player info success", name)
		} else {
			t.Error("update player info failed")
		}
	}

}
