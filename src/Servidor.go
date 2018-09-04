package main

import (
	"bufio"
	"fmt"
	"net"
)

/* Inicia un servidor */
func main() {

	servidor, conexion1 := net.Listen("tcp", "localhost:8080")

	if conexion1 != nil {
		fmt.Println("No se creo el servidor")
	}

	for {
		conexion, disponible := servidor.Accept()
		if disponible != nil {
			fmt.Println("Esperando conexion")
			continue
		}
		mensaje, _ := bufio.NewReader(conexion).ReadString('\n')
		fmt.Println(mensaje)
		conexion.Write([]byte(mensaje))
		if mensaje == "close" {
			conexion.Close()
		}
	}
}
