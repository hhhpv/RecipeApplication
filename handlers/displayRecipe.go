package handlers

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"text/template"

	Dbase "../database"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

type Unique struct {
	UID string `json:"uid"`
}

type DisplayRecipe struct {
	Name              string `json:"username"`
	Date              string `json:"created"`
	RecipeName        string `json:"recipename"`
	RecipeDescription string `json:"recipedescription"`
	RecipeSteps       string `json:"recipesteps"`
	UID               string `json:"uid"`
}

type PageDisplayRecipe struct {
	Title             string
	Body              []byte
	Name              string
	Date              string
	RecipeName        string
	RecipeDescription string
	RecipeSteps       string
}

func loadPageDisplayRecipe(title string, user string, recipename string, recipedesrciption string, recipesteps string, createdat string) *PageDisplayRecipe {
	filename := title + ".txt"
	body, _ := ioutil.ReadFile(filename)
	return &PageDisplayRecipe{Title: title, Body: body, Name: user, RecipeDescription: recipedesrciption, Date: createdat, RecipeName: recipename, RecipeSteps: recipesteps}
}
func DisplayRecipeHandler(r *mux.Router) {
	r.HandleFunc("/display_recipe/{uid}", CheckSecurity(DisplayRecipeHandlerFunc))
}
func DisplayRecipeHandlerFunc(w http.ResponseWriter, r *http.Request) {
	var result DisplayRecipe
	collection := Dbase.Client.Database("recipe").Collection("recipe_database")
	params := mux.Vars(r)
	uid := params["uid"]
	fmt.Println(uid)
	filter := bson.D{{"uid", uid}}
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result.RecipeSteps)
	title := r.URL.Path[len(""):]
	p := loadPageDisplayRecipe(title, result.Name, result.RecipeName, result.RecipeDescription, result.RecipeSteps, result.Date)
	t, _ := template.ParseFiles("./views/displayRecipe.html")
	t.Execute(w, p)
}
