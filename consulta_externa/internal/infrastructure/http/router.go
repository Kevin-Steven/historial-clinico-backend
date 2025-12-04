package http

import "github.com/gin-gonic/gin"

func SetupRouter(handler *ConsultaExternaHandler) *gin.Engine {
	r := gin.Default()
	handler.RegisterRoutes(r)
	return r
}
