package domain

import "time"

type EstadoUsuario string

type RolUsuario string

const (
	EstadoActivo    EstadoUsuario = "ACTIVO"
	EstadoInactivo  EstadoUsuario = "INACTIVO"
	EstadoBloqueado EstadoUsuario = "BLOQUEADO"

	RolAdmin  RolUsuario = "ADMIN"
	RolMedico RolUsuario = "MEDICO"
)

type Usuario struct {
	ID            int64
	Username      string
	Email         string
	PasswordHash  string
	Estado        EstadoUsuario
	Rol           RolUsuario
	ProfesionalID *int64
	FechaCreacion time.Time
	UltimoLogin   *time.Time
}
