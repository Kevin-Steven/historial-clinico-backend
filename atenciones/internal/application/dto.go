package application

type CrearAtencionRequest struct {
	PacienteID        int64   `json:"id_paciente"`
	EstablecimientoID int64   `json:"id_establecimiento"`
	TipoAtencionID    int8    `json:"id_tipo_atencion"`
	MotivoConsulta    *string `json:"motivo_consulta"`
	EnfermedadActual  *string `json:"enfermedad_actual"`
	UsuarioCreaID     *int64  `json:"id_usuario_crea"`
}

type AtencionResponse struct {
	ID                int64 `json:"id"`
	PacienteID        int64 `json:"id_paciente"`
	EstablecimientoID int64 `json:"id_establecimiento"`
	TipoAtencionID    int8  `json:"id_tipo_atencion"`
}
