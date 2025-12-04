package application

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTService interface {
	GenerateToken(usuarioID int64, rol string) (string, error)
}

type jwtServiceImpl struct {
	secret []byte
}

func NewJWTService(secret string) JWTService {
	return &jwtServiceImpl{secret: []byte(secret)}
}

type CustomClaims struct {
	UsuarioID int64  `json:"uid"`
	Rol       string `json:"rol"`
	jwt.RegisteredClaims
}

func (s *jwtServiceImpl) GenerateToken(usuarioID int64, rol string) (string, error) {
	claims := CustomClaims{
		UsuarioID: usuarioID,
		Rol:       rol,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(8 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.secret)
}
