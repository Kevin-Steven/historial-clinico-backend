package domain

import "time"

type EvolucionPrescripcion struct {
	ID             int64
	AtencionID     int64
	FechaHora      time.Time
	NotaEvolucion  *string
	Farmacoterapia *string
	Indicaciones   *string
	ProfesionalID  *int64
	UsuarioID      *int64
}
