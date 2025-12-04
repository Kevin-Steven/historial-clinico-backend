package application

import "historial-clinico-backend/evolucion/internal/domain"

type ListarEvolucionesUseCase struct {
	repo domain.EvolucionRepository
}

func NewListarEvolucionesUseCase(repo domain.EvolucionRepository) *ListarEvolucionesUseCase {
	return &ListarEvolucionesUseCase{repo: repo}
}

func (uc *ListarEvolucionesUseCase) Execute(atencionID int64) ([]EvolucionDetalleResponse, error) {
	res, err := uc.repo.ListByAtencionID(atencionID)
	if err != nil {
		return nil, err
	}
	out := make([]EvolucionDetalleResponse, 0, len(res))
	for _, e := range res {
		out = append(out, EvolucionDetalleResponse{
			ID:             e.ID,
			AtencionID:     e.AtencionID,
			FechaHora:      e.FechaHora,
			NotaEvolucion:  e.NotaEvolucion,
			Farmacoterapia: e.Farmacoterapia,
			Indicaciones:   e.Indicaciones,
		})
	}
	return out, nil
}
