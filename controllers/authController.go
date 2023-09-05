package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/matheusferreira165/jwt-authentication/database"
	"github.com/matheusferreira165/jwt-authentication/models"
	"golang.org/x/crypto/bcrypt"
)

const SecretKey = "secret"

func Register(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()

	data, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var user models.User
	json.Unmarshal(data, &user)
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	user.Password = string(hashedPassword)

	userJson, _ := json.Marshal(user)

	database.DB.Create(&user)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(userJson)

}

func Login(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()

	data, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var userRequest models.User
	json.Unmarshal(data, &userRequest)

	var userResponse models.User

	database.DB.Where("email = ?", userRequest.Email).First(&userResponse)

	if userResponse.Id == 0 {

		notFoundMessage := struct {
			Message string `json:"message"`
		}{
			Message: "user not found",
		}

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(notFoundMessage)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userResponse.Password), []byte(userRequest.Password)); err != nil {
		incorrectPassMessage := struct {
			Message string `json:"message"`
		}{
			Message: "incorrect password",
		}

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(incorrectPassMessage)
		return

	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(userResponse.Id)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := claims.SignedString([]byte(SecretKey))
	if err != nil {
		incorrectLogin := struct {
			Message string `json:"message"`
		}{
			Message: "could not login",
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(incorrectLogin)
	}

	cookie := http.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)

	sucess := struct {
		Message string `json:"message"`
	}{
		Message: "sucess",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(sucess)
}

func User(w http.ResponseWriter, r *http.Request) {

	cookie, _ := r.Cookie("jwt")

	token, err := jwt.ParseWithClaims(cookie.Value, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {

		statusUnauthorized := struct {
			Message string `json:"message"`
		}{
			Message: "unauthentication",
		}
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(statusUnauthorized)

		return
	}

	claims := token.Claims

	json.NewEncoder(w).Encode(claims)

}
