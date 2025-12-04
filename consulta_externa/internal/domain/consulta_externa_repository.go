package domain

type ConsultaExternaRepository interface {
	Create(c *ConsultaExterna) error
	GetByID(id int64) (*ConsultaExterna, error)
}
