package application

type RegistrarPacienteRequest struct {
	NumeroHClinica  string  `json:"numero_hclinica"`
	PrimerApellido  string  `json:"primer_apellido"`
	SegundoApellido *string `json:"segundo_apellido"`
	PrimerNombre    string  `json:"primer_nombre"`
	SegundoNombre   *string `json:"segundo_nombre"`
	Sexo            string  `json:"sexo"`
	FechaNacimiento *string `json:"fecha_nacimiento"`
	TipoDocumento   *string `json:"tipo_documento"`
	NumeroDocumento *string `json:"numero_documento"`
	Telefono        *string `json:"telefono"`
	Direccion       *string `json:"direccion"`
}

type PacienteResponse struct {
	ID              int64   `json:"id"`
	NumeroHClinica  string  `json:"numero_hclinica"`
	NombreCompleto  string  `json:"nombre_completo"`
	Sexo            string  `json:"sexo"`
	Edad            *int    `json:"edad,omitempty"`
	TipoDocumento   *string `json:"tipo_documento"`
	NumeroDocumento *string `json:"numero_documento"`
}
