package main

import (
	"net/http"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	conn, err := gorm.Open(postgres.Open(""), &gorm.Config{})

	if err != nil {
		panic("could not connect to the database")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	http.ListenAndServe(":3000", nil)
}
