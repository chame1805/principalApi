package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"github.com/joho/godotenv"

	reserva "principalApi/src/infraestructure"
)

func main() {
	// Cargar el archivo .env
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error al cargar el archivo env: %v", err)
	}

	// Crear una instancia de Gin
	r := gin.Default()

	// Configuración de CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://localhost:3000"}, // Agrega el puerto correcto
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))
	
	// Configuración de rutas de reservas
	reserva.SetReservaHospitalRoutes(r)

	// Ejecutar el servidor en el puerto 8080
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Error al ejecutar el servidor: %v", err)
	}
}
