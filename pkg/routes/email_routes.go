package routes

import (
	"github.com/gorilla/mux"
	"github.com/niharikabhavaraju/go_scheduler/pkg/controllers"
)

var RegisterEmailRoutes = func(router *mux.Router) {
	router.HandleFunc("/email", controllers.CreateEmail).Methods("POST")
	router.HandleFunc("/email", controllers.GetEmails).Methods("GET")
	router.HandleFunc("/email/{id}", controllers.GetEmailById).Methods("GET")
	router.HandleFunc("/email/{id}", controllers.DeleteEmail).Methods("DELETE")
	router.HandleFunc("/email/{id}", controllers.UpdateEmail).Methods("PUT")
}
