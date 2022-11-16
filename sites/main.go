package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// Routes
	e.POST("/name", createUser)
	e.GET("/name/:key", getUser)
	//e.GET("/name", allUsers)
	e.PUT("/name/:key", updateUser)
	e.DELETE("/name/:key", deleteUser)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

type (
	user struct {
		KEY  int    `json:"key"`
		Name string `json:"name"`
	}
)

var (
	users = map[int]*user{}
	pr    = 1
)

func createUser(c echo.Context) error {
	u := &user{
		KEY: pr,
	}

	if err := c.Bind(u); err != nil {
		return err
	}
	users[u.KEY] = u
	pr++
	return c.JSON(http.StatusCreated, u)
}

func getUser(c echo.Context) error {
	key, _ := strconv.Atoi(c.Param("key"))
	return c.JSON(http.StatusOK, users[key])
}

func updateUser(c echo.Context) error {
	u := new(user)
	if err := c.Bind(u); err != nil {
		return err
	}
	key, _ := strconv.Atoi(c.Param("key"))
	users[key].Name = u.Name
	return c.JSON(http.StatusOK, users[key])
}

func deleteUser(c echo.Context) error {
	key, _ := strconv.Atoi(c.Param("key"))
	delete(users, key)
	return c.NoContent(http.StatusNoContent)
}

//func allUsers(c echo.Context) error {

//return c.JSON(http.StatusOK, &aif)
//}
