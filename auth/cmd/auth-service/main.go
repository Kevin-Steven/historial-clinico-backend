package main

import (
	"log"
	"os"

	"historial-clinico-backend/auth/internal/application"
	"historial-clinico-backend/auth/internal/config"
	httptransport "historial-clinico-backend/auth/internal/infrastructure/http"
	"historial-clinico-backend/auth/internal/infrastructure/repository"
)

func main() {
	cfg := config.Load()

	db, err := repository.NewMySQLDB(cfg)
	if err != nil {
		log.Fatalf("error inicializando base de datos auth: %v", err)
	}

	usuarioRepo := repository.NewMySQLUsuarioRepository(db)
	tokenSvc := application.NewJWTService(cfg.JWTSecret)
	loginUC := application.NewLoginUsuarioUseCase(usuarioRepo, tokenSvc)
	crearUC := application.NewCrearUsuarioUseCase(usuarioRepo)

	handler := httptransport.NewAuthHandler(loginUC, crearUC)
	router := httptransport.SetupRouter(handler)

	addr := ":" + cfg.HTTPPort
	log.Println("Servicio de Auth escuchando en", addr)
	if err := router.Run(addr); err != nil {
		log.Println("error al iniciar servidor HTTP auth:", err)
		os.Exit(1)
	}
}
