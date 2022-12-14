package internal

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"site01/internal/db"
	"site01/internal/models"
	"time"
)

func CreateUser(c echo.Context) error {
	for {
		u := models.User{
			TimeCreate: time.Now(),
		}
		_, ok := db.DbS[c.Param("dbName")]
		if !ok {
			return c.String(http.StatusNotFound, "Please creat first db")
			continue
		}
		if err := c.Bind(&u); err != nil {
			return c.String(http.StatusBadRequest, "bad request")
		}
		db.DbU[c.Param("key")] = u

		db.DbS[c.Param("dbName")][c.Param("key")] = u

		return c.JSON(http.StatusCreated, u)

	}
}

func GetUser(c echo.Context) error {
	res, ok := db.DbU[c.Param("key")]
	if !ok {
		return c.String(http.StatusNotFound, "There is no such key")
	}
	return c.JSON(http.StatusOK, res)
}

func UpdateUser(c echo.Context) error {
	u := models.User{
		TimeUpdate: time.Now(),
	}
	if err := c.Bind(&u); err != nil {
		return err
	}
	db.DbS[c.Param("dbName")][c.Param("key")] = u
	return c.JSON(http.StatusOK, u.Name)
}

func DeleteUser(c echo.Context) error {
	delete(db.DbS[c.Param("dbName")], c.Param("key"))
	return c.NoContent(http.StatusNoContent)
}

func CreateDb(c echo.Context) error {
	mm, ok := db.DbS[c.Param("dbName")]
	if !ok {
		mm = make(map[string]models.User)
		db.DbS[c.Param("dbName")] = mm
	}
	return c.JSON(http.StatusCreated, "Created")
}

func GetDb(c echo.Context) error {
	ss := len(db.DbS[c.Param("dbName")])
	mm, err := db.DbS[c.Param("dbName")]
	ee := fmt.Sprintf("ddd %n ff %b", ss, mm)
	if !err {
		return c.String(http.StatusNotFound, "There is no such key")
	}
	return c.JSON(http.StatusOK, ee)
}

func ListDb(c echo.Context) error {
	return c.JSON(http.StatusOK, db.DbS)
}

func DeleteDb(c echo.Context) error {
	delete(db.DbS, c.Param("dbName"))
	return c.NoContent(http.StatusNoContent)
}
