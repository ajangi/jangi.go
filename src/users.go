package src

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

type User struct {
	gorm.Model
	Name string `json:"name"`
}
func GetPersonalInfo (c echo.Context) (err error) {

}
