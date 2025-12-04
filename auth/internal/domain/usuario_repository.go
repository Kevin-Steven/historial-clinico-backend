package domain

type UsuarioRepository interface {
	Create(u *Usuario) error
	GetByUsername(username string) (*Usuario, error)
	GetByID(id int64) (*Usuario, error)
}
