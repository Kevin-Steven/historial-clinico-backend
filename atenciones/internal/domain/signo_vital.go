package domain

import "time"

type SignoVital struct {
	ID                int64
	AtencionID        int64
	FechaMedicion     time.Time
	Origen            string
	TemperaturaC      *float64
	PresionSistolica  *int
	PresionDiastolica *int
	PulsoXMin         *int
	FrecuenciaResp    *int
	PesoKg            *float64
	TallaCm           *float64
}
