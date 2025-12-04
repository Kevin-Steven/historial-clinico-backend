package http

import (
	"net/http"

	"historial-clinico-backend/auth/internal/application"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	loginUC *application.LoginUsuarioUseCase
	crearUC *application.CrearUsuarioUseCase
}

func NewAuthHandler(loginUC *application.LoginUsuarioUseCase, crearUC *application.CrearUsuarioUseCase) *AuthHandler {
	return &AuthHandler{loginUC: loginUC, crearUC: crearUC}
}

func (h *AuthHandler) RegisterRoutes(r *gin.Engine) {
	auth := r.Group("/auth")
	{
		auth.POST("/login", h.Login)
	}

	usuarios := r.Group("/usuarios")
	{
		usuarios.POST("", h.CrearUsuario)
	}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req application.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "payload invalido"})
		return
	}

	resp, err := h.loginUC.Execute(req)
	if err != nil {
		if err == application.ErrCredencialesInvalidas {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error interno"})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *AuthHandler) CrearUsuario(c *gin.Context) {
	var req application.CrearUsuarioRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "payload invalido"})
		return
	}

	resp, err := h.crearUC.Execute(req)
	if err != nil {
		if err == application.ErrDatosUsuarioInvalidos {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error interno"})
		return
	}

	c.JSON(http.StatusCreated, resp)
}
