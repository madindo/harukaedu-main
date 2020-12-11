package router

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	maincontroller "github.com/madindo/harukaedu-main/controllers"
	usercontroller "github.com/madindo/harukaedu-users/controllers"
)

func Web() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", maincontroller.HomeIndex ).Methods("GET")

	myRouter.HandleFunc("/users", usercontroller.UserIndex).Methods("GET")
	myRouter.HandleFunc("/users", usercontroller.UserStore).Methods("POST")
	myRouter.HandleFunc("/users/{id}", usercontroller.UserUpdate).Methods("PUT")
	myRouter.HandleFunc("/users/{id}", usercontroller.UserDelete).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", myRouter))
}
