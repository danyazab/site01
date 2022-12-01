package main

import (
	_ "database/sql"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
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
		KEY  string `json:"key"`
		Name string `json:"name"`
	}
	base struct {
		dbName string
		Vmist  interface{}
	}
)

var perR interface{}

func createUser(c echo.Context) error {
	//per1 := c.QueryParam("dbName")
	per := c.QueryParam("key")
	perR = c.QueryParam("dbName")

	//var x user
	u := &user{
		//DbName: per1,
		KEY: per,
	}

	//DbU["key"] = &user{
	//	//DbName: per1,
	//	KEY: per,
	//}

	if err := c.Bind(u); err != nil {
		return err
	}
	DbU[c.Param("key")] = u

	return c.JSON(http.StatusCreated, u)
}

func getUser(c echo.Context) error {
	res, err := DbU[c.Param("key")]
	if !err {
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
	key, _ := strconv.Atoi(c.Param("key"))
	delete(DbU, string(key))
	return c.NoContent(http.StatusNoContent)
}

func createDb(c echo.Context) error {
	//un := new(base)
	per := c.QueryParam("dbName")
	//ne := new(db)
	//gra := user{
	//	KEY:  un.KEY,
	//	Name: un.Name}

	u := &base{
		dbName: per,
		Vmist:  DbU,
	}
	//DbS["dbName"] = &base{
	//	dbName: per,
	//	Vmist:  ne,
	//}

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
	key, _ := strconv.Atoi(c.Param("dbName"))
	delete(DbS, string(key))
	return c.NoContent(http.StatusNoContent)
}
