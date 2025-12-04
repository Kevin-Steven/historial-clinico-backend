package domain

type ProfesionalRepository interface {
	GetByID(id int64) (*Profesional, error)
}
