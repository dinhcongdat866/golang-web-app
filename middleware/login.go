package middleware

import (
	"fmt"
	"golang-web-app/utils"
	"net/http"

	"github.com/gorilla/mux"
)

func Login(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]
	password := vars["pwd"]

	userExist := utils.SearchUserPwd(username, password)

	//homeURL := "localhost:80/home"
	if userExist {
		http.Redirect(w, r, "localhost:80/home/dat", http.StatusFound)
	} else {
		fmt.Fprintf(w, "Account does not exist")
	}
}
