# SistemaGestionStreaming-SISStream-Lite
Proyecto académico desarrollado para la asignatura Programación Orientada a Objetos. Incluye la planificación, estructura funcional y primeros componentes del sistema de gestión de streaming SISStream Lite, implementado en Go


Materia: Programación Orientada a Objetos
Docente: Palacios Morocho Milton Ricardo
Estudiante: Christian José Escalante
Periodo: 2025

Programación funcional en Go

Uso de funciones puras

Manejo de estructuras como maps

Validación mediante expresiones regulares

Lógica de control, iteraciones y manejo de datos

El proyecto representa las primeras fases de un sistema real, enfocado en la gestión de usuarios, contenido y suscripciones.

Objetivo General

Diseñar y planificar un sistema de streaming simple que permita administrar usuarios, catálogo de contenido y estados de suscripción, empleando principios fundamentales de Go y lógica vista en clase.

Funcionalidades principales
✔ Gestión de Usuarios

Registrar un usuario

Validar correo y contraseña con expresiones regulares

Consultar suscripción activa

✔ Gestión de Contenido

Registrar películas o series

Listar contenido

Buscar contenido

✔ Gestión de Suscripciones

Activar suscripción

Verificar estado

Permitir o bloquear acceso al contenido

Tecnologías utilizadas

Lenguaje: Go (Golang)

Paradigma: Programación funcional

Paquetes estándar empleados:

fmt

bufio

os

regexp

time

Estructura recomendada del proyecto
SISStream-Lite/
│
├── src/
│   ├── users.go
│   ├── content.go
│   ├── subscriptions.go
│   └── utils.go
│
├── main.go
├── README.md
└── LICENSE

Ejemplo de lógica funcional en Go
func ValidarCorreo(correo string) bool {
    patron := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
    return patron.MatchString(correo)
}

func RegistrarUsuario(nombre, correo string) map[string]string {
    return map[string]string{
        "nombre": nombre,
        "correo": correo,
        "suscripcion": "inactiva",
    }
}

Arquitectura del sistema

(Aquí puedes insertar la imagen creada en Napkin AI)

Estado actual del proyecto

Este repositorio contiene:

La planeación completa del sistema

Estructura base en Go

Readme profesional

Código inicial para los módulos

Archivo LICENSE (MIT)

Más adelante se desarrollarán:

Persistencia de datos

Consola interactiva

Expansión de funcionalidades

Licencia

Este proyecto está bajo licencia MIT, lo que permite su uso académico y libre distribución.

Autor

Christian José Escalante
Estudiante de Programación Orientada a Objetos
