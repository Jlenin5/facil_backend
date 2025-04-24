package repository

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // Driver de PostgreSQL
)

func ConnectDB() *sqlx.DB {
	errr := godotenv.Load()
  if errr != nil {
    log.Fatal("Error loading .env file")
  }

	// Obtener variables de entorno
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	if host == "" || port == "" || user == "" || password == "" || dbname == "" {
		log.Fatalf("Faltan variables de entorno. Verifica DB_HOST, DB_PORT, DB_USER, DB_PASSWORD y DB_NAME.")
	}

	// Crear la cadena de conexión
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Conectar a la base de datos
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatalf("No se pudo conectar a la base de datos: %v", err)
	}

	return db
}