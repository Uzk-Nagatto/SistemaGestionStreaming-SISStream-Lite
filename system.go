package src

import (
	"errors"
	"regexp"
)

// StreamingSystem coordina todo el sistema
type StreamingSystem struct {
	usuarios      map[string]*User // key: correo
	contenidos    []*Content
	suscripciones map[string]*Subscription // key: correo usuario
}

func NewStreamingSystem() *StreamingSystem {
	return &StreamingSystem{
		usuarios:      make(map[string]*User),
		contenidos:    []*Content{},
		suscripciones: make(map[string]*Subscription),
	}
}

// Validar correo con regex (comentado porque es parte clave)
func validarCorreo(correo string) bool {
	// Patrón simple para validar estructura de correo: texto@dominio.extension
	patron := `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(patron)
	return re.MatchString(correo)
}

func (s *StreamingSystem) RegistrarUsuario(u *User) error {
	if !validarCorreo(u.GetCorreo()) {
		return errors.New("correo inválido")
	}
	if _, existe := s.usuarios[u.GetCorreo()]; existe {
		return errors.New("usuario ya registrado")
	}
	s.usuarios[u.GetCorreo()] = u
	return nil
}

func (s *StreamingSystem) RegistrarContenido(c *Content) {
	s.contenidos = append(s.contenidos, c)
}

func (s *StreamingSystem) ActivarSuscripcion(correo string) error {
	u, ok := s.usuarios[correo]
	if !ok {
		return errors.New("usuario no existe")
	}
	sub := NewSubscription("Básico", 30)
	s.suscripciones[correo] = sub
	u.SetEstadoSuscripcion("activa")
	return nil
}

// ReproducirContenido verifica suscripción antes de permitir acceso
func (s *StreamingSystem) ReproducirContenido(correo string, contenido Playable) error {
	sub, ok := s.suscripciones[correo]
	if !ok || !sub.EstaActiva() {
		return errors.New("suscripción vencida o inexistente")
	}
	return nil
}

func (s *StreamingSystem) ListarContenidos() []*Content {
	return s.contenidos
}
