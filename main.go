package main

import (
	"net/http"

	"github.com/matheusferreira165/jwt-authentication/database"
	"github.com/matheusferreira165/jwt-authentication/routes"
)

func main() {
	database.Connect()
	routes.Setup()

	http.ListenAndServe(":3000", nil)
}
