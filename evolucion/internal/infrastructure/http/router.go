package http

import "github.com/gin-gonic/gin"

func SetupRouter(handler *EvolucionHandler, jwtSecret string) *gin.Engine {
	r := gin.Default()
	r.Use(NewAuthMiddleware(jwtSecret))
	handler.RegisterRoutes(r)
	return r
}
