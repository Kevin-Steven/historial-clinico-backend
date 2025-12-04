package repository

import (
	"errors"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"historial-clinico-backend/consulta_externa/internal/config"
	"historial-clinico-backend/consulta_externa/internal/domain"
)

type ConsultaExternaModel struct {
	ID                     int64   `gorm:"column:id_consulta_externa;primaryKey;autoIncrement"`
	AtencionID             int64   `gorm:"column:id_atencion"`
	AntecedentesPersonales *string `gorm:"column:antecedentes_personales"`
	AntecedentesFamiliares *string `gorm:"column:antecedentes_familiares"`
	RevisionSistemas       *string `gorm:"column:revision_sistemas"`
	ExamenFisicoRegional   *string `gorm:"column:examen_fisico_regional"`
	PlanTratamiento        *string `gorm:"column:plan_tratamiento"`
}

func (ConsultaExternaModel) TableName() string { return "consulta_externa" }

func NewMySQLDB(cfg config.Config) (*gorm.DB, error) {
	if cfg.DSN == "" {
		return nil, errors.New("CONSULTA_EXT_DB_DSN no configurado")
	}
	return gorm.Open(mysql.Open(cfg.DSN), &gorm.Config{})
}

type MySQLConsultaExternaRepository struct {
	db *gorm.DB
}

func NewMySQLConsultaExternaRepository(db *gorm.DB) *MySQLConsultaExternaRepository {
	return &MySQLConsultaExternaRepository{db: db}
}

func (r *MySQLConsultaExternaRepository) Create(c *domain.ConsultaExterna) error {
	m := ConsultaExternaModel{
		AtencionID:             c.AtencionID,
		AntecedentesPersonales: c.AntecedentesPersonales,
		AntecedentesFamiliares: c.AntecedentesFamiliares,
		RevisionSistemas:       c.RevisionSistemas,
		ExamenFisicoRegional:   c.ExamenFisicoRegional,
		PlanTratamiento:        c.PlanTratamiento,
	}
	if err := r.db.Create(&m).Error; err != nil {
		return err
	}
	c.ID = m.ID
	return nil
}

func (r *MySQLConsultaExternaRepository) GetByID(id int64) (*domain.ConsultaExterna, error) {
	var m ConsultaExternaModel
	if err := r.db.First(&m, "id_consulta_externa = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &domain.ConsultaExterna{
		ID:                     m.ID,
		AtencionID:             m.AtencionID,
		AntecedentesPersonales: m.AntecedentesPersonales,
		AntecedentesFamiliares: m.AntecedentesFamiliares,
		RevisionSistemas:       m.RevisionSistemas,
		ExamenFisicoRegional:   m.ExamenFisicoRegional,
		PlanTratamiento:        m.PlanTratamiento,
	}, nil
}
