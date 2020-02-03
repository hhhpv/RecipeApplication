package handlers

import (
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

func AboutViewHandler(r *mux.Router) {
	r.HandleFunc("/user_profile", CheckSecurity(AboutViewHandlerFunc))
}
func AboutViewHandlerFunc(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len(""):]
	p := loadPage(title)
	t, _ := template.ParseFiles("./views/about.html")
	t.Execute(w, p)
}
