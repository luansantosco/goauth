package db

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var Db *sql.DB

func ConnectDB() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Arquivo .env não encontrado. Lendo variáveis diretamente...")
	}

	databaseUrl := os.Getenv("DATABASE_URL")
	if databaseUrl == "" {
		log.Fatal("DATABASE_URL não configurada.")
	}

	Db, err = sql.Open("postgres", databaseUrl)
	if err != nil {
		log.Fatal("Erro ao abrir conexão com banco:", err)
	}

	err = Db.Ping()
	if err != nil {
		log.Fatal("Erro ao conectar no banco de dados:", err)
	}

	log.Println("✅ Banco de dados conectado com sucesso!")
}