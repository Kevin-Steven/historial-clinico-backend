package repository

import (
	"errors"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"historial-clinico-backend/atenciones/internal/config"
	"historial-clinico-backend/atenciones/internal/domain"
)

type AtencionModel struct {
	ID                int64      `gorm:"column:id_atencion;primaryKey;autoIncrement"`
	PacienteID        int64      `gorm:"column:id_paciente"`
	EstablecimientoID int64      `gorm:"column:id_establecimiento"`
	TipoAtencionID    int8       `gorm:"column:id_tipo_atencion"`
	FechaIngreso      time.Time  `gorm:"column:fecha_ingreso"`
	FechaAlta         *time.Time `gorm:"column:fecha_alta"`
	MotivoConsulta    *string    `gorm:"column:motivo_consulta"`
	EnfermedadActual  *string    `gorm:"column:enfermedad_actual"`
	UsuarioCreaID     *int64     `gorm:"column:id_usuario_crea"`
}

func (AtencionModel) TableName() string { return "atenciones" }

func NewMySQLDB(cfg config.Config) (*gorm.DB, error) {
	if cfg.DSN == "" {
		return nil, errors.New("ATENCIONES_DB_DSN no configurado")
	}
	return gorm.Open(mysql.Open(cfg.DSN), &gorm.Config{})
}

type MySQLAtencionRepository struct {
	db *gorm.DB
}

func NewMySQLAtencionRepository(db *gorm.DB) *MySQLAtencionRepository {
	return &MySQLAtencionRepository{db: db}
}

func (r *MySQLAtencionRepository) Create(a *domain.Atencion) error {
	m := AtencionModel{
		PacienteID:        a.PacienteID,
		EstablecimientoID: a.EstablecimientoID,
		TipoAtencionID:    a.TipoAtencionID,
		FechaIngreso:      a.FechaIngreso,
		FechaAlta:         a.FechaAlta,
		MotivoConsulta:    a.MotivoConsulta,
		EnfermedadActual:  a.EnfermedadActual,
		UsuarioCreaID:     a.UsuarioCreaID,
	}
	if err := r.db.Create(&m).Error; err != nil {
		return err
	}
	a.ID = m.ID
	return nil
}

func (r *MySQLAtencionRepository) GetByID(id int64) (*domain.Atencion, error) {
	var m AtencionModel
	if err := r.db.First(&m, "id_atencion = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &domain.Atencion{
		ID:                m.ID,
		PacienteID:        m.PacienteID,
		EstablecimientoID: m.EstablecimientoID,
		TipoAtencionID:    m.TipoAtencionID,
		FechaIngreso:      m.FechaIngreso,
		FechaAlta:         m.FechaAlta,
		MotivoConsulta:    m.MotivoConsulta,
		EnfermedadActual:  m.EnfermedadActual,
		UsuarioCreaID:     m.UsuarioCreaID,
	}, nil
}
