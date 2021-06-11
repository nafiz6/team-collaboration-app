package accountsHandler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
)

var store = sessions.NewCookieStore([]byte("t0p-s3cr3t"))

func init() {
}

func Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")

	session, err := store.Get(r, "session-key")
	if err != nil {
		log.Println(err.Error())
		fmt.Fprintln(w, err.Error())
		return
	}

	fmt.Println(username)

	// TODO : get this from db where the hashed version will be stored
	hash := getHashedPassword("password")

	err = bcrypt.CompareHashAndPassword(hash, []byte(password))
	if err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}

	// get clients cookies
	session, _ = store.Get(r, "session-key")
	session.Values["id"] = "1123" // TODO : store an uid
	session.Save(r, w)

	log.Println("logged in")
	fmt.Fprintln(w, "logged in")

}

func Register(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")
	// TODO OTHER FORM FIELDS???

	// hash the password
	hash := getHashedPassword(password)

	log.Println(username)
	if hash == nil {
		return
	}
	/*
		TODO :
		store hash and username to db
		check unique username
	*/

	// Get session cookies from client
	session, err := store.Get(r, "session-key")
	if err != nil {
		log.Println(err.Error())
	}

	// TODO : get UID from db
	// store hash and user id in clients cookies
	session.Values["id"] = "55269"

	// Save.
	session.Save(r, w)

	log.Println("Successfully reged")
	fmt.Fprint(w, "Successfully reged")

}

func Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-key")
	session.Values["id"] = ""
	session.Options.MaxAge = -1
	session.Save(r, w)
	fmt.Fprint(w, "Logged out")
}

func SecretPage(w http.ResponseWriter, r *http.Request) {
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
