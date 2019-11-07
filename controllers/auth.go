package controllers

import (
	"encoding/json"
	"net/http"
	"strings"
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/ifeanyilawrence/go-task-api/config"
	"github.com/ifeanyilawrence/go-task-api/models"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

//GetToken : return an authentication token
func GetToken(w http.ResponseWriter, r *http.Request) {

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	var existingUser models.User

	_ = config.Users.Find(bson.M{"email": user.Email}).One(&existingUser)

	if existingUser.Email == "" {
		http.Error(w, http.StatusText(404), http.StatusNotFound)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password))

	if err != nil {
		http.Error(w, http.StatusText(404), http.StatusNotFound)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": existingUser.Email,
		"name":  existingUser.Name,
		"id":    existingUser.ID,
	})

	tokenString, err := token.SignedString([]byte("secret"))

	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	result := struct {
		Name  string `json:"name"`
		Email string `json:"email"`
		Token string `json:"token"`
	}{
		Name:  existingUser.Name,
		Email: existingUser.Email,
		Token: tokenString,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(result)
}

//Profile : returns user profile
func Profile(w http.ResponseWriter, r *http.Request) {

	status, token := Authenticate(r)

	if !status {
		http.Error(w, "Invalid Token", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		result := struct{ Email, Name string }{
			Email: claims["email"].(string),
			Name:  claims["name"].(string),
		}

		json.NewEncoder(w).Encode(result)
		return
	}

	http.Error(w, "Invalid Token", http.StatusBadRequest)
	return
}

//Authenticate : authenticates a user
func Authenticate(r *http.Request) (bool, *jwt.Token) {

	fmt.Println("Am here")
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		return false, nil
	}

	tokenSplit := strings.Split(tokenString, " ")[1]

	token, _ := jwt.Parse(tokenSplit, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return false, nil
		}
		return []byte("secret"), nil
	})

	fmt.Println("Am here too", token)
	if token.Valid {
		return true, token
	}

	return false, nil
}
