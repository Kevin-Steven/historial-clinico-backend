package application

import (
	"time"

	"historial-clinico-backend/signos_vitales/internal/domain"
)

type RegistrarSignosUseCase struct {
	signosRepo   domain.SignoVitalRepository
	atencionRepo domain.AtencionRepository
}

func NewRegistrarSignosUseCase(signosRepo domain.SignoVitalRepository, atencionRepo domain.AtencionRepository) *RegistrarSignosUseCase {
	return &RegistrarSignosUseCase{signosRepo: signosRepo, atencionRepo: atencionRepo}
}

func (uc *RegistrarSignosUseCase) Execute(atencionID int64, req RegistrarSignosRequest) (*SignoVitalResponse, error) {
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

	var fecha time.Time
	if req.FechaMedicion != nil && *req.FechaMedicion != "" {
		fecha, err = time.Parse(time.RFC3339, *req.FechaMedicion)
		if err != nil {
			return nil, ErrDatosSignosInvalidos
		}
	} else {
		fecha = time.Now()
	}

	if req.Origen == "" {
		req.Origen = "TRIAGE"
	}

	s := &domain.SignoVital{
		AtencionID:        atencionID,
		FechaMedicion:     fecha,
		Origen:            req.Origen,
		TemperaturaC:      req.TemperaturaC,
		PresionSistolica:  req.PresionSistolica,
		PresionDiastolica: req.PresionDiastolica,
		PulsoXMin:         req.PulsoXMin,
		FrecuenciaResp:    req.FrecuenciaResp,
		PesoKg:            req.PesoKg,
		TallaCm:           req.TallaCm,
	}

	if err := uc.signosRepo.Create(s); err != nil {
		return nil, err
	}

	return &SignoVitalResponse{
		ID:            s.ID,
		AtencionID:    s.AtencionID,
		FechaMedicion: s.FechaMedicion,
		Origen:        s.Origen,
	}, nil
}
