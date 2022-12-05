package models

type (
	User struct {
		KEY        string
		Name       string `json:"name"`
		TimeCreate interface{}
	}
	Base struct {
		DbName string
		Vmist  interface{}
	}
)
