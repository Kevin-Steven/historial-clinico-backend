package repository

import (
	"errors"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"historial-clinico-backend/auth/internal/config"
	"historial-clinico-backend/auth/internal/domain"
)

type UsuarioModel struct {
	ID            int64  `gorm:"column:id_usuario;primaryKey;autoIncrement"`
	Username      string `gorm:"column:username"`
	Email         string `gorm:"column:email"`
	PasswordHash  string `gorm:"column:password_hash"`
	Estado        string `gorm:"column:estado"`
	Rol           string `gorm:"column:rol"`
	ProfesionalID *int64 `gorm:"column:id_profesional"`
}

func (UsuarioModel) TableName() string { return "usuarios" }

func NewMySQLDB(cfg config.Config) (*gorm.DB, error) {
	if cfg.DSN == "" {
		return nil, errors.New("AUTH_DB_DSN no configurado")
	}
	return gorm.Open(mysql.Open(cfg.DSN), &gorm.Config{})
}

type MySQLUsuarioRepository struct {
	db *gorm.DB
}

func NewMySQLUsuarioRepository(db *gorm.DB) *MySQLUsuarioRepository {
	return &MySQLUsuarioRepository{db: db}
}

func (r *MySQLUsuarioRepository) toDomain(m UsuarioModel) *domain.Usuario {
	return &domain.Usuario{
		ID:            m.ID,
		Username:      m.Username,
		Email:         m.Email,
		PasswordHash:  m.PasswordHash,
		Estado:        domain.EstadoUsuario(m.Estado),
		Rol:           domain.RolUsuario(m.Rol),
		ProfesionalID: m.ProfesionalID,
	}
}

func (r *MySQLUsuarioRepository) Create(u *domain.Usuario) error {
	m := UsuarioModel{
		Username:      u.Username,
		Email:         u.Email,
		PasswordHash:  u.PasswordHash,
		Estado:        string(u.Estado),
		Rol:           string(u.Rol),
		ProfesionalID: u.ProfesionalID,
	}
	if err := r.db.Create(&m).Error; err != nil {
		return err
	}
	u.ID = m.ID
	return nil
}

func (r *MySQLUsuarioRepository) GetByUsername(username string) (*domain.Usuario, error) {
	var m UsuarioModel
	if err := r.db.Where("username = ?", username).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return r.toDomain(m), nil
}

func (r *MySQLUsuarioRepository) GetByID(id int64) (*domain.Usuario, error) {
	var m UsuarioModel
	if err := r.db.First(&m, "id_usuario = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return r.toDomain(m), nil
}
