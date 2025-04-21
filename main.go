package main

import (
	"fmt"
	"goauth/db"
	"log"
	"net/http"
)

func main() {
	db.ConnectDB()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "GoAuth rodando! 🚀")
	})

	log.Println("Servidor rodando na porta 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}