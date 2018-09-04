package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

type Usuario struct {
	nombre string
	conexion net.Conn
}

var listaUsuarios []Usuario

/* Inicia un servidor */
func main() {

	if len(os.Args) != 3 {
		fmt.Println("No agregaste correctamente los datos de la direccion y puerto")
		os.Exit(1)
	}

	fmt.Println("  Iniciando el servidor ..... ")

	servidor, conexion1 := net.Listen("tcp", os.Args[1] + ":" + os.Args[2])

	listaUsuarios = make([]Usuario, 1)

	revisaError(conexion1)

	//conexion, disponible := servidor.Accept()

	for {
		conexion, disponible := servidor.Accept()
		if disponible != nil {
			fmt.Println("Esperando conexion")
			continue
		}
		 manejaCliente(conexion)
	}
}

func agregaUsuario(nombre string, conexion net.Conn) {
	var nuevoUsuario Usuario
	nuevoUsuario.nombre = nombre
	nuevoUsuario.conexion = conexion
	listaUsuarios = append(listaUsuarios, nuevoUsuario)
}

func revisaError(err error) {
	if err != nil {
		fmt.Println("Fallo la conexion al servidor")
	}
}

func manejaCliente(conexion net.Conn) {
	var usuarioActual string
	for _, u  := range listaUsuarios {
		if u.conexion == conexion{
		usuarioActual = u.nombre
		}
	}
	for {
		mensaje, _ := bufio.NewReader(conexion).ReadString('\n')
		fmt.Println(usuarioActual + ": " + string(mensaje))
		conexion.Write([]byte(mensaje))
		if mensaje == "/close" {
			conexion.Close()
			os.Exit(1)
			break
		}
	}
}
