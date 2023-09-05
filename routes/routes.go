package routes

import (
	"net/http"

	"github.com/matheusferreira165/jwt-authentication/controllers"
	"github.com/rs/cors"
)

func Setup() {

	c := cors.New(cors.Options{
		AllowCredentials: true,
	})

	http.HandleFunc("/api/v1/register", controllers.Register)
	http.HandleFunc("/api/v1/login", controllers.Login)
	http.HandleFunc("/api/v1/user", controllers.User)
	http.HandleFunc("/api/v1/logout", controllers.Logout)
	http.Handle("/", c.Handler(http.DefaultServeMux))
}
