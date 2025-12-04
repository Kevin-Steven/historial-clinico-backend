package domain

import "time"

type Atencion struct {
	ID                int64
	PacienteID        int64
	EstablecimientoID int64
	TipoAtencionID    int8
	FechaIngreso      time.Time
	FechaAlta         *time.Time
	MotivoConsulta    *string
	EnfermedadActual  *string
	UsuarioCreaID     *int64
}
