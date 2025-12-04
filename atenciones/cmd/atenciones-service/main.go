package main

import (
	"log"
	"os"

	"historial-clinico-backend/atenciones/internal/application"
	"historial-clinico-backend/atenciones/internal/config"
	httptransport "historial-clinico-backend/atenciones/internal/infrastructure/http"
	"historial-clinico-backend/atenciones/internal/infrastructure/repository"
)

func main() {
	cfg := config.Load()

	db, err := repository.NewMySQLDB(cfg)
	if err != nil {
		log.Fatalf("error inicializando base de datos atenciones: %v", err)
	}

	atencionRepo := repository.NewMySQLAtencionRepository(db)
	signosRepo := repository.NewMySQLSignoVitalRepository(db)
	crearAtencionUC := application.NewCrearAtencionUseCase(atencionRepo)
	getAtencionUC := application.NewGetAtencionUseCase(atencionRepo)
	registrarSignosUC := application.NewRegistrarSignosUseCase(signosRepo, atencionRepo)

	handler := httptransport.NewAtencionesHandler(crearAtencionUC, getAtencionUC, registrarSignosUC)
	router := httptransport.SetupRouter(handler, cfg.JWTSecret)

	addr := ":" + cfg.HTTPPort
	log.Println("Servicio de Atenciones escuchando en", addr)
	if err := router.Run(addr); err != nil {
		log.Println("error al iniciar servidor HTTP de atenciones:", err)
		os.Exit(1)
	}
}
