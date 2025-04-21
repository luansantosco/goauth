package configs

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func ConnectDB() *sql.DB {
	dbHost := GetEnv("DB_HOST", "localhost")
	dbPort := GetEnv("DB_PORT", "5432")
	dbUser := GetEnv("DB_USER", "goauth_user")
	dbPassword := GetEnv("DB_PASSWORD", "think")
	dbName := GetEnv("DB_NAME", "goauth")

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("❌ Error opening database connection:", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("❌ Error pinging database:", err)
	}

	log.Println("✅ Database successfully connected!")
	return db
}
