package main

import (
	"log"
	"os"

	"historial-clinico-backend/evolucion/internal/application"
	"historial-clinico-backend/evolucion/internal/config"
	httptransport "historial-clinico-backend/evolucion/internal/infrastructure/http"
	"historial-clinico-backend/evolucion/internal/infrastructure/repository"
)

func main() {
	cfg := config.Load()

	db, err := repository.NewMySQLDB(cfg)
	if err != nil {
		log.Fatalf("error inicializando base de datos evolucion: %v", err)
	}

	repo := repository.NewMySQLEvolucionRepository(db)
	crearUC := application.NewCrearEvolucionUseCase(repo)
	listarUC := application.NewListarEvolucionesUseCase(repo)

	handler := httptransport.NewEvolucionHandler(crearUC, listarUC)
	router := httptransport.SetupRouter(handler, cfg.JWTSecret)

	addr := ":" + cfg.HTTPPort
	log.Println("Servicio de Evolucion escuchando en", addr)
	if err := router.Run(addr); err != nil {
		log.Println("error al iniciar servidor HTTP de evolucion:", err)
		os.Exit(1)
	}
}
