package application

import (
	"errors"
	"time"

	"historial-clinico-backend/pacientes/internal/domain"
)

var (
	ErrDatosObligatorios      = errors.New("faltan datos obligatorios del paciente")
	ErrFechaNacimientoInvalida = errors.New("fecha_nacimiento con formato inválido, use YYYY-MM-DD")
	ErrPacienteDuplicado      = errors.New("ya existe un paciente con ese numero_documento")
)

type RegistrarPacienteUseCase struct {
	repo domain.PacienteRepository
}

func NewRegistrarPacienteUseCase(repo domain.PacienteRepository) *RegistrarPacienteUseCase {
	return &RegistrarPacienteUseCase{repo: repo}
}

func (uc *RegistrarPacienteUseCase) Execute(req RegistrarPacienteRequest) (*PacienteResponse, error) {
	if req.PrimerApellido == "" || req.PrimerNombre == "" || req.Sexo == "" {
		return nil, ErrDatosObligatorios
	}

	var fechaNac *time.Time
	if req.FechaNacimiento != nil && *req.FechaNacimiento != "" {
		t, err := time.Parse("2006-01-02", *req.FechaNacimiento)
		if err != nil {
			return nil, ErrFechaNacimientoInvalida
		}
		fechaNac = &t
	}

	p := domain.Paciente{
		NumeroHClinica:  req.NumeroHClinica,
		PrimerApellido:  req.PrimerApellido,
		SegundoApellido: req.SegundoApellido,
		PrimerNombre:    req.PrimerNombre,
		SegundoNombre:   req.SegundoNombre,
		Sexo:            req.Sexo,
		FechaNacimiento: fechaNac,
		TipoDocumento:   req.TipoDocumento,
		NumeroDocumento: req.NumeroDocumento,
		Telefono:        req.Telefono,
		Direccion:       req.Direccion,
	}

	if req.NumeroDocumento != nil {
		existente, err := uc.repo.GetByNumeroDocumento(*req.NumeroDocumento)
		if err != nil {
			return nil, err
		}
		if existente != nil {
			return nil, ErrPacienteDuplicado
		}
	}

	id, err := uc.repo.Create(p)
	if err != nil {
		return nil, err
	}

	resp := &PacienteResponse{
		ID:             id,
		NumeroHClinica: p.NumeroHClinica,
		NombreCompleto: p.PrimerNombre + " " + p.PrimerApellido,
		Sexo:           p.Sexo,
		TipoDocumento:  p.TipoDocumento,
		NumeroDocumento: p.NumeroDocumento,
	}

	return resp, nil
}
