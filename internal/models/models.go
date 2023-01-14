package models

import "time"

type (
	User struct {
		Name       string `json:"name" example:"Danya"`
		TimeCreate time.Time
	}
)
