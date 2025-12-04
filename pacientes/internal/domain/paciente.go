package domain

import "time"

type Paciente struct {
	ID              int64
	NumeroHClinica  string
	PrimerApellido  string
	SegundoApellido *string
	PrimerNombre    string
	SegundoNombre   *string
	Sexo            string
	FechaNacimiento *time.Time
	TipoDocumento   *string
	NumeroDocumento *string
	Telefono        *string
	Direccion       *string
}
