package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	Dbase "../database"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

type DeleteRecipe struct {
	Name  string `json:"username"`
	Token string `json:"token"`
	UID   string `json:"uid"`
}
type DeleteResult struct {
	Name   string `json:"username"`
	Token  string `json:"token"`
	Status string `json:"status"`
}

func DeleteRecipeHandler(r *mux.Router) {
	r.HandleFunc("/delete_recipe", CheckSecurity(DeleteRecipeHandlerFunc))
}
func DeleteRecipeHandlerFunc(w http.ResponseWriter, r *http.Request) {
	collection := Dbase.Client.Database("recipe").Collection("recipe_database")
	var req DeleteRecipe
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	filter := bson.M{"name": req.Name, "uid": req.UID}
	_, merr := collection.DeleteOne(context.TODO(), filter)
	if merr != nil {
		result := DeleteResult{
			Name:   req.Name,
			Token:  req.Token,
			Status: "failure",
		}
		json.NewEncoder(w).Encode(result)
	} else {
		result := DeleteResult{
			Name:   req.Name,
			Token:  req.Token,
			Status: "success",
		}
		json.NewEncoder(w).Encode(result)
	}
}
