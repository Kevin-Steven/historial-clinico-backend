package http

import (
	"net/http"

	"historial-clinico-backend/consulta_externa/internal/application"

	"github.com/gin-gonic/gin"
)

type ConsultaExternaHandler struct {
	crearUC *application.CrearConsultaExternaUseCase
}

func NewConsultaExternaHandler(crearUC *application.CrearConsultaExternaUseCase) *ConsultaExternaHandler {
	return &ConsultaExternaHandler{crearUC: crearUC}
}

func (h *ConsultaExternaHandler) RegisterRoutes(r *gin.Engine) {
	grp := r.Group("/consultas-externas")
	{
		grp.POST("", h.CrearConsulta)
	}
}

func (h *ConsultaExternaHandler) CrearConsulta(c *gin.Context) {
	var req application.CrearConsultaExternaRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "payload invalido"})
		return
	}

	resp, err := h.crearUC.Execute(req)
	if err != nil {
		if err == application.ErrDatosConsultaInvalidos {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error interno"})
		return
	}

	c.JSON(http.StatusCreated, resp)
}
