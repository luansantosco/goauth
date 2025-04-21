package main

import (
	"goauth/controllers"
	"goauth/db"
	"log"
	"net/http"
)

func main() {
	db.ConnectDB()

	http.HandleFunc("/register", controllers.RegisterHandler)

	log.Println("Servidor rodando na porta 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}