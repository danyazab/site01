//package main
//
//import (
//	"net/http"
//
//	"github.com/labstack/echo/v4"
//)
//
//func main() {
//	e := echo.New()
//	Meps := map[string]interface{}{}
//
//	e.GET("/", func(c echo.Context) error {
//		return c.String(http.StatusOK, "хелов ворд")
//	})
//	e.Logger.Fatal(e.Start(":8080"))
//
//	url := "/key/:name"
//
//	e.POST(url, func(c echo.Context) error {
//		return c.JSON(http.StatusOK, Meps)
//	})
//
//	e.POST(url, func(c echo.Context) error {
//		ss := struct {
//			Value interface{} `json:"value"`
//		}{}
//		err := c.Bind(&url)
//		if err != nil {
//			return nil
//		}
//		Meps[c.Param("name")] = ss.Value
//
//		return c.NoContent(http.StatusCreated)
//	})
//	e.GET(url, func(c echo.Context) error {
//		res, pp := Meps[c.Param("name")]
//
//		if !pp {
//			return c.String(http.StatusNotFound, "the kay not found")
//
//		}
//		return c.JSON(http.StatusOK, res)
//	})
//
//}

package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
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

type (
	user struct {
		KEY  interface{} `json:"key"`
		Name interface{} `json:"name"`
	}
)

var (
	users = map[interface{}]*user{}
)

func createUser(c echo.Context) error {
	per := c.QueryParam("key")
	u := &user{
		KEY: per,
	}

	if err := c.Bind(u); err != nil {
		return err
	}
	users[c.Param("key")] = u

	return c.JSON(http.StatusCreated, u)
}

func getUser(c echo.Context) error {
	res, err := users[c.Param("key")]
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
	users[c.Param("key")].Name = u.Name
	return c.JSON(http.StatusOK, u.Name)
}

func deleteUser(c echo.Context) error {
	key, _ := strconv.Atoi(c.Param("key"))
	delete(users, key)
	return c.NoContent(http.StatusNoContent)
}

//func allUsers(c echo.Context) error {

//return c.JSON(http.StatusOK, &aif)
//}

func createDb(c echo.Context) error {

	//c.Param(("Db") := new(map[interface{}]*user)

	return c.JSON(http.StatusCreated, "")
}
func getDb(c echo.Context) error {

	return c.JSON(http.StatusOK, "")

}
func listDb(c echo.Context) error {

	return c.JSON(http.StatusOK, "")

}
func deleteDb(c echo.Context) error {

	return c.NoContent(http.StatusNoContent)

}
