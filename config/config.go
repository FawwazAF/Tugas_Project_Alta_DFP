package config

import (
	"alta/project/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var HTTP_PORT int

func InitDb() {

	//Set connection string here, use mysql username password and schema at your pc
	// envVar := "root:Minus12345@tcp(localhost:3306)/new_schema?charset=utf8&parseTime=True&loc=Local"
	// connectionString := os.Getenv(envVar)
	connectionString := "root:Minus12345@tcp(localhost:3306)/new_schema?charset=utf8&parseTime=True&loc=Local"

	// connectionString := "root:Minus12345@tcp(localhost:3306)/new_schema?charset=utf8&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	InitMigrate()
}

func InitPort() {
	// var err error

	HTTP_PORT = 8080 //Port Setting

	// HTTP_PORT, err = strconv.Atoi(os.Getenv("HTTP_PORT"))
	// if err != nil {
	// 	panic(err)
	// }
}

func InitMigrate() {
	DB.AutoMigrate(&model.User{})
}
