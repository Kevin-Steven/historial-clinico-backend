package application

import (
	"errors"

	"golang.org/x/crypto/bcrypt"

	"historial-clinico-backend/auth/internal/domain"
)

var (
	ErrCredencialesInvalidas = errors.New("credenciales invalidas")
)

type LoginUsuarioUseCase struct {
	repo      domain.UsuarioRepository
	tokenServ JWTService
}

func NewLoginUsuarioUseCase(repo domain.UsuarioRepository, tokenServ JWTService) *LoginUsuarioUseCase {
	return &LoginUsuarioUseCase{repo: repo, tokenServ: tokenServ}
}

func (uc *LoginUsuarioUseCase) Execute(req LoginRequest) (*LoginResponse, error) {
	u, err := uc.repo.GetByUsername(req.Username)
	if err != nil {
		return nil, err
	}
	if u == nil {
		return nil, ErrCredencialesInvalidas
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(req.Password)); err != nil {
		return nil, ErrCredencialesInvalidas
	}

	token, err := uc.tokenServ.GenerateToken(u.ID, string(u.Rol))
	if err != nil {
		return nil, err
	}

	return &LoginResponse{
		AccessToken: token,
		Rol:         string(u.Rol),
		UsuarioID:   u.ID,
	}, nil
}
