package application

import "historial-clinico-backend/atenciones/internal/domain"

type GetAtencionUseCase struct {
	repo domain.AtencionRepository
}

func NewGetAtencionUseCase(repo domain.AtencionRepository) *GetAtencionUseCase {
	return &GetAtencionUseCase{repo: repo}
}

func (uc *GetAtencionUseCase) Execute(id int64) (*AtencionResponse, error) {
	at, err := uc.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if at == nil {
		return nil, nil
	}
	return &AtencionResponse{
		ID:                at.ID,
		PacienteID:        at.PacienteID,
		EstablecimientoID: at.EstablecimientoID,
		TipoAtencionID:    at.TipoAtencionID,
	}, nil
}
