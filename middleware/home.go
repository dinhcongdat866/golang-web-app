package middleware

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//Need buttons for forwarding
func Home(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]

	fmt.Fprintf(w, "Hello %s!", username)
}
