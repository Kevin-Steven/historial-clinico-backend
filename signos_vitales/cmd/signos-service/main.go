package main

import (
	"log"
	"os"

	"historial-clinico-backend/signos_vitales/internal/application"
	"historial-clinico-backend/signos_vitales/internal/config"
	httptransport "historial-clinico-backend/signos_vitales/internal/infrastructure/http"
	"historial-clinico-backend/signos_vitales/internal/infrastructure/repository"
)

func main() {
	cfg := config.Load()

	db, err := repository.NewMySQLDB(cfg)
	if err != nil {
		log.Fatalf("error inicializando base de datos signos_vitales: %v", err)
	}

	// repositorios
	signosRepo := repository.NewMySQLSignoVitalRepository(db)
	atencionesRepo := repository.NewMySQLAtencionRepository(db)

	// casos de uso
	registrarUC := application.NewRegistrarSignosUseCase(signosRepo, atencionesRepo)
	listarUC := application.NewListarSignosUseCase(signosRepo, atencionesRepo)

	handler := httptransport.NewSignosHandler(registrarUC, listarUC)
	router := httptransport.SetupRouter(handler, cfg.JWTSecret)

	addr := ":" + cfg.HTTPPort
	log.Println("Servicio de Signos Vitales escuchando en", addr)
	if err := router.Run(addr); err != nil {
		log.Println("error al iniciar servidor HTTP de signos_vitales:", err)
		os.Exit(1)
	}
}
