package http

import (
	"net/http"
	"strconv"

	"historial-clinico-backend/evolucion/internal/application"

	"github.com/gin-gonic/gin"
)

type EvolucionHandler struct {
	crearUC  *application.CrearEvolucionUseCase
	listarUC *application.ListarEvolucionesUseCase
}

func NewEvolucionHandler(crearUC *application.CrearEvolucionUseCase, listarUC *application.ListarEvolucionesUseCase) *EvolucionHandler {
	return &EvolucionHandler{crearUC: crearUC, listarUC: listarUC}
}

func (h *EvolucionHandler) RegisterRoutes(r *gin.Engine) {
	r.POST("/evoluciones", h.CrearEvolucion)
	r.GET("/atenciones/:id/evoluciones", h.ListarPorAtencion)
}

func (h *EvolucionHandler) CrearEvolucion(c *gin.Context) {
	var req application.CrearEvolucionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "payload invalido"})
		return
	}

	resp, err := h.crearUC.Execute(req)
	if err != nil {
		if err == application.ErrDatosEvolucionInvalidos {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error interno"})
		return
	}

	c.JSON(http.StatusCreated, resp)
}

func (h *EvolucionHandler) ListarPorAtencion(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id invalido"})
		return
	}

	res, err := h.listarUC.Execute(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error interno"})
		return
	}

	c.JSON(http.StatusOK, res)
}
