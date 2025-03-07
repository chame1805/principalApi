package main

import (
	"log"
	

	reserva "principalApi/src/infraestructure"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error al cargar el archivo env: %v", err)
	}


	r:= gin.Default()

	


	reserva.SetReservaHospitalRoutes(r)
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Error al ejecutar el servidor: %v", err)
	}
}