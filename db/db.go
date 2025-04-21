package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var Db *sql.DB

func ConnectDB() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Arquivo .env não encontrado. Lendo variáveis diretamente do ambiente...")
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	if host == "" || port == "" || user == "" || password == "" || dbname == "" {
		log.Fatal("Variáveis de ambiente para conexão com banco de dados não configuradas corretamente.")
	}

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	Db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Erro ao abrir conexão com banco:", err)
	}

	err = Db.Ping()
	if err != nil {
		log.Fatal("Erro ao conectar no banco de dados:", err)
	}

	log.Println("✅ Banco de dados conectado com sucesso!")
}