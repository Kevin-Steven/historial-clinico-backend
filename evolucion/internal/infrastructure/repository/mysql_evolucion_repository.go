package repository

import (
	"errors"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"historial-clinico-backend/evolucion/internal/config"
	"historial-clinico-backend/evolucion/internal/domain"
)

type EvolucionModel struct {
	ID             int64     `gorm:"column:id_registro;primaryKey;autoIncrement"`
	AtencionID     int64     `gorm:"column:id_atencion"`
	FechaHora      time.Time `gorm:"column:fecha_hora"`
	NotaEvolucion  *string   `gorm:"column:nota_evolucion"`
	Farmacoterapia *string   `gorm:"column:farmacoterapia"`
	Indicaciones   *string   `gorm:"column:indicaciones"`
	ProfesionalID  *int64    `gorm:"column:id_profesional"`
	UsuarioID      *int64    `gorm:"column:id_usuario"`
}

func (EvolucionModel) TableName() string { return "evolucion_prescripciones" }

func NewMySQLDB(cfg config.Config) (*gorm.DB, error) {
	if cfg.DSN == "" {
		return nil, errors.New("EVOLUCION_DB_DSN no configurado")
	}
	return gorm.Open(mysql.Open(cfg.DSN), &gorm.Config{})
}

type MySQLEvolucionRepository struct {
	db *gorm.DB
}

func NewMySQLEvolucionRepository(db *gorm.DB) *MySQLEvolucionRepository {
	return &MySQLEvolucionRepository{db: db}
}

func (r *MySQLEvolucionRepository) Create(e *domain.EvolucionPrescripcion) error {
	m := EvolucionModel{
		AtencionID:     e.AtencionID,
		FechaHora:      e.FechaHora,
		NotaEvolucion:  e.NotaEvolucion,
		Farmacoterapia: e.Farmacoterapia,
		Indicaciones:   e.Indicaciones,
		ProfesionalID:  e.ProfesionalID,
		UsuarioID:      e.UsuarioID,
	}
	if err := r.db.Create(&m).Error; err != nil {
		return err
	}
	e.ID = m.ID
	return nil
}

func (r *MySQLEvolucionRepository) ListByAtencionID(atencionID int64) ([]domain.EvolucionPrescripcion, error) {
	var models []EvolucionModel
	if err := r.db.Where("id_atencion = ?", atencionID).Order("fecha_hora ASC").Find(&models).Error; err != nil {
		return nil, err
	}
	res := make([]domain.EvolucionPrescripcion, 0, len(models))
	for _, m := range models {
		res = append(res, domain.EvolucionPrescripcion{
			ID:             m.ID,
			AtencionID:     m.AtencionID,
			FechaHora:      m.FechaHora,
			NotaEvolucion:  m.NotaEvolucion,
			Farmacoterapia: m.Farmacoterapia,
			Indicaciones:   m.Indicaciones,
			ProfesionalID:  m.ProfesionalID,
			UsuarioID:      m.UsuarioID,
		})
	}
	return res, nil
}
