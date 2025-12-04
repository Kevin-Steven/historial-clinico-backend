package application

import "time"

type RegistrarSignosRequest struct {
	FechaMedicion     *string  `json:"fecha_medicion"` // opcional, ISO8601; si no viene, now
	Origen            string   `json:"origen"`
	TemperaturaC      *float64 `json:"temperatura_c"`
	PresionSistolica  *int     `json:"presion_sistolica"`
	PresionDiastolica *int     `json:"presion_diastolica"`
	PulsoXMin         *int     `json:"pulso_x_min"`
	FrecuenciaResp    *int     `json:"frecuencia_resp"`
	PesoKg            *float64 `json:"peso_kg"`
	TallaCm           *float64 `json:"talla_cm"`
}

type SignoVitalResponse struct {
	ID            int64     `json:"id"`
	AtencionID    int64     `json:"id_atencion"`
	FechaMedicion time.Time `json:"fecha_medicion"`
	Origen        string    `json:"origen"`
}
