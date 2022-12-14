package db

import (
	"site01/internal/models"
)

type dbUser map[string]models.User
type db map[string]dbUser

var (
	DbU = make(map[string]models.User)

	DbS db = make(map[string]dbUser)
)
