package auth

import (
	"fmt"
	"github.com/zentooling/graide/internal/logger"
	"net/http"
	"time"
)

var log = logger.New("auth")
var jwtKey = []byte("my_secret_key") // TODO make legit

// func Login(user string, password string, session http.ResponseWriter)

func Login(w http.ResponseWriter, r *http.Request) {

	// get user and password from POST body
	err := r.ParseForm()
	if err != nil {
		log.Println("parse form", err)
		// appropriate message?
		w.WriteHeader(http.StatusPartialContent)
		return
	}
	// TODO user/password must be passed over https
	user := r.FormValue("user_name")
	password := r.FormValue("password")

	// for now assume lookup is good

	if err := userLookup(user, password); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Print(w, "not authorized")
		return
	}

	var cookie = http.Cookie{
		Name:    "session_token",
		Value:   "valid",
		Expires: time.Now().Add(120 * time.Second),
	}

	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/index", http.StatusFound)

}

func userLookup(user, password string) error {
	_ = user
	_ = password
	return nil
}

func IsAuthorized(role string, r http.Request) bool {

	var ret = true

	// TODO add role check
	cookie, err := r.Cookie("session_token")
	if err != nil || cookie == nil {
		ret = false
	}

	return ret
}
