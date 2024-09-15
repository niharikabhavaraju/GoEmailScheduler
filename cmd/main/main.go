package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/niharikabhavaraju/go_scheduler/pkg/routes"
	"github.com/niharikabhavaraju/go_scheduler/pkg/scheduler"
)

func main() {
	go scheduler.StartScheduler()
	router := mux.NewRouter()
	router.Use(mux.CORSMethodMiddleware(router))
	routes.RegisterEmailRoutes(router)
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8080", router))
}
