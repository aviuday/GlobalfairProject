package routes 

import (
	"github.com/gorilla/mux"
	"github.com/aviuday/Globalfair-project/pkg/controllers"
)

var RegisterRoutes = func(router *mux.Router){
	router.HandleFunc("/user", controllers.UserDetails).Methods("GET")
	router.HandleFunc("/user/create", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/user/{userId}/details", controllers.GetUserById).Methods("GET")
	router.HandleFunc("/user/{userId}/update", controllers.UpdateUser).Methods("POST")
	router.HandleFunc("/user/{userId}/delete", controllers.DeleteUser).Methods("DELETE")
	// router.HandleFunc("/transaction/purchase", controllers.Purchase).Methods("POST")
	// router.HandleFunc("/transaction/payment", controllers.Payment).Methods("POST")
	// router.HandleFunc("/transaction/latest", controllers.Purchase).Methods("GET")
}