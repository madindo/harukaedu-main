package router

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	homecontroller "github.com/madindo/harukaedu-main/controllers"
	usercontroller "github.com/madindo/harukaedu-users/controllers"
)

func Web() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homecontroller.HomeIndex ).Methods("GET")

	myRouter.HandleFunc("/users", usercontroller.UserIndex).Methods("GET")
	myRouter.HandleFunc("/users", usercontroller.UserStore).Methods("POST")
	myRouter.HandleFunc("/users/{id}", usercontroller.UserUpdate).Methods("PUT")
	myRouter.HandleFunc("/users/{id}", usercontroller.UserDelete).Methods("DELETE")
	myRouter.HandleFunc("/loaderio-e8b5a8357a3427f049705ec89394e548.txt", homecontroller.HomeLoaderIo).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", myRouter))
}
