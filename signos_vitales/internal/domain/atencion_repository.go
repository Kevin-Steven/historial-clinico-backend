package domain

type AtencionRepository interface {
	GetByID(id int64) (*Atencion, error)
}
