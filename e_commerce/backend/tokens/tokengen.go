package tokens

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type SignedDetails struct {
	Email      string
	First_Name string
	Last_Name  string
	Uid        string
	jwt.StandardClaims
}

func TokenGenerator(email string, firstname string, lastname string, uid string) (signedToken string, signedfreshtoken string, err error) {
	claims := &SignedDetails{
		Email:      email,
		First_Name: firstname,
		Last_Name:  lastname,
		Uid:        uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}
	refreshclaims := &SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}
	token, _ := jwt.NewWithClaims

}

func ValidateToken() {

}

func UpdateAllTokens() {

}
