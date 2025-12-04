package domain

type EvolucionRepository interface {
	Create(e *EvolucionPrescripcion) error
	ListByAtencionID(atencionID int64) ([]EvolucionPrescripcion, error)
}
