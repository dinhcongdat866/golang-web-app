package main

import (
	m "golang-web-app/middleware"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

//newly added

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/home/{username}", m.Home)
	r.HandleFunc("/login/{username}/{pwd}", m.Login)
	r.HandleFunc("/profile/{username}", m.User)
	r.HandleFunc("/chat", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "websockets.html")
	})
	r.HandleFunc("/echo", m.Echo)

	http.ListenAndServe(":80", r)
}
