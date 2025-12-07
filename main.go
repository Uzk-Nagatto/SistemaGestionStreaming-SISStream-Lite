package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"SistemaGestionStreaming-SISStream-Lite/src"
)

func main() {
	sys := src.NewStreamingSystem()
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n--- SISStream Lite ---")
		fmt.Println("1. Registrar usuario")
		fmt.Println("2. Registrar contenido")
		fmt.Println("3. Activar suscripción")
		fmt.Println("4. Listar contenidos")
		fmt.Println("5. Salir")
		fmt.Print("Selecciona una opción: ")

		opStr, _ := reader.ReadString('\n')
		opStr = strings.TrimSpace(opStr)
		op, _ := strconv.Atoi(opStr)

		switch op {
		case 1:
			fmt.Print("Nombre: ")
			nombre, _ := reader.ReadString('\n')
			fmt.Print("Correo: ")
			correo, _ := reader.ReadString('\n')
			fmt.Print("Clave: ")
			clave, _ := reader.ReadString('\n')

			u := src.NewUser(len(sys.ListarContenidos())+1,
				strings.TrimSpace(nombre),
				strings.TrimSpace(correo),
				strings.TrimSpace(clave),
			)

			if err := sys.RegistrarUsuario(u); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Usuario registrado correctamente.")
			}

		case 2:
			fmt.Print("Título: ")
			titulo, _ := reader.ReadString('\n')
			fmt.Print("Género: ")
			genero, _ := reader.ReadString('\n')
			fmt.Print("Año: ")
			anioStr, _ := reader.ReadString('\n')
			anio, _ := strconv.Atoi(strings.TrimSpace(anioStr))

			c := src.NewContent(len(sys.ListarContenidos())+1,
				strings.TrimSpace(titulo),
				strings.TrimSpace(genero),
				anio,
			)
			sys.RegistrarContenido(c)
			fmt.Println("Contenido registrado.")

		case 3:
			fmt.Print("Correo del usuario: ")
			correo, _ := reader.ReadString('\n')
			correo = strings.TrimSpace(correo)

			if err := sys.ActivarSuscripcion(correo); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Suscripción activada.")
			}

		case 4:
			fmt.Println("\nCatálogo:")
			for _, c := range sys.ListarContenidos() {
				fmt.Println("-", c.GetTitulo())
			}

		case 5:
			fmt.Println("Saliendo...")
			return

		default:
			fmt.Println("Opción inválida.")
		}
	}
}
