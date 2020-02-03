package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	Dbase "../database"
	"github.com/gorilla/mux"
)

type NewRecipeRequest struct {
	Name              string `json:"username"`
	Token             string `json:"token"`
	RecipeName        string `json:"recipename"`
	RecipeDescription string `json:"recipedescription"`
	RecipeSteps       string `json:"recipesteps"`
}

type NewRecipeInsert struct {
	Name              string `json:"username"`
	Date              string `json:"created"`
	RecipeName        string `json:"recipename"`
	RecipeDescription string `json:"recipedescription"`
	RecipeSteps       string `json:"recipesteps"`
	UID               string `json:"uid"`
}

type RecipeCreatedResponse struct {
	Name   string `json:"username"`
	Token  string `json:"token"`
	Status string `json:"status`
}

func NewRecipeHandler(r *mux.Router) {
	r.HandleFunc("/new_recipe", CheckSecurity(NewRecipeHandlerFunc))
}

func NewRecipeHandlerFunc(w http.ResponseWriter, r *http.Request) {
	var req NewRecipeRequest
	collection := Dbase.Client.Database("recipe").Collection("recipe_database")
	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var currentDate string
	currentTime := time.Now()
	currentDate = currentTime.Format("01-02-2006")
	newRecipe := NewRecipeInsert{
		Name:              req.Name,
		Date:              currentDate,
		RecipeName:        req.RecipeName,
		RecipeDescription: req.RecipeDescription,
		RecipeSteps:       req.RecipeSteps,
		UID:               req.Name + currentDate + req.RecipeName,
	}
	_, insertErr := collection.InsertOne(context.TODO(), newRecipe)
	if insertErr != nil {
		log.Fatal(insertErr)
		resp := RecipeCreatedResponse{
			Name:   req.Name,
			Token:  req.Token,
			Status: "failure",
		}
		json.NewEncoder(w).Encode(resp)
	} else {
		resp := RecipeCreatedResponse{
			Name:   req.Name,
			Token:  req.Token,
			Status: "success",
		}
		json.NewEncoder(w).Encode(resp)
	}
}
