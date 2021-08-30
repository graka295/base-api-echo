package db

import (
	"fmt"

	"api-echo/services"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo/v4"
)

var Conn *gorm.DB

// GetConnection set connection string
func GetConnection() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		services.GetConfiguration("connection.db.DB_USER"),
		services.GetConfiguration("connection.db.DB_PASSWORD"),
		services.GetConfiguration("connection.db.DB_HOST"),
		services.GetConfiguration("connection.db.DB_PORT"),
		services.GetConfiguration("connection.db.DB_NAME"),
	)
}

// Connect function for get connection db
func Connect(e *echo.Echo) (db *gorm.DB, err error) {
	db, err = gorm.Open("mysql", GetConnection())
	Conn = db
	return db, err
}
