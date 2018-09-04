package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

type Usuario struct {
	nombre string
}

var listaUsuarios []Usuario

/* Inicia un servidor */
func main() {

	fmt.Println("Iniciando el servidor ..... ")

	servidor, conexion1 := net.Listen("tcp", "localhost:8080")

	listaUsuarios = make([]Usuario, 1)

	revisaError(conexion1)

	conexion, disponible := servidor.Accept()

	for {
		if disponible != nil {
			fmt.Println("Esperando conexion")
			continue
		}
		manejaCliente(conexion)
	}
}

func agregaUsuario(nombre string) {
	var nuevoUsuario Usuario
	nuevoUsuario.nombre = nombre
	listaUsuarios = append(listaUsuarios, nuevoUsuario)
}

func revisaError(err error) {
	if err != nil {
		fmt.Println("Fallo la conexion al servidor")
	}
}

func manejaCliente(conexion net.Conn) {
	mensaje, _ := bufio.NewReader(conexion).ReadString('\n')
	fmt.Println(string(mensaje))
	conexion.Write([]byte(mensaje))
	if mensaje == "/close" {
		os.Exit(1)
	}

}
