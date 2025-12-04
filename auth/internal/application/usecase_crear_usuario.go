package application

import (
	"errors"

	"golang.org/x/crypto/bcrypt"

	"historial-clinico-backend/auth/internal/domain"
)

var (
	ErrDatosUsuarioInvalidos = errors.New("datos de usuario invalidos")
)

type CrearUsuarioUseCase struct {
	repo domain.UsuarioRepository
}

func NewCrearUsuarioUseCase(repo domain.UsuarioRepository) *CrearUsuarioUseCase {
	return &CrearUsuarioUseCase{repo: repo}
}

func (uc *CrearUsuarioUseCase) Execute(req CrearUsuarioRequest) (*UsuarioResponse, error) {
	if req.Username == "" || req.Email == "" || req.Password == "" || req.Rol == "" {
		return nil, ErrDatosUsuarioInvalidos
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	usuario := &domain.Usuario{
		Username:      req.Username,
		Email:         req.Email,
		PasswordHash:  string(hash),
		Estado:        domain.EstadoActivo,
		Rol:           domain.RolUsuario(req.Rol),
		ProfesionalID: req.ProfesionalID,
	}

	if err := uc.repo.Create(usuario); err != nil {
		return nil, err
	}

	return &UsuarioResponse{
		ID:       usuario.ID,
		Username: usuario.Username,
		Email:    usuario.Email,
		Rol:      string(usuario.Rol),
	}, nil
}
