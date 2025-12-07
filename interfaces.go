package src

// Playable: interfaz para contenido "reproducible"
type Playable interface {
	GetTitulo() string
}

// Repository: interfaz para CRUD futuro
type Repository[T any] interface {
	Add(item T) error
	GetAll() []T
}
