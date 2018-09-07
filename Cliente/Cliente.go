package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"github.com/RodrigoVelazquez99/proyecto1/Usuario"
)

func main() {

	if len(os.Args) != 3 {
		fmt.Println("No agregaste correctamente los datos de la direccion y puerto")
		os.Exit(1)
	}

	var nombre string

	fmt.Println("Introduce el nombre de usuario")
	fmt.Scanln(& nombre)
	fmt.Println("Bienvenido " + nombre)
	conexion, err := net.Dial("tcp", os.Args[1] + ":" + os.Args[2])

	if err != nil {
		err.Error()
	}

	Usuario.AgregaUsuario(nombre, conexion)

	for {

		informacionDevuelta := bufio.NewReader(os.Stdin)
		fmt.Print(nombre + ">")
		lector, _ := informacionDevuelta.ReadString('\n')
		fmt.Fprintf(conexion, lector+"\n")
		mensaje, _ := bufio.NewReader(conexion).ReadString('\n')
		fmt.Print(string(mensaje))

	}
}
