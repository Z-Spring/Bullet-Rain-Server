package main

import (
	"BulletRain_server/database"
	"BulletRain_server/netmanager"
)

func main() {
	database.Connect()
	netmanager.NetManager(8080)
}
