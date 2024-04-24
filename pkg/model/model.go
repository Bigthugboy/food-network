package model

import (
	_ "encoding/json"

	middeware "github.com/bigthugboy/food-network/cmd/middleware"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Customer struct {
	gorm.Model
	FirstName   string `gorm: ""json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	Password    string `json:"password"`
}

func init() {
	middeware.Connect()
	db = middeware.GetDB()
	db.AutoMigrate(&Customer{})
}
