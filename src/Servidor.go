package main

/* Controla el serivdor */
import (
	"bufio"
	"fmt"
	"net"
)

/* Inicia un servidor */
func creaServidor(ip string) {
	direccion, conexion := net.ResolveTCPAddr("tcp", ip)
	if conexion != nil {
		fmt.Println(conexion.Error())
	}
	servidor, conexion1 := net.ListenTCP("tcp", direccion)
	if conexion1 != nil {
		fmt.Println("No se creo el servidor")
	}
	ejecutaServidor(direccion, servidor)
}

func ejecutaServidor(direccion *net.TCPAddr, servidor net.Listener) {
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
