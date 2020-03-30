package handlers

import (
	"crypto/rsa"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/jwtauth"
)

type Jwt struct {
	VerifyKey *rsa.PublicKey
	SignKey   *rsa.PrivateKey
}

func (j *Jwt) JwtFinalHandler(w http.ResponseWriter, r *http.Request) {
	token, _, err := jwtauth.FromContext(r.Context())
	if err != nil {
		log.Printf("error from jwt %s", err.Error())
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	if token == nil || !token.Valid {
		log.Println("invalid token")
		http.Error(w, "invalid token", http.StatusUnauthorized)
		return
	}
	fmt.Fprintf(w, "jwt is %s", "valid")
}
