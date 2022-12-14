package main

import (
	"github.com/labstack/echo/v4"
	"site01/internal"
)

func main() {
	e := echo.New()

	e.POST("/my-db/:dbName/:key", internal.CreateUser)
	e.GET("/my-db/:dbName/:key", internal.GetUser)
	e.PUT("/my-db/:dbName/:key", internal.UpdateUser)
	e.DELETE("/my-db/:dbName/:key", internal.DeleteUser)

	e.POST("/my-db/:dbName", internal.CreateDb)
	e.GET("/my-db/:dbName", internal.GetDb)
	e.GET("/my-db/", internal.ListDb)
	e.DELETE("/my-db/:dbName", internal.DeleteDb)

	e.Logger.Fatal(e.Start(":8088"))

}
