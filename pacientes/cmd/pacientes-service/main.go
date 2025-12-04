package main

import (
	"log"
	"os"

	"historial-clinico-backend/pacientes/internal/application"
	"historial-clinico-backend/pacientes/internal/config"
	httptransport "historial-clinico-backend/pacientes/internal/infrastructure/http"
	"historial-clinico-backend/pacientes/internal/infrastructure/repository"
)

func main() {
	cfg := config.Load()

	db, err := repository.NewMySQLDB(cfg)
	if err != nil {
		log.Fatalf("error inicializando base de datos: %v", err)
	}

	pacienteRepo := repository.NewMySQLPacienteRepository(db)
	registrarUC := application.NewRegistrarPacienteUseCase(pacienteRepo)
	handler := httptransport.NewPacientesHandler(registrarUC)
	router := httptransport.SetupRouter(handler, cfg.JWTSecret)

	addr := ":" + cfg.HTTPPort
	log.Println("Servicio de Pacientes escuchando en", addr)
	if err := router.Run(addr); err != nil {
		log.Println("error al iniciar servidor HTTP:", err)
		os.Exit(1)
	}
}
