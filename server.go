package main

import (
	"api-echo/routes"
	"api-echo/services"
	"net/http"

	"api-echo/database/connections/db"

	_ "api-echo/docs/swagger"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	//db connection
	conn := getDBConnection(e)
	defer conn.Close()
	db.Conn.LogMode(true)

	services.SwaggerInfo()
	services.Validation()
	routes.Main(e)
	s := &http.Server{
		Addr: ":" + services.GetConfiguration("server.port"),
	}
	e.Logger.Fatal(e.StartServer(s))
}

func getDBConnection(e *echo.Echo) *gorm.DB {
	conn, err := db.Connect(e)
	if err != nil {
		panic(err.Error())
	}
	return conn
}
