package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func main() {
	e := echo.New()

	e.POST("/my-db/:dbName/:key", createUser)
	e.GET("/my-db/:dbName/:key", getUser)
	e.PUT("/my-db/:dbName/:key", updateUser)
	e.DELETE("/my-db/:dbName/:key", deleteUser)

	e.POST("/my-db/:dbName", createDb)
	e.GET("/my-db/:dbName", getDb)
	e.GET("/my-db/", listDb)
	e.DELETE("/my-db/:dbName", deleteDb)

	e.Logger.Fatal(e.Start(":8080"))

}

type dbUser map[string]*user
type db map[string]*base

var (
	DbU dbUser = make(map[string]*user)

	DbS db = make(map[string]*base)
)

type (
	user struct {
		KEY        string
		Name       string `json:"name"`
		TimeCreate interface{}
	}
	base struct {
		dbName string
		Vmist  interface{}
	}
)

func createUser(c echo.Context) error {
	for {
		per := c.QueryParam("key")
		per1 := c.QueryParam("dbName")

		u := &user{
			KEY:        per,
			TimeCreate: time.Now(),
		}

		if err := c.Bind(u); err != nil {
			return err
		}
		_, ok := DbU[c.Param("key")]
		if ok {
			return c.String(http.StatusBadRequest, "It is not possible to create an existing one")
			continue
		}
		DbU[c.Param("key")] = u

		s := DbU
		ss := &base{
			dbName: per1,
			Vmist:  s,
		}
		DbS[c.Param("dbName")] = ss

		return c.JSON(http.StatusCreated, u)

	}
}

func getUser(c echo.Context) error {
	res, ok := DbU[c.Param("key")]
	if !ok {
		return c.String(http.StatusNotFound, "There is no such key")
	}
	return c.JSON(http.StatusOK, res)
}

func updateUser(c echo.Context) error {
	u := new(user)
	if err := c.Bind(u); err != nil {
		return err
	}
	DbU[c.Param("key")].Name = u.Name
	return c.JSON(http.StatusOK, u.Name)
}

func deleteUser(c echo.Context) error {
	delete(DbU, c.Param("key"))
	return c.NoContent(http.StatusNoContent)
}

func createDb(c echo.Context) error {
	per := c.QueryParam("dbName")

	u := &base{
		dbName: per,
		//Vmist:  user{},
	}
	if err := c.Bind(u); err != nil {
		return err
	}

	DbS[c.Param("dbName")] = u

	return c.JSON(http.StatusCreated, u)
}

func getDb(c echo.Context) error {

	res, err := DbS[c.Param("dbName")]
	if !err {
		return c.String(http.StatusNotFound, "There is no such key")
	}
	return c.JSON(http.StatusOK, res)
}

func listDb(c echo.Context) error {
	return c.JSON(http.StatusOK, DbS)
}

func deleteDb(c echo.Context) error {
	delete(DbS, c.Param("dbName"))
	return c.NoContent(http.StatusNoContent)
}
