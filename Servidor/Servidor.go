package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"github.com/RodrigoVelazquez99/proyecto1/Usuario"
)

/* Inicia un servidor */
func main() {

	if len(os.Args) != 3 {
		fmt.Println("No agregaste correctamente los datos de la direccion y puerto")
		os.Exit(1)
	}

	fmt.Println("  Iniciando el servidor ..... ")

	servidor, conexion1 := net.Listen("tcp", os.Args[1] + ":" + os.Args[2])

	revisaError(conexion1)

	for {
		conexion, disponible := servidor.Accept()
		if disponible != nil {
			fmt.Println("Esperando conexion")
			os.Exit(1)
		}
		 manejaCliente(conexion)
	}
}

func revisaError(err error) {
	if err != nil {
		fmt.Println("Fallo la conexion al servidor")
	}
}

func manejaCliente(conexion net.Conn) {
	usuarioActual := Usuario.BuscaUsuario(conexion)
	for {
		mensaje, _ := bufio.NewReader(conexion).ReadString('\n')
		fmt.Println(usuarioActual + ": " + string(mensaje))
		conexion.Write([]byte(mensaje))
		if mensaje == "/close" {
			conexion.Close()
			break
		}
	}
}
