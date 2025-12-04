package repository

import (
	"errors"

	"gorm.io/gorm"

	"historial-clinico-backend/auth/internal/domain"
)

type ProfesionalModel struct {
	ID             int64   `gorm:"column:id_profesional;primaryKey;autoIncrement"`
	NombreCompleto string  `gorm:"column:nombre_completo"`
	Especialidad   *string `gorm:"column:especialidad"`
	NumeroRegistro *string `gorm:"column:numero_registro"`
}

func (ProfesionalModel) TableName() string { return "profesionales" }

func (r *MySQLUsuarioRepository) GetProfesionalByID(id int64) (*domain.Profesional, error) {
	var m ProfesionalModel
	if err := r.db.First(&m, "id_profesional = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &domain.Profesional{
		ID:             m.ID,
		NombreCompleto: m.NombreCompleto,
		Especialidad:   m.Especialidad,
		NumeroRegistro: m.NumeroRegistro,
	}, nil
}
