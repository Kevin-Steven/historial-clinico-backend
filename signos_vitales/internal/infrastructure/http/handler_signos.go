package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"historial-clinico-backend/signos_vitales/internal/application"
)

type SignosHandler struct {
	registrarUC *application.RegistrarSignosUseCase
	listarUC    *application.ListarSignosUseCase
}

func NewSignosHandler(registrarUC *application.RegistrarSignosUseCase, listarUC *application.ListarSignosUseCase) *SignosHandler {
	return &SignosHandler{registrarUC: registrarUC, listarUC: listarUC}
}

func (h *SignosHandler) Registrar(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id de atencion invalido"})
		return
	}

	var req application.RegistrarSignosRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "payload invalido"})
		return
	}

	res, err := h.registrarUC.Execute(id, req)
	if err != nil {
		if err == application.ErrDatosSignosInvalidos {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error registrando signos"})
		return
	}

	c.JSON(http.StatusCreated, res)
}

func (h *SignosHandler) ListarPorAtencion(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id de atencion invalido"})
		return
	}

	items, err := h.listarUC.Execute(id)
	if err != nil {
		if err == application.ErrDatosSignosInvalidos {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error listando signos"})
		return
	}

	c.JSON(http.StatusOK, items)
}
