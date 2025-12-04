package repository

import (
	"time"

	"gorm.io/gorm"

	"historial-clinico-backend/signos_vitales/internal/domain"
)

type SignoVitalModel struct {
	ID                int64     `gorm:"column:id_signo;primaryKey;autoIncrement"`
	AtencionID        int64     `gorm:"column:id_atencion"`
	FechaMedicion     time.Time `gorm:"column:fecha_medicion"`
	Origen            string    `gorm:"column:origen"`
	TemperaturaC      *float64  `gorm:"column:temperatura_c"`
	PresionSistolica  *int      `gorm:"column:presion_sistolica"`
	PresionDiastolica *int      `gorm:"column:presion_diastolica"`
	PulsoXMin         *int      `gorm:"column:pulso_x_min"`
	FrecuenciaResp    *int      `gorm:"column:frecuencia_resp"`
	PesoKg            *float64  `gorm:"column:peso_kg"`
	TallaCm           *float64  `gorm:"column:talla_cm"`
}

func (SignoVitalModel) TableName() string { return "signos_vitales" }

type MySQLSignoVitalRepository struct {
	db *gorm.DB
}

func NewMySQLSignoVitalRepository(db *gorm.DB) *MySQLSignoVitalRepository {
	return &MySQLSignoVitalRepository{db: db}
}

func (r *MySQLSignoVitalRepository) Create(s *domain.SignoVital) error {
	m := SignoVitalModel{
		AtencionID:        s.AtencionID,
		FechaMedicion:     s.FechaMedicion,
		Origen:            s.Origen,
		TemperaturaC:      s.TemperaturaC,
		PresionSistolica:  s.PresionSistolica,
		PresionDiastolica: s.PresionDiastolica,
		PulsoXMin:         s.PulsoXMin,
		FrecuenciaResp:    s.FrecuenciaResp,
		PesoKg:            s.PesoKg,
		TallaCm:           s.TallaCm,
	}
	if err := r.db.Create(&m).Error; err != nil {
		return err
	}
	s.ID = m.ID
	return nil
}

func (r *MySQLSignoVitalRepository) GetByAtencionID(atencionID int64) ([]domain.SignoVital, error) {
	var models []SignoVitalModel
	if err := r.db.Where("id_atencion = ?", atencionID).Find(&models).Error; err != nil {
		return nil, err
	}
	res := make([]domain.SignoVital, 0, len(models))
	for _, m := range models {
		res = append(res, domain.SignoVital{
			ID:                m.ID,
			AtencionID:        m.AtencionID,
			FechaMedicion:     m.FechaMedicion,
			Origen:            m.Origen,
			TemperaturaC:      m.TemperaturaC,
			PresionSistolica:  m.PresionSistolica,
			PresionDiastolica: m.PresionDiastolica,
			PulsoXMin:         m.PulsoXMin,
			FrecuenciaResp:    m.FrecuenciaResp,
			PesoKg:            m.PesoKg,
			TallaCm:           m.TallaCm,
		})
	}
	return res, nil
}
