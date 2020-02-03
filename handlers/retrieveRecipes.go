package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"

	Dbase "../database"
	"github.com/gorilla/mux"
)

type RequestRecipe struct {
	Name  string `json:"username`
	Token string `json:"token"`
}
type GetRecipe struct {
	Name              string `json:"username"`
	Date              string `json:"created"`
	RecipeName        string `json:"recipename"`
	RecipeDescription string `json:"recipedescription"`
	RecipeSteps       string `json:"recipesteps"`
	UID               string `json:"uid"`
}

type ReturnRecipes struct {
	Name    string      `json:"username"`
	Recipes []GetRecipe `json:"recipes"`
	Status  string      `json:"status"`
}

func RetrieveRecipeHandler(r *mux.Router) {
	r.HandleFunc("/get_recipes", CheckSecurity(RetrieveRecipeHandlerFunc))
}

func RetrieveRecipeHandlerFunc(w http.ResponseWriter, r *http.Request) {
	var req RequestRecipe
	collection := Dbase.Client.Database("recipe").Collection("recipe_database")
	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	recipes := []GetRecipe{}
	pipeline := ([]bson.M{bson.M{"$sample": bson.M{"size": 6}}})
	cursor, merr := collection.Aggregate(context.Background(), pipeline)
	for cursor.Next(context.Background()) {
		var temp GetRecipe
		cursor.Decode(&temp)
		recipes = append(recipes, temp)
	}
	if merr != nil {
		log.Fatal(err)
		resp := ReturnRecipes{
			Name:    req.Name,
			Recipes: []GetRecipe{},
			Status:  "failure",
		}
		json.NewEncoder(w).Encode(resp)
	} else {
		resp := ReturnRecipes{
			Name:    req.Name,
			Recipes: recipes,
			Status:  "success",
		}
		json.NewEncoder(w).Encode(resp)
	}
}
