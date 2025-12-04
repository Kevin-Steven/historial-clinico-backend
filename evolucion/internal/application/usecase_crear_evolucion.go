package application

import (
	"time"

	"historial-clinico-backend/evolucion/internal/domain"
)

type CrearEvolucionUseCase struct {
	repo domain.EvolucionRepository
}

func NewCrearEvolucionUseCase(repo domain.EvolucionRepository) *CrearEvolucionUseCase {
	return &CrearEvolucionUseCase{repo: repo}
}

func (uc *CrearEvolucionUseCase) Execute(req CrearEvolucionRequest) (*EvolucionResponse, error) {
	if req.AtencionID == 0 {
		return nil, ErrDatosEvolucionInvalidos
	}

	var t time.Time
	var err error
	if req.FechaHora != nil && *req.FechaHora != "" {
		t, err = time.Parse(time.RFC3339, *req.FechaHora)
		if err != nil {
			return nil, ErrDatosEvolucionInvalidos
		}
	} else {
		t = time.Now()
	}

	e := &domain.EvolucionPrescripcion{
		AtencionID:     req.AtencionID,
		FechaHora:      t,
		NotaEvolucion:  req.NotaEvolucion,
		Farmacoterapia: req.Farmacoterapia,
		Indicaciones:   req.Indicaciones,
		ProfesionalID:  req.ProfesionalID,
		UsuarioID:      req.UsuarioID,
	}

	if err := uc.repo.Create(e); err != nil {
		return nil, err
	}

	return &EvolucionResponse{
		ID:         e.ID,
		AtencionID: e.AtencionID,
		FechaHora:  e.FechaHora,
	}, nil
}
