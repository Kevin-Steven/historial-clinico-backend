package http

import (
	"net/http"
	"strconv"

	"historial-clinico-backend/atenciones/internal/application"

	"github.com/gin-gonic/gin"
)

type AtencionesHandler struct {
	crearUC           *application.CrearAtencionUseCase
	getUC             *application.GetAtencionUseCase
	registrarSignosUC *application.RegistrarSignosUseCase
}

func NewAtencionesHandler(
	crearUC *application.CrearAtencionUseCase,
	getUC *application.GetAtencionUseCase,
	registrarSignosUC *application.RegistrarSignosUseCase,
) *AtencionesHandler {
	return &AtencionesHandler{
		crearUC:           crearUC,
		getUC:             getUC,
		registrarSignosUC: registrarSignosUC,
	}
}

func (h *AtencionesHandler) RegisterRoutes(r *gin.Engine) {
	grp := r.Group("/atenciones")
	{
		grp.POST("", h.CrearAtencion)
		grp.GET(":id", h.GetAtencion)
		grp.POST(":id/signos-vitales", h.RegistrarSignos)
	}
}

func (h *AtencionesHandler) CrearAtencion(c *gin.Context) {
	var req application.CrearAtencionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "payload invalido"})
		return
	}

	resp, err := h.crearUC.Execute(req)
	if err != nil {
		if err == application.ErrDatosAtencionInvalidos {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error interno"})
		return
	}

	c.JSON(http.StatusCreated, resp)
}

func (h *AtencionesHandler) GetAtencion(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id invalido"})
		return
	}

	resp, err := h.getUC.Execute(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error interno"})
		return
	}
	if resp == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "atencion no encontrada"})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *AtencionesHandler) RegistrarSignos(c *gin.Context) {
	idStr := c.Param("id")
	atencionID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id invalido"})
		return
	}

	var req application.RegistrarSignosRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "payload invalido"})
		return
	}

	resp, err := h.registrarSignosUC.Execute(atencionID, req)
	if err != nil {
		if err == application.ErrDatosSignosInvalidos {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error interno"})
		return
	}

	c.JSON(http.StatusCreated, resp)
}
