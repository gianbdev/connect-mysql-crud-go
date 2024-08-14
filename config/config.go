package config

import (
	"os"
)

var (
	JwtSecret = os.Getenv("JWT_SECRET")
	DbDSN     = os.Getenv("DB_DSN") // Ejemplo: "user:password@tcp(localhost:3306)/dbname"
)
