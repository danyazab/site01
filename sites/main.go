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
