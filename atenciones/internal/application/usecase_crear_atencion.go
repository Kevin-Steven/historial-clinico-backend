package application

import (
	"errors"
	"time"

	"historial-clinico-backend/atenciones/internal/domain"
)

var (
	ErrDatosAtencionInvalidos = errors.New("datos de atencion invalidos")
)

type CrearAtencionUseCase struct {
	repo domain.AtencionRepository
}

func NewCrearAtencionUseCase(repo domain.AtencionRepository) *CrearAtencionUseCase {
	return &CrearAtencionUseCase{repo: repo}
}

func (uc *CrearAtencionUseCase) Execute(req CrearAtencionRequest) (*AtencionResponse, error) {
	if req.PacienteID == 0 || req.EstablecimientoID == 0 || req.TipoAtencionID == 0 {
		return nil, ErrDatosAtencionInvalidos
	}

	now := time.Now()

	at := &domain.Atencion{
		PacienteID:        req.PacienteID,
		EstablecimientoID: req.EstablecimientoID,
		TipoAtencionID:    req.TipoAtencionID,
		FechaIngreso:      now,
		MotivoConsulta:    req.MotivoConsulta,
		EnfermedadActual:  req.EnfermedadActual,
		UsuarioCreaID:     req.UsuarioCreaID,
	}

	if err := uc.repo.Create(at); err != nil {
		return nil, err
	}

	return &AtencionResponse{
		ID:                at.ID,
		PacienteID:        at.PacienteID,
		EstablecimientoID: at.EstablecimientoID,
		TipoAtencionID:    at.TipoAtencionID,
	}, nil
}
