package domain

type AtencionRepository interface {
	Create(a *Atencion) error
	GetByID(id int64) (*Atencion, error)
}
