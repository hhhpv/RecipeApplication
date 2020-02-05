package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"

	Dbase "../database"
	"go.mongodb.org/mongo-driver/bson"
)

type Signin struct {
	Username string
	Password string
}

type SignInResponse struct {
	Name   string `json:"name"`
	Token  string `json:"token"`
	Result string `json:"result"`
}

const (
	MySigningKey = "someAuthenticationString"
)

func SignMeInHandler(r *mux.Router) {
	r.HandleFunc("/signmein", SignMeInHandlerFunc)
}

func SignMeInHandlerFunc(w http.ResponseWriter, r *http.Request) {
	collection := Dbase.Client.Database("recipe").Collection("users")
	w.Header().Set("Content-Type", "application/json")
	var regFields RegFields
	regerr := json.NewDecoder(r.Body).Decode(&regFields)
	if regerr != nil {
		fmt.Println(r.Body)
		fmt.Println(regerr)
		http.Error(w, regerr.Error(), http.StatusBadRequest)
		return
	}
	name := regFields.Username
	password := regFields.Password

	filter := bson.D{{"username", name}}

	var result Signin

	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		fmt.Println(err)
	}
	verifyUser := comparePasswords(result.Password, []byte(password))
	if verifyUser {
		createdToken, terr := GenerateToken([]byte(MySigningKey), name)
		if terr != nil {
			fmt.Println("Creating token failed")
		}
		if ParseToken(createdToken, MySigningKey) {
			resp := SignInResponse{
				Name: name, Token: createdToken, Result: "success",
			}
			json.NewEncoder(w).Encode(resp)
		} else {
			resp := SignInResponse{
				Name: name, Token: "", Result: "failure",
			}
			json.NewEncoder(w).Encode(resp)
		}
	} else {
		fmt.Println("Invalid credentials")
		resp := SignInResponse{
			Name: name, Token: "", Result: "failure",
		}
		json.NewEncoder(w).Encode(resp)
	}
	defer r.Body.Close()
}

func hashAndSalt(pwd []byte, cost int) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, cost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}
func comparePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}
func GenerateToken(mySigningKey []byte, username string) (string, error) {
	// Create the token
	token := jwt.New(jwt.SigningMethodHS256)
	// Set some claims
	claims := make(jwt.MapClaims)
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	token.Claims = claims
	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString(mySigningKey)
	return tokenString, err
}

func ParseToken(myToken string, myKey string) bool {
	token, err := jwt.Parse(myToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(myKey), nil
	})
	if err == nil && token.Valid {
		return true
	} else {
		return false
	}
}
