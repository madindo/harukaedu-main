package controllers
import (
	"net/http"
	"fmt"
)

func HomeIndex(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "HomeResponse")
}

func HomeLoaderIo(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "HomeResponseloaderio-e8b5a8357a3427f049705ec89394e548")
}