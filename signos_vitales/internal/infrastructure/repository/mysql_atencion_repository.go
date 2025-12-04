package repository

import (
	"gorm.io/gorm"

	"historial-clinico-backend/signos_vitales/internal/domain"
)

type AtencionModel struct {
	ID int64 `gorm:"column:id_atencion;primaryKey"`
}

func (AtencionModel) TableName() string { return "atenciones" }

type MySQLAtencionRepository struct {
	db *gorm.DB
}

func NewMySQLAtencionRepository(db *gorm.DB) *MySQLAtencionRepository {
	return &MySQLAtencionRepository{db: db}
}

func (r *MySQLAtencionRepository) GetByID(id int64) (*domain.Atencion, error) {
	var m AtencionModel
	if err := r.db.First(&m, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &domain.Atencion{ID: m.ID}, nil
}
