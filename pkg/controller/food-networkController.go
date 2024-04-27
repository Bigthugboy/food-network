package controller

import (
	"encoding/json"

	"log"
	"net/http"

	_ "github.com/bigthugboy/food-network/db"
	"github.com/bigthugboy/food-network/db/query"
	"github.com/bigthugboy/food-network/db/repo"
	"github.com/bigthugboy/food-network/pkg/config"
	"github.com/bigthugboy/food-network/pkg/model"
	"github.com/go-playground/validator"

	"github.com/jinzhu/gorm"
)

type foodNearby struct {
	app *config.AppTools
	DB  query.DBstore
}

func newFoodNetwork(app *config.AppTools, db *gorm.DB) db.mainstore {
	return &foodNearby{
		app: app,
		DB:  repo.NewFoodNetworkDB(app, db),
	}
}

func (f *foodNearby) Welcome() {
	return

}

func (f *foodNearby) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var customer model.Customer
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		log.Println("Error parsing form:", err)
		return
	}
	customer.Password, _ = config.Encrypt(customer.Password)
	// Validation
	if err := f.app.Validate.Struct(&customer); err != nil {
		if _, ok := err.(*validator.InvalidValidationError); !ok {
			http.Error(w, "Bad request", http.StatusBadRequest)
			log.Println("Validation error:", err)
			return
		}
	}
	track, err := f.DB.InsertCustomer(customer)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		log.Println("Error adding user to database:", err)
		return
	}
	switch track {
	case 1:
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	case 0:
		response := map[string]string{"message": "Registered Successfully"}

		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			log.Println("Error encoding JSON:", err)
			return
		}
	}
}

// login function
func (f *foodNearby) LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		log.Println("Error parsing form:", err)
		return
	}
	email := r.Form.Get("email")
	password := r.Form.Get("password")
	if !config.IsEmailValid(email) {
		http.Error(w, "Invalid email format", http.StatusBadRequest)
		log.Println("Invalid email format")
		return
	}
	_, customerPassword, err := f.DB.GetCustomerByEmail(email)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		log.Println("Error retrieving user from database:", err)
		return
	}
	if _, err := config.Verify(password, customerPassword); err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		log.Println("Invalid email or password:", err)
		return
	}
	return
}

// get all resturant
func (f *foodNearby) GetFoodNearbyHandler(w http.ResponseWriter, r *http.Request) {
	data, err := f.DB.GetAllResturants()
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		log.Println("Error retrieving restaurants from database:", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if len(data) >= 1 {
		json.NewEncoder(w).Encode(map[string]interface{}{"restaurants": data})
	} else {
		json.NewEncoder(w).Encode(map[string]string{"message": "No restaurants available"})
	}
}

// SelectRestaurantHandler handles requests to search for restaurants by name, menu list, or location
func (f *foodNearby) SelectRestaurantHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	query := r.URL.Query()
	name := query.Get("name")
	menuList := query["menu"]
	location := query.Get("location")
	matchedRestaurants := f.searchRestaurants(name, menuList, location)
	if len(matchedRestaurants) > 0 {
		json.NewEncoder(w).Encode(matchedRestaurants)
	} else {
		json.NewEncoder(w).Encode(map[string]string{"message": "No restaurants found matching the criteria"})
	}
}
