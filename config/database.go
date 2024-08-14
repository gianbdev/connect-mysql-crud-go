package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql" // Asegúrate de importar el driver MySQL
)

var DB *sql.DB

func ConnectDB() {
	var err error
	// Ajusta la cadena de conexión, omitiendo la contraseña
	DB, err = sql.Open("mysql", "user:@tcp(localhost:3306)/go-api-db")
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	// Verifica la conexión
	if err = DB.Ping(); err != nil {
		log.Fatalf("Error pinging the database: %v", err)
	}
	log.Println("Successfully connected to the database")
}
