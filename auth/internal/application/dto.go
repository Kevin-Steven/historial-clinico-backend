package application

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
	Rol         string `json:"rol"`
	UsuarioID   int64  `json:"usuario_id"`
}

type CrearUsuarioRequest struct {
	Username      string `json:"username"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	Rol           string `json:"rol"`
	ProfesionalID *int64 `json:"id_profesional"`
}

type UsuarioResponse struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Rol      string `json:"rol"`
}
