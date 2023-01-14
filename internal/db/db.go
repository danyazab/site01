package db

import (
	"site01/internal/models"
)

type DbUser map[string]models.User
type Db map[string]DbUser

var (
	DbU = make(map[string]models.User)

	DbS Db = make(map[string]DbUser)
)
