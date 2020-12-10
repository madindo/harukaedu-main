package controllers
import (
	"net/http"
	"fmt"
)

func HomeIndex(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "HomeResponse")
}