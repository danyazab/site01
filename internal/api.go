package internal

import (
	"fmt"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"site01/config"
)

func RunServer(cfg *config.Database) error {
	port := 8088

	e := echo.New()
	e.POST("/my-db/:dbName/:key", CreateUser)
	e.GET("/my-db/:dbName/:key", GetUser)
	e.PUT("/my-db/:dbName/:key", UpdateUser)
	e.DELETE("/my-db/:dbName/:key", DeleteUser)

	e.POST("/my-db/:dbName", CreateDb)
	e.GET("/my-db/:dbName", GetDb)
	e.GET("/my-db/", ListDb)
	e.DELETE("/my-db/:dbName", DeleteDb)

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	return e.Start(fmt.Sprintf(":%d", port))
}
