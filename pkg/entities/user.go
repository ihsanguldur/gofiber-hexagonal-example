package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Lat      string `json:"lat"`
	Lon      string `json:"lon"`
}
