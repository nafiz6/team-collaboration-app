package accountsHandler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"teams/middleware/db"
	. "teams/models"

	"teams/middleware/cors"

	"github.com/gorilla/sessions"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var store = sessions.NewCookieStore([]byte("t0p-s3cr3t"))

func init() {
}

func Login(w http.ResponseWriter, r *http.Request) {
	cors.EnableCorsCredentials(&w)

	decoder := json.NewDecoder(r.Body)
	var t UserLogin
	err := decoder.Decode(&t)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error Decording JSON: " + err.Error()))
		log.Println("Error Decording JSON")
		log.Println(err.Error())
		return
	}
	username := t.Username
	password := t.Password

	session, err := store.Get(r, "session-key")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error Fetching Session: " + err.Error()))
		log.Println("Error in session")
		log.Println(err.Error())
		return
	}

	log.Println("logging in username " + username)

	// TODO : get this from db where the hashed version will be stored

	var userDetails UserDetailsNew

	err = db.Users.FindOne(context.TODO(), bson.D{{"username", t.Username}}).Decode(&userDetails)

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Username doesn't exit!"))
		log.Println("USER NOT FOUND")
		return
	}

	storedPassword := userDetails.Password

	err = bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(password))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Incorrect Password!"))
		return
	}

	// get clients cookies
	session, _ = store.Get(r, "session-key")
	log.Println("logged in with id " + userDetails.ID.Hex())
	session.Values["id"] = userDetails.ID.Hex() // TODO : store an uid
	session.Save(r, w)

	log.Println("logged in")
	fmt.Fprintln(w, "logged in")

}

func Register(w http.ResponseWriter, r *http.Request) {
	cors.EnableCorsCredentials(&w)
	decoder := json.NewDecoder(r.Body)
	var userDetails UserRegistration
	err := decoder.Decode(&userDetails)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error Decording JSON: " + err.Error()))
		log.Println("Error Decording JSON")
		log.Println(err.Error())
		return
	}
	username := userDetails.Username
	password := userDetails.Password
	// TODO OTHER FORM FIELDS???

	// hash the password
	hash := getHashedPassword(password)

	log.Println("registration username " + username)
	if hash == nil {
		return
	}
	/*
		TODO :	TODONE
		store hash and username to db
		check unique username
	*/

	userDetails.ID = primitive.NewObjectID()

	validationErr := validateRegistrationInput(userDetails)
	if validationErr != "" {
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte(validationErr))
		return
	}

	userDetails.Password = string(hash[:])
	if userDetails.Dp == "" {
		userDetails.Dp = "http://localhost:8080/static/default.jpg"
	}

	var userFound UserDetailsNew

	findUserErr := db.Users.FindOne(context.TODO(), bson.D{{"username", username}}).Decode(&userFound)

	if findUserErr != mongo.ErrNoDocuments {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Username already exists!"))
		log.Println("USER NOT FOUND")
		return
	}

	_, insertErr := db.Users.InsertOne(context.Background(), userDetails)

	if insertErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error inserting into database"))
		return
	}
	//insertedID := insertResult.InsertedID

	// Get session cookies from client
	session, err := store.Get(r, "session-key")
	if err != nil {
		log.Println(err.Error())
	}

	// TODO : get UID from db	DONE
	// store hash and user id in clients cookies
	//session.Values["id"] = "55269"
	session.Values["id"] = userDetails.ID.Hex()

	// Save.
	session.Save(r, w)

	log.Println("Successfully reged")
	fmt.Fprint(w, "Successfully reged")

}

func Logout(w http.ResponseWriter, r *http.Request) {
	cors.EnableCorsCredentials(&w)
	session, err := store.Get(r, "session-key")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error retrieving session"))
		return
	}

	session.Values["id"] = ""
	session.Options.MaxAge = -1
	session.Save(r, w)
	fmt.Fprint(w, "Logged out")
}

func SecretPage(w http.ResponseWriter, r *http.Request) {
	cors.EnableCorsCredentials(&w)
	username := GetUserId(r)
	if username == "" {
		fmt.Fprint(w, "Not logged in")
		return
	}
	fmt.Fprint(w, "Hello "+username)

}

func GetUserId(r *http.Request) string {
	session, _ := store.Get(r, "session-key")
	untyped, ok := session.Values["id"]
	if !ok {
		log.Println(ok)
		log.Println("1")

		return ""
	}

	username, ok := untyped.(string)
	if !ok {
		log.Print(ok)
		return ""
	}

	log.Println(username)

	return username
}

func getHashedPassword(password string) []byte {
	cost := bcrypt.DefaultCost

	hash, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return nil
	}

	return hash
}

func validateRegistrationInput(userDetails UserRegistration) string {
	if len(userDetails.Name) == 0 || !HasOnlyLetters(userDetails.Name) {
		return "Invalid Name"
	}
	if len(userDetails.Username) < 6 {
		return "Username must be atleast 6 characters"
	}
	if len(userDetails.Password) < 6 {
		return "Password must be atleast 6 characters."
	}

	return ""
}

func HasOnlyLetters(s string) bool {
	for _, r := range s {
		if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') && r != ' ' {
			return false
		}
	}
	return true
}
