package src

import "time"

// User representa un usuario del sistema
type User struct {
	id                int
	nombre            string
	correo            string
	clave             string
	estadoSuscripcion string
}

// Constructor funcional estilo POO
func NewUser(id int, nombre, correo, clave string) *User {
	return &User{
		id:                id,
		nombre:            nombre,
		correo:            correo,
		clave:             clave,
		estadoSuscripcion: "vencida",
	}
}

// Getters (encapsulación)
func (u *User) GetCorreo() string            { return u.correo }
func (u *User) GetNombre() string            { return u.nombre }
func (u *User) GetEstadoSuscripcion() string { return u.estadoSuscripcion }

// Setter controlado
func (u *User) SetEstadoSuscripcion(estado string) {
	u.estadoSuscripcion = estado
}

// Content representa una película o serie
type Content struct {
	id     int
	titulo string
	genero string
	anio   int
}

func NewContent(id int, titulo, genero string, anio int) *Content {
	return &Content{id: id, titulo: titulo, genero: genero, anio: anio}
}

func (c *Content) GetTitulo() string { return c.titulo }

// Subscription gestiona suscripción de usuario
type Subscription struct {
	plan        string
	estado      string
	fechaInicio time.Time
	fechaFin    time.Time
}

func NewSubscription(plan string, dias int) *Subscription {
	inicio := time.Now()
	fin := inicio.AddDate(0, 0, dias)
	return &Subscription{
		plan: plan, estado: "activa",
		fechaInicio: inicio, fechaFin: fin,
	}
}

func (s *Subscription) EstaActiva() bool {
	return s.estado == "activa" && time.Now().Before(s.fechaFin)
}
