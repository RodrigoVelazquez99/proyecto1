package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"log"
	"io"
)

var listaConexiones []net.Conn

func main() {

	if len(os.Args) != 3 {
		fmt.Println("No agregaste correctamente los datos de la direccion y puerto")
		os.Exit(1)
	}

	fmt.Println("  Iniciando el servidor .....  ")
	servidor, err := net.Listen("tcp", os.Args[1] + ":" + os.Args[2])
 	revisaError(err)
	listaConexiones = make([]net.Conn,1)

	for {
		conexion, err1 := servidor.Accept()
		listaConexiones = append(listaConexiones, conexion)
		revisaError(err1)
		go manejaCliente(conexion)
	}
}

func revisaError(err error) {
	if err != nil {
		fmt.Println("Fallo la conexion al servidor")
		os.Exit(1)
	}
}

func manejaCliente(conexion net.Conn) {
	defer conexion.Close()
	mensaje := make([]byte,256)
	for  {
		for  {
			cadena, err := conexion.Read(mensaje)
			if err != nill{
				if err == io.EOF {
					break
				}
			}
			mensaje = bytes.Trim(mensaje[:cadena], "\x00")
		}
		enviaClientes(mensaje, conexion)
		mensaje = make([]byte, 0)
	}
}

func enviaClientes(mensaje byte[], cliente net.Conn){
	for _, conexion := range listaConexiones {
		conexion.Write([]byte(mensaje))
	}
}
