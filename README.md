# SistemaGestionStreaming-SISStream-Lite
Proyecto académico desarrollado para la asignatura Programación Orientada a Objetos. Incluye la planificación, estructura funcional y primeros componentes del sistema de gestión de streaming SISStream Lite, implementado en Go


ChatGPT dijo:

Perfecto, Christian.
Aquí tienes un README más profesional, formal y limpio, listo para copiar y pegar directamente en GitHub, sin emojis ni adornos.

SISStream Lite
Sistema de Gestión de Streaming – Proyecto Académico

Materia: Programación Orientada a Objetos
Docente: Palacios Morocho Milton Ricardo
Estudiante: Christian José Escalante
Periodo académico: 2025

Descripción del Proyecto

SISStream Lite es un sistema básico de gestión de plataformas de streaming desarrollado como parte del Aprendizaje Autónomo 1 – Etapa 1: Planeación del Software. El proyecto se enmarca dentro del estudio de estructuras, funciones y principios base de la programación orientada a objetos y la lógica planteada durante la Unidad 1 de la materia.

El sistema se planifica siguiendo una arquitectura sencilla que incluye gestión de usuarios, catálogo de contenido y suscripciones, aplicando técnicas de programación funcional en Go y el uso apropiado de maps, funciones puras y expresiones regulares.

Objetivo General

Diseñar y planificar un sistema de gestión de streaming que permita administrar usuarios, contenido y suscripciones, integrando conceptos fundamentales de programación funcional y lógica de datos en Go.

Objetivos Específicos

Establecer la estructura funcional del sistema.

Definir los módulos principales necesarios para el funcionamiento básico.

Aplicar funciones puras, validaciones y uso de maps.

Documentar los paquetes utilizados durante el desarrollo.

Relacionar adecuadamente los contenidos de clase con el proyecto.

Módulos del Sistema
1. Gestión de Usuarios

Registro de usuarios.

Validación de correo y contraseña mediante expresiones regulares.

Consulta del estado de suscripción.

2. Gestión de Contenido

Registro de nuevos elementos de contenido.

Listado general del catálogo.

Búsqueda de contenido.

3. Gestión de Suscripciones

Activación o actualización de una suscripción.

Verificación de acceso.

Control básico de estados (activa / inactiva).

Tecnologías y Paquetes Utilizados

Lenguaje: Go (Golang)
Paradigma aplicado: Programación funcional

Paquetes estándar empleados:

fmt

bufio

os

regexp

time

Estructura del Proyecto
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

Ejemplo de Lógica Funcional
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

Estado Actual del Proyecto

Este repositorio contiene la planeación inicial, la estructura de organización del sistema y ejemplos base de implementación en Go.
En las siguientes etapas se añadirá:

Persistencia de datos.

Consola interactiva.

Ampliación de funcionalidades del sistema.

Licencia

Este proyecto está disponible bajo la licencia MIT, permitiendo su uso académico y libre distribución.

Autor

Christian José Escalante
Estudiante de Programación Orientada a Objetos
