package database

import (
	"log"
	"regexp"

	"BulletRain_server/databaseModel"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func Connect() {
	var err error
	dbname := "root:admin@tcp(localhost:3306)/bulletrain?charset=utf8mb4&parseTime=True&loc=Local"
	Db, err = gorm.Open(mysql.Open(dbname), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
}

func IsSafeString(str string) bool {
	matched, err := regexp.MatchString(`[-|;|,|\/|\(|\)|\[|\]|\}|\{|%|@|\*|!|\'\\]`, str)
	if err != nil {
		log.Fatal(err)
	}
	return !matched
}

func IsAccountExist(id string) bool {
	if !IsSafeString(id) {
		return false
	}
	var account databaseModel.Account
	result := Db.First(&account, "id=?", id)
	return result.RowsAffected > 0
}

func Register(id, pw string) bool {
	if !IsSafeString(id) && !IsSafeString(pw) {
		return false
	}
	if IsAccountExist(id) {
		return false
	}

	account := databaseModel.Account{
		Id: id,
		Pw: pw,
	}
	result := Db.Create(&account)
	return result.RowsAffected > 0
}

func CreatePlayerData(id string) bool {
	if !IsSafeString(id) {
		return false
	}

	playerData := databaseModel.PlayerData{
		Id:   id,
		Coin: 0,
		Text: "",
		Win:  0,
		Loss: 0,
	}
	result := Db.Create(&playerData)
	return result.RowsAffected > 0
}

func CheckPassword(id, pw string) bool {
	if !IsSafeString(id) && !IsSafeString(pw) {
		return false
	}
	account := databaseModel.Account{}
	result := Db.First(&account, "id=? and pw=?", id, pw)
	return result.RowsAffected > 0
}

func GetPlayerData(id string) databaseModel.PlayerData {
	if !IsSafeString(id) {
		return databaseModel.PlayerData{}
	}

	var playerData databaseModel.PlayerData
	result := Db.First(&playerData, "id=?", id)
	if result.RowsAffected <= 0 {
		return databaseModel.PlayerData{}
	}

	return playerData
}

func UpdatePlayerData(id string, data *databaseModel.PlayerData) bool {
	if !IsSafeString(id) {
		return false
	}

	var playerData databaseModel.PlayerData
	result := Db.First(&playerData, "id=?", id)
	if result.RowsAffected <= 0 {
		log.Printf("can't find id = `%s` player data\n", id)
		return false
	}

	Db.Model(&playerData).Updates(data)
	return result.RowsAffected > 0
}
