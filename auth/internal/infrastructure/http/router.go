package http

import "github.com/gin-gonic/gin"

func SetupRouter(handler *AuthHandler) *gin.Engine {
	r := gin.Default()
	handler.RegisterRoutes(r)
	return r
}
