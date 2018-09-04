package main

import (
	"bufio"
	"fmt"
	"net"
)

type Usuario struct {
	nombre string
}

/* Lista de los usuarios del chat */
listaUsuarios := list.New()

/* Inicia un servidor */
func main() {

	fmt.Println("Iniciando el servidor ")

	servidor, conexion1 := net.Listen("tcp", "localhost:8080")

	revisaError(conexion1)

	conexion, disponible := servidor.Accept()

	for {
		if disponible != nil {
			fmt.Println("Esperando conexion")
			continue
		}
		mensaje, _ := bufio.NewReader(conexion).ReadString('\n')
		fmt.Println(string(mensaje))
		conexion.Write([]byte(mensaje))
		if mensaje == "close" {
			conexion.Close()
		}
	}
}

func revisaError(err error) {
	if err != nil {
		fmt.Println("Fallo la conexion al servidor")
	}
}
