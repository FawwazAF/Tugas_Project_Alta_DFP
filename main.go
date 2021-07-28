package main

import (
	"alta/project/config"
	"alta/project/routes"
	"fmt"

	"github.com/labstack/echo" //dont use v4
)

func main() {

	e := echo.New()
	config.InitDb()
	config.InitPort()
	routes.New(e)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.HTTP_PORT)))
	//new
}
