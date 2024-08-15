package main

import (
	"log"
	"os"

	"go-api-crud/controllers"
	"go-api-crud/database"
	"go-api-crud/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Cargar variables de entorno desde el archivo .env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Obtener la cadena de conexión de la base de datos desde las variables de entorno
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		log.Fatal("DB_DSN is not set in the .env file")
	}

	// Inicializar la base de datos
	database.Connect(dsn)
	database.Migrate()

	// Inicializar el enrutador
	router := initRouter()
	router.Run(":8080")
}

func initRouter() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api/v1")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/login", controllers.GenerateToken)
			auth.POST("/register", controllers.RegisterUser)
			auth.POST("/refresh-token", controllers.RefreshToken)
		}

		secured := api.Group("").Use(middlewares.Auth())
		{
			secured.GET("/ping", controllers.Ping)
		}
	}

	return router
}
