package middleware

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"golang-web-app/utils"
)

//Check if account requested exist in database
func User(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]

	userExist := utils.SearchUser(username)

	//homeURL := "localhost:80/home"
	if userExist {
		http.Redirect(w, r, "localhost:80/home/dat", http.StatusFound)
	} else {
		fmt.Fprintf(w, "Account does not exist")
	}
}
