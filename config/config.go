package config

import (
	"alta/project/model"
	_ "alta/project/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var HTTP_PORT int

func InitDb() {

	// Framecode for environmental
	/*
		envVar := "root:Minus12345@tcp(localhost:3306)/new_schema?charset=utf8&parseTime=True&loc=Local"
		connectionString := os.Getenv(envVar)
	*/

	//Set connection string here, use mysql username password and schema at your pc
	//connectionString := "root:Minus12345@tcp(localhost:3306)/alta_shop_project?charset=utf8&parseTime=True&loc=Local"
	connectionString := "root:02021996Doni*@tcp(localhost:3306)/alta_shop_project?charset=utf8&parseTime=True&loc=Local" // doni local computer
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

	// Framecode for environmental
	/*
		HTTP_PORT, err = strconv.Atoi(os.Getenv("HTTP_PORT"))
		if err != nil {
		panic(err)
		}
	*/
}

func InitMigrate() {
	DB.AutoMigrate(&model.User{})
	DB.AutoMigrate(&model.Product{})
	DB.AutoMigrate(&model.Shopping_cart{})
}
