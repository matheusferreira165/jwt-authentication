package main

import (
	"fmt"
	"net/http"

	"github.com/matheusferreira165/jwt-authentication/database"
	"github.com/matheusferreira165/jwt-authentication/routes"
)

func main() {
	database.Connect()
	routes.Setup()
	fmt.Println("Servidor Iniciado")
	http.ListenAndServe(":3000", nil)
}
