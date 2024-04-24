package routes

import (
	"net/http"

	"github.com/bigthugboy/food-network/pkg/controller"
	"github.com/gorilla/mux"
)

var foodNearbyRoute = func(r mux.Router) {

	r.HandleFunc("/", controller.Welcome).Methods("GET")
	r.HandleFunc("/register", controller.RegisterHandler).Methods("POST")
	r.HandleFunc("/login", controller.LoginHandler).Methods("GET")
	r.HandleFunc("/get-foodNearby", controller.getFoodNearbyHandler).Methods("GET")
	r.HandleFunc("/select-resturant", controller.SelectResturantHandler).Methods("GET")
	http.Handle("/", r)
}
