package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"text/template"
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

const (
	MySigningKey = "someAuthenticationString"
)

func SignMeInHandler(r *mux.Router) {
	r.HandleFunc("/signmein", SignMeInHandlerFunc)
}

func SignMeInHandlerFunc(w http.ResponseWriter, r *http.Request) {

	//For JSON request processing
	// var b signin
	// decoder := json.NewDecoder(r.Body)
	// err := decoder.Decode(&b)
	// if err != nil {
	// 	fmt.Println("error")
	// 	fmt.Println(err)
	// }
	// fmt.Println((b))
	// u := r.Form.Get("username")
	//body, _ := ioutil.ReadAll(r.Body)
	// u := r.Form.Get("username")

	collection := Dbase.Client.Database("recipe").Collection("users")
	err := r.ParseForm()
	if err != nil {
		fmt.Println("cannot parse")
	}
	name := r.PostFormValue("username")
	password := r.PostFormValue("passcode")
	//passwordLength := len(password)
	fmt.Println("username is ", name)
	fmt.Println("password is ", password)
	//hashedPassword := hashAndSalt([]byte(password), passwordLength)

	// filter := bson.M{"username": name, "password": hashedPassword}
	filter := bson.D{{"username", name}}
	// bson.D{"username", name, "password", hashedPassword}

	var result Signin
	fmt.Println(filter)
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		fmt.Println(err)
		fmt.Println("there is an error")
	}
	verifyUser := comparePasswords(result.Password, []byte(password))
	if verifyUser {
		fmt.Println("Valid credentials")
		fmt.Println(result)
		createdToken, err := GenerateToken([]byte(MySigningKey), name)
		if err != nil {
			fmt.Println("Creating token failed")
		}
		ParseToken(createdToken, MySigningKey)
	} else {
		fmt.Println("Invalid credentials")
		fmt.Println(result)
	}
	defer r.Body.Close()
	title := r.URL.Path[len(""):]
	p := loadPage(title)
	t, _ := template.ParseFiles("./views/HomePage.html")
	t.Execute(w, p)
}

func hashAndSalt(pwd []byte, cost int) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, cost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}
func comparePasswords(hashedPwd string, plainPwd []byte) bool {
	// Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice
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
		fmt.Println("Your token is valid.  I like your style.")
		fmt.Println(token)
		return true
	} else {
		fmt.Println("This token is terrible!  I cannot accept this.")
		return false
	}
}
