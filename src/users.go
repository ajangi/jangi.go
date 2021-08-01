package main

import (
	"github.com/jinzhu/gorm"
	//"github.com/labstack/echo/v4"
)

type User struct {
	gorm.Model
	Name string `json:"name"`
}
