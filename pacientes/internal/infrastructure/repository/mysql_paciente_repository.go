package repository

import (
	"errors"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"historial-clinico-backend/pacientes/internal/config"
	"historial-clinico-backend/pacientes/internal/domain"
)

type PacienteModel struct {
	ID              int64  `gorm:"column:id_paciente;primaryKey;autoIncrement"`
	NumeroHClinica  string `gorm:"column:numero_hclinica"`
	PrimerApellido  string `gorm:"column:primer_apellido"`
	SegundoApellido *string `gorm:"column:segundo_apellido"`
	PrimerNombre    string `gorm:"column:primer_nombre"`
	SegundoNombre   *string `gorm:"column:segundo_nombre"`
	Sexo            string `gorm:"column:sexo"`
	FechaNacimiento *gorm.DeletedAt `gorm:"column:fecha_nacimiento"` // usaremos time en mapping manual
	TipoDocumento   *string `gorm:"column:tipo_documento"`
	NumeroDocumento *string `gorm:"column:numero_documento"`
	Telefono        *string `gorm:"column:telefono"`
	Direccion       *string `gorm:"column:direccion"`
}

func (PacienteModel) TableName() string { return "pacientes" }

func NewMySQLDB(cfg config.Config) (*gorm.DB, error) {
	if cfg.DSN == "" {
		return nil, errors.New("PACIENTES_DB_DSN no configurado")
	}
	return gorm.Open(mysql.Open(cfg.DSN), &gorm.Config{})
}

type MySQLPacienteRepository struct {
	db *gorm.DB
}

func NewMySQLPacienteRepository(db *gorm.DB) *MySQLPacienteRepository {
	return &MySQLPacienteRepository{db: db}
}

func (r *MySQLPacienteRepository) toDomain(m PacienteModel) domain.Paciente {
	// Nota: aquí podrías mapear manualmente fechas si las usas
	return domain.Paciente{
		ID:              m.ID,
		NumeroHClinica:  m.NumeroHClinica,
		PrimerApellido:  m.PrimerApellido,
		SegundoApellido: m.SegundoApellido,
		PrimerNombre:    m.PrimerNombre,
		SegundoNombre:   m.SegundoNombre,
		Sexo:            m.Sexo,
		TipoDocumento:   m.TipoDocumento,
		NumeroDocumento: m.NumeroDocumento,
		Telefono:        m.Telefono,
		Direccion:       m.Direccion,
	}
}

func (r *MySQLPacienteRepository) Create(p domain.Paciente) (int64, error) {
	m := PacienteModel{
		NumeroHClinica:  p.NumeroHClinica,
		PrimerApellido:  p.PrimerApellido,
		SegundoApellido: p.SegundoApellido,
		PrimerNombre:    p.PrimerNombre,
		SegundoNombre:   p.SegundoNombre,
		Sexo:            p.Sexo,
		TipoDocumento:   p.TipoDocumento,
		NumeroDocumento: p.NumeroDocumento,
		Telefono:        p.Telefono,
		Direccion:       p.Direccion,
	}
	if err := r.db.Create(&m).Error; err != nil {
		return 0, err
	}
	return m.ID, nil
}

func (r *MySQLPacienteRepository) Update(p domain.Paciente) error {
	return r.db.Model(&PacienteModel{}).
		Where("id_paciente = ?", p.ID).
		Updates(PacienteModel{
			NumeroHClinica:  p.NumeroHClinica,
			PrimerApellido:  p.PrimerApellido,
			SegundoApellido: p.SegundoApellido,
			PrimerNombre:    p.PrimerNombre,
			SegundoNombre:   p.SegundoNombre,
			Sexo:            p.Sexo,
			TipoDocumento:   p.TipoDocumento,
			NumeroDocumento: p.NumeroDocumento,
			Telefono:        p.Telefono,
			Direccion:       p.Direccion,
		}).Error
}

func (r *MySQLPacienteRepository) GetByID(id int64) (*domain.Paciente, error) {
	var m PacienteModel
	if err := r.db.First(&m, "id_paciente = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	d := r.toDomain(m)
	return &d, nil
}

func (r *MySQLPacienteRepository) GetByNumeroDocumento(doc string) (*domain.Paciente, error) {
	var m PacienteModel
	if err := r.db.Where("numero_documento = ?", doc).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	d := r.toDomain(m)
	return &d, nil
}

func (r *MySQLPacienteRepository) Search(filter domain.PacienteFilter, limit, offset int) ([]domain.Paciente, error) {
	var models []PacienteModel
	q := r.db.Model(&PacienteModel{})
	if filter.NumeroDocumento != nil {
		q = q.Where("numero_documento = ?", *filter.NumeroDocumento)
	}
	if filter.Nombre != nil {
		like := "%" + *filter.Nombre + "%"
		q = q.Where("primer_nombre LIKE ? OR primer_apellido LIKE ?", like, like)
	}
	if limit > 0 {
		q = q.Limit(limit)
	}
	if offset > 0 {
		q = q.Offset(offset)
	}
	if err := q.Find(&models).Error; err != nil {
		return nil, err
	}
	res := make([]domain.Paciente, 0, len(models))
	for _, m := range models {
		res = append(res, r.toDomain(m))
	}
	return res, nil
}
