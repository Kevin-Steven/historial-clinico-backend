package http

import "github.com/gin-gonic/gin"

type SignosRouterConfig struct {
	Handler   *SignosHandler
	JWTSecret string
}

func SetupRouter(handler *SignosHandler, jwtSecret string) *gin.Engine {
	r := gin.Default()

	auth := JWTAuthMiddleware(jwtSecret)

	group := r.Group("/atenciones/:id")
	group.Use(auth)
	{
		group.POST("/signos-vitales", handler.Registrar)
		group.GET("/signos-vitales", handler.ListarPorAtencion)
	}

	return r
}
