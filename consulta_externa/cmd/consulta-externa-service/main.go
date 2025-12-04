package main

import (
	"log"
	"os"

	"historial-clinico-backend/consulta_externa/internal/application"
	"historial-clinico-backend/consulta_externa/internal/config"
	httptransport "historial-clinico-backend/consulta_externa/internal/infrastructure/http"
	"historial-clinico-backend/consulta_externa/internal/infrastructure/repository"
)

func main() {
	cfg := config.Load()

	db, err := repository.NewMySQLDB(cfg)
	if err != nil {
		log.Fatalf("error inicializando base de datos consulta_externa: %v", err)
	}

	repo := repository.NewMySQLConsultaExternaRepository(db)
	crearUC := application.NewCrearConsultaExternaUseCase(repo)

	handler := httptransport.NewConsultaExternaHandler(crearUC)
	router := httptransport.SetupRouter(handler)

	addr := ":" + cfg.HTTPPort
	log.Println("Servicio de Consulta Externa escuchando en", addr)
	if err := router.Run(addr); err != nil {
		log.Println("error al iniciar servidor HTTP de consulta externa:", err)
		os.Exit(1)
	}
}
