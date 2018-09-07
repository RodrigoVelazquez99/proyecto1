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
		conexion, disponible := servidor.Accept()
		listaConexiones = append(listaConexiones, conexion)
		if disponible != nil {
			fmt.Println("Conexion fallida")
			os.Exit(1)
		}
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
			cadena, err := conexion.Read(buffer)
			if err != nill{
				if err == io.EOF {
					log.Println("Ocurrio un error")
					os.Exit(1)
				}
				os.Exit(1)
			}
			mensaje = bytes.Trim(mensaje[:cadena], "\x00")
		}
		enviaClientes(mensaje, conexion)
	}
}

func enviaClientes(mensaje byte[], cliente net.Conn){

}
