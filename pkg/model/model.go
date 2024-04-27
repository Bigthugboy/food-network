package model

import (
	_ "encoding/json"
	"time"

	middeware "github.com/bigthugboy/food-network/cmd/middleware"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Customer struct {
	gorm.Model
	FirstName   string   `gorm: ""json:"firstName"`
	LastName    string   `json:"lastName"`
	Email       string   `json:"email"`
	PhoneNumber string   `json:"phoneNumber"`
	Password    string   `json:"password"`
	GeoLocation Location `josn:location`
}

type Resturant struct {
	gorm.Model
	Name     string   `gorm: "json:"name"`
	Location Location `json:location`
	MenuList []string `json:"menulist"`
}
type Location struct {
	Address string `json:"address"`
	City    string `json:"city"`
}

type PaymentSystem struct {
	ID                int64     `json:"_id"`
	PhoneNumber       string    `json:"first_name"`
	Email             string    `json:"email"`
	AmountPaid        string    `json:"amount_payed"`
	DatePaid          time.Time `json:"date_payed"`
	IsSuccessful      bool      `json:"is_successful"`
	DateExpired       time.Time `json:"date_expired"`
	GetPayment        int64     `json:"get_payment"`
	FirstName         string    `json:"first_name" Usage:"required,alpha"`
	LastName          string    `json:"last_name" Usage:"required,alpha"`
	Reference         string    `json:"reference"`
	Status            bool      `json:"status"`
	CustomerID        string    `json:"customerID"`
	AuthorizationCode string    `json:"authorizationCode"`
}
type PaymentRequest struct {
	Email       string    `json:"email"`
	Amount      string    `json:"amount"`
	SubAccount  string    `json:"subaccount"`
	Currency    string    `json:"currency"`
	FirstName   string    `json:"first_name" Usage:"required,alpha"`
	LastName    string    `json:"last_name" Usage:"required,alpha"`
	DatePayed   time.Time `json:"date_payed"`
	PhoneNumber string    `json:"phone" Usage:"required"`
}

type PaymentResponse struct {
	Status  bool         `json:"status"`
	Message string       `json:"message"`
	Data    ResponseData `json:"data"`
}

type Authorizations struct {
	AuthorizationCode string `json:"authorizationCode"`
}

type ResponseData struct {
	AuthorizationUrl string         `json:"authorization_url"`
	AccessCode       string         `json:"access_code"`
	Reference        string         `json:"reference"`
	Amount           string         `json:"amount"`
	Status           bool           `json:"status"`
	Authorization    Authorizations `json:"authorization"`
	StatusCode       string         `json:"status_code"`
}

type ValidateResponse struct {
	Status  bool         `json:"status"`
	Message string       `json:"message"`
	Data    ResponseData `json:"data"`
}

func init() {
	middeware.Connect()
	db = middeware.GetDB()
	db.AutoMigrate(&Customer{})
}
