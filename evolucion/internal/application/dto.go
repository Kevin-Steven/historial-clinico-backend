package application

import "time"

type CrearEvolucionRequest struct {
	AtencionID     int64   `json:"id_atencion"`
	FechaHora      *string `json:"fecha_hora"`
	NotaEvolucion  *string `json:"nota_evolucion"`
	Farmacoterapia *string `json:"farmacoterapia"`
	Indicaciones   *string `json:"indicaciones"`
	ProfesionalID  *int64  `json:"id_profesional"`
	UsuarioID      *int64  `json:"id_usuario"`
}

type EvolucionResponse struct {
	ID         int64     `json:"id"`
	AtencionID int64     `json:"id_atencion"`
	FechaHora  time.Time `json:"fecha_hora"`
}

type EvolucionDetalleResponse struct {
	ID             int64     `json:"id"`
	AtencionID     int64     `json:"id_atencion"`
	FechaHora      time.Time `json:"fecha_hora"`
	NotaEvolucion  *string   `json:"nota_evolucion"`
	Farmacoterapia *string   `json:"farmacoterapia"`
	Indicaciones   *string   `json:"indicaciones"`
}
