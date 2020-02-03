package handlers

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

func AddRecipeHandler(r *mux.Router) {
	r.HandleFunc("/add_recipe", CheckSecurity(AddRecipeHandlerFunc))
}

func AddRecipeHandlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Cookie)
	title := r.URL.Path[len(""):]
	p := loadPage(title)
	t, _ := template.ParseFiles("./views/addRecipe.html")
	t.Execute(w, p)
}
