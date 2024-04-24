package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/bigthugboy/food-network/db/query"
	"github.com/bigthugboy/food-network/db/repo"
	"github.com/bigthugboy/food-network/pkg/config"
	"github.com/bigthugboy/food-network/pkg/model"
	_ "github.com/bigthugbpoy/food-network/db/store"
	"github.com/go-playground/validator"

	"github.com/jinzhu/gorm"
)

type foodNearby struct {
	app *config.AppTools
	DB  query.DBstore
}

func newFoodNetwork(app *config.AppTools, db *gorm.DB) store.mainstore {
	return &foodNearby{
		app: app,
		DB:  repo.NewFoodNetworkDB(app, db),
	}
}

func (f *foodNearby) Welcome() {

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
