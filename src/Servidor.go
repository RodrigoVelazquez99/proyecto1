package main

import (
	"bufio"
	"fmt"
	"net"
)

/* Inicia un servidor */
func main() {
	fmt.Println("Iniciando el servidor ")

	servidor, conexion1 := net.Listen("tcp", "localhost:8080")

	if conexion1 != nil {
		fmt.Println("No se creo el servidor")
	}

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
