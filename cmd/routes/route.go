package routes

import (
	"github.com/bigthugboy/food-network/pkg/controller"
	"github.com/gorilla/mux"
)

func foodNearbyRoute(r *mux.Router) {

	r.HandleFunc("/", controller.Welcome).Methods("GET")
	r.HandleFunc("/register", controller.RegisterHandler).Methods("POST")
	r.HandleFunc("/login", controller.LoginHandler).Methods("GET")
	r.HandleFunc("/get-foodNearby", controller.GetFoodNearbyHandler).Methods("GET")
	r.HandleFunc("/select-restaurant", controller.SelectRestaurantHandler).Methods("GET")
}
