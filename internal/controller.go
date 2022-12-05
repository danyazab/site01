package internal

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"site01/internal/db"
	"site01/internal/models"
	"time"
)

func CreateUser(c echo.Context) error {
	for {
		per := c.QueryParam("key")
		per1 := c.QueryParam("dbName")

		u := &models.User{
			KEY:        per,
			TimeCreate: time.Now(),
		}

		if err := c.Bind(u); err != nil {
			return err
		}
		_, ok := db.DbU[c.Param("key")]
		if ok {
			return c.String(http.StatusBadRequest, "It is not possible to create an existing one")
			continue
		}
		db.DbU[c.Param("key")] = u

		s := db.DbU
		ss := &models.Base{
			DbName: per1,
			Vmist:  s,
		}
		db.DbS[c.Param("dbName")] = ss

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
	u := new(models.User)
	if err := c.Bind(u); err != nil {
		return err
	}
	db.DbU[c.Param("key")].Name = u.Name
	return c.JSON(http.StatusOK, u.Name)
}

func DeleteUser(c echo.Context) error {
	delete(db.DbU, c.Param("key"))
	return c.NoContent(http.StatusNoContent)
}

func CreateDb(c echo.Context) error {
	per := c.QueryParam("dbName")

	u := &models.Base{
		DbName: per,
		//Vmist:  user{},
	}
	if err := c.Bind(u); err != nil {
		return err
	}

	db.DbS[c.Param("dbName")] = u

	return c.JSON(http.StatusCreated, u)
}

func GetDb(c echo.Context) error {

	res, err := db.DbS[c.Param("dbName")]
	if !err {
		return c.String(http.StatusNotFound, "There is no such key")
	}
	return c.JSON(http.StatusOK, res)
}

func ListDb(c echo.Context) error {
	return c.JSON(http.StatusOK, db.DbS)
}

func DeleteDb(c echo.Context) error {
	delete(db.DbS, c.Param("dbName"))
	return c.NoContent(http.StatusNoContent)
}
