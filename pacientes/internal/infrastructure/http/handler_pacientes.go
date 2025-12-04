package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"historial-clinico-backend/pacientes/internal/application"
)

type PacientesHandler struct {
	registrarUC *application.RegistrarPacienteUseCase
}

func NewPacientesHandler(registrarUC *application.RegistrarPacienteUseCase) *PacientesHandler {
	return &PacientesHandler{registrarUC: registrarUC}
}

func (h *PacientesHandler) RegisterRoutes(r *gin.Engine) {
	grp := r.Group("/pacientes")
	{
		grp.POST("", h.RegistrarPaciente)
	}
}

func (h *PacientesHandler) RegistrarPaciente(c *gin.Context) {
	var req application.RegistrarPacienteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "payload invalido"})
		return
	}

	resp, err := h.registrarUC.Execute(req)
	if err != nil {
		switch err {
		case application.ErrDatosObligatorios, application.ErrFechaNacimientoInvalida, application.ErrPacienteDuplicado:
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error interno"})
		}
		return
	}

	c.JSON(http.StatusCreated, resp)
}
