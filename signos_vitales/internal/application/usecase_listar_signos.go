package application

import "historial-clinico-backend/signos_vitales/internal/domain"

type ListarSignosUseCase struct {
	signosRepo   domain.SignoVitalRepository
	atencionRepo domain.AtencionRepository
}

func NewListarSignosUseCase(signosRepo domain.SignoVitalRepository, atencionRepo domain.AtencionRepository) *ListarSignosUseCase {
	return &ListarSignosUseCase{signosRepo: signosRepo, atencionRepo: atencionRepo}
}

func (uc *ListarSignosUseCase) Execute(atencionID int64) ([]SignoVitalResponse, error) {
	if atencionID == 0 {
		return nil, ErrDatosSignosInvalidos
	}

	at, err := uc.atencionRepo.GetByID(atencionID)
	if err != nil {
		return nil, err
	}
	if at == nil {
		return nil, ErrDatosSignosInvalidos
	}

	items, err := uc.signosRepo.GetByAtencionID(atencionID)
	if err != nil {
		return nil, err
	}

	res := make([]SignoVitalResponse, 0, len(items))
	for _, s := range items {
		res = append(res, SignoVitalResponse{
			ID:            s.ID,
			AtencionID:    s.AtencionID,
			FechaMedicion: s.FechaMedicion,
			Origen:        s.Origen,
		})
	}
	return res, nil
}
