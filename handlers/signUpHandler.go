package handlers

import (
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

func SignUpHandler(r *mux.Router) {
	r.HandleFunc("/signup", SignUpHandlerFunc)
}

func SignUpHandlerFunc(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len(""):]
	p := loadPage(title)
	t, _ := template.ParseFiles("./views/signup.html")
	t.Execute(w, p)
}
