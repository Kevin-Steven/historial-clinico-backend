package application

import (
	"historial-clinico-backend/consulta_externa/internal/domain"
)

type CrearConsultaExternaUseCase struct {
	repo domain.ConsultaExternaRepository
}

func NewCrearConsultaExternaUseCase(repo domain.ConsultaExternaRepository) *CrearConsultaExternaUseCase {
	return &CrearConsultaExternaUseCase{repo: repo}
}

func (uc *CrearConsultaExternaUseCase) Execute(req CrearConsultaExternaRequest) (*ConsultaExternaResponse, error) {
	if req.AtencionID == 0 {
		return nil, ErrDatosConsultaInvalidos
	}

	c := &domain.ConsultaExterna{
		AtencionID:             req.AtencionID,
		AntecedentesPersonales: req.AntecedentesPersonales,
		AntecedentesFamiliares: req.AntecedentesFamiliares,
		RevisionSistemas:       req.RevisionSistemas,
		ExamenFisicoRegional:   req.ExamenFisicoRegional,
		PlanTratamiento:        req.PlanTratamiento,
	}

	if err := uc.repo.Create(c); err != nil {
		return nil, err
	}

	return &ConsultaExternaResponse{
		ID:         c.ID,
		AtencionID: c.AtencionID,
	}, nil
}
