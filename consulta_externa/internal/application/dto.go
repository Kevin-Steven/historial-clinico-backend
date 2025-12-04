package application

type CrearConsultaExternaRequest struct {
	AtencionID             int64   `json:"id_atencion"`
	AntecedentesPersonales *string `json:"antecedentes_personales"`
	AntecedentesFamiliares *string `json:"antecedentes_familiares"`
	RevisionSistemas       *string `json:"revision_sistemas"`
	ExamenFisicoRegional   *string `json:"examen_fisico_regional"`
	PlanTratamiento        *string `json:"plan_tratamiento"`
}

type ConsultaExternaResponse struct {
	ID         int64 `json:"id"`
	AtencionID int64 `json:"id_atencion"`
}
