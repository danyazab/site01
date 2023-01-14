package internal

import (
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"site01/internal/db"
	"site01/internal/models"
	"site01/internal/resp"

	"sync"
	"time"
)

var lock = sync.RWMutex{}

// CreateUser godoc
// @Tags User
// @Description Get User
// @Accept json
// @Summary Create user in DB
// @Produce json
// @Param   enter_dbName path    string           true "Enter which database"
// @Param   rep_type   path    string             true "Enter which key to create"
// @Param   input  body      models.User true  "CreateUser JSON"
// @Success  201    {object}  resp.JSONSuccessResult{data=models.User,code=int,message=string}
// @Router  /my-db/{dbName}/{key} [post]
// @Failure	 404 	{object} resp.JSONIntServerErrReqResult{code=int,message=string}
// @Failure  400    {object} resp.JSONBadReqResult{code=int,message=string}
func CreateUser(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()
	u := models.User{
		TimeCreate: time.Now(),
	}
	_, ok := db.DbS[c.Param("dbName")]
	if !ok {
		return resp.FailResponse(c, http.StatusNotFound, "Please creat first db")
	}

	if err := c.Bind(&u); err != nil {
		return resp.FailResponse(c, http.StatusBadRequest, err.Error())
	}
	if len(u.Name) > 25 {
		return errors.New("too long")
	}
	if len(u.Name) < 2 {
		return errors.New("too minimal")
	}

	db.DbU[c.Param("key")] = u
	db.DbS[c.Param("dbName")][c.Param("key")] = u
	//docs.SwaggerInfo

	return resp.SuccessResponse(c, http.StatusCreated, u)
}

// GetUser godoc
// @Tags User
// @Description Get User
// @Accept json
// @Summary Get user in DB
// @Produce json
// @Param   dbName 		path    string               true "Write out from which database to extract"
// @Param   key  	path    string               true "Enter which key to extract"
// @Success   200   {object}  resp.JSONSuccessResult{data=models.User,code=int,message=string}
// @Router  /my-db/{dbName}/{key} 	[get]
// @Failure	  404 	{object} resp.JSONIntServerErrReqResult{code=int,message=string}
func GetUser(c echo.Context) error {
	res, ok := db.DbU[c.Param("key")]
	if !ok {
		return resp.FailResponse(c, http.StatusNotFound, "There is no such key")
	}
	return resp.SuccessResponse(c, http.StatusOK, res)
}

// UpdateUser godoc
// @Tags User
// @Description Update User
// @Update User
// @Accept json
// @Summary Update user in DB
// @Produce json
// @Param   dbName path    string               true "Enter the changes in which database"(path)
// @Param   key path    string               true "Commit the changes to the database" Format(path)
// @Param        input  body      models.User  true  "UpdateUser JSON"
// @Success      200   {object}  resp.JSONSuccessResult{data=models.User,code=int,message=string}
// @Router  /my-db/{dbName}/{key} [put]
// @Failure	 404 	{object} resp.JSONIntServerErrReqResult{code=int,message=string}
func UpdateUser(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()
	u := models.User{
		TimeCreate: time.Now(),
	}
	if err := c.Bind(&u); err != nil {
		return err
	}
	db.DbS[c.Param("dbName")][c.Param("key")] = u
	return resp.SuccessResponse(c, http.StatusOK, u.Name)
}

// DeleteUser godoc
// @Tags User
// @Description Delete User
// @DB Delete User
// @Accept json
// @Summary Delete User in DB
// @Produce json
// @Param   dbName path    string               true "Write out from which database to extract"
// @Param   key path    string               true "Enter which key to extract"
// /// @Success    204   {object}	 resp.JSONSuccessResult{code=int,message=string}
// @Router  /my-db/{dbName}/{key} [delete]
// @Failure	 404 	{object} resp.JSONIntServerErrReqResult{code=int,message=string}
func DeleteUser(c echo.Context) error {
	delete(db.DbS[c.Param("dbName")], c.Param("key"))
	return resp.SuccessResponse(c, http.StatusNoContent, "")
}

// CreateDb godoc
// @Tags DB
// @Description Create DB
// @DB Create DB
// @Accept json
// @Summary Create DB
// @Produce json
// @Param   dbName	 path      string           true "Enter in which database to create"
// @Success      201 {object} resp.JSONSuccessResult{code=int,message=string}
// @Router  /my-db/{dbName} [post]
// @Failure	 404 	{object} resp.JSONIntServerErrReqResult{code=int,message=string}
// @Failure 400 {object}  resp.JSONBadReqResult{code=int,message=string}
func CreateDb(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()
	mm, ok := db.DbS[c.Param("dbName")]
	if !ok {
		mm = make(map[string]models.User)
		db.DbS[c.Param("dbName")] = mm
	} else {
		resp.FailResponse(c, http.StatusBadRequest, "Please change name of db")
	}

	return resp.SuccessResponse(c, http.StatusCreated, "")
}

// GetDb godoc
// @Tags DB
// @Description Get DB
// @DB Get DB
// @Accept json
// @Summary Get user in DB
// @Produce json
// @Param   dbName 	path    string               true "Write out from which database to extract"
// @Success      200   {object}  resp.JSONSuccessResult{data=db.DbUser,code=int,message=string}
// @Router  /my-db/{dbName} [get]
// @Failure	 404 	{object} resp.JSONIntServerErrReqResult{code=int,message=string}
func GetDb(c echo.Context) error {
	mm, err := db.DbS[c.Param("dbName")]
	if !err {
		return resp.FailResponse(c, http.StatusNotFound, "There is no such key")
	}
	return resp.SuccessResponse(c, http.StatusOK, mm)
}

// ListDb godoc
// @Tags my-db
// @Description Get All list
// @List All
// @Accept json
// @Summary Get List with all DB and all User
// @Produce json
// @Success   200   {object}  resp.JSONSuccessResult{data=db.Db,code=int,message=string}
// @Router  /my-db/ [get]
func ListDb(c echo.Context) error {
	return resp.SuccessResponse(c, http.StatusOK, db.DbS)
}

// DeleteDb godoc
// @Tags DB
// @Description Delete DB
// @DB Delete
// @Accept json
// @Summary Delete DB
// @Produce json
// @Param   dbName path    string          true "Enter which base to delete"
// ////@Success    204  {object} resp.JSONSuccessResult{code=int,message=string}
// @Router  /my-db/{dbName} [delete]
func DeleteDb(c echo.Context) error {
	delete(db.DbS, c.Param("dbName"))
	return resp.SuccessResponse(c, http.StatusNoContent, "")
}
