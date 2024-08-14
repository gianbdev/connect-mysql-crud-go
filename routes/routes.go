package routes

import (
	"go-api-crud/controllers"
	"go-api-crud/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/register", controllers.Register).Methods("POST")
	router.HandleFunc("/login", controllers.Login).Methods("POST")
	router.Handle("/profile", middlewares.AuthMiddleware(http.HandlerFunc(controllers.GetProfile))).Methods("GET")
}
