package domain

type SignoVitalRepository interface {
	Create(s *SignoVital) error
	GetByAtencionID(atencionID int64) ([]SignoVital, error)
}
