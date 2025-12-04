package domain

type ConsultaExterna struct {
	ID                     int64
	AtencionID             int64
	AntecedentesPersonales *string
	AntecedentesFamiliares *string
	RevisionSistemas       *string
	ExamenFisicoRegional   *string
	PlanTratamiento        *string
}
