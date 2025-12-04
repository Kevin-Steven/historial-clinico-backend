package domain

type PacienteFilter struct {
	NumeroDocumento *string
	Nombre          *string
}

type PacienteRepository interface {
	Create(p Paciente) (int64, error)
	Update(p Paciente) error
	GetByID(id int64) (*Paciente, error)
	GetByNumeroDocumento(doc string) (*Paciente, error)
	Search(filter PacienteFilter, limit, offset int) ([]Paciente, error)
}
