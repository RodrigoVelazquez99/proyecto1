package main

import (
	"fmt"
	"net"
	"os"
	"io"
	"bytes"
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
	var mensaje []byte
	buffer := make([]byte,256)
	for  {
		for  {
			cadena, err := conexion.Read(buffer)
			if err != nil{
				if err == io.EOF {
					break
				}
			}
			buffer = bytes.Trim(buffer[:cadena], "\x00")
			mensaje = append(mensaje,buffer...)
			if mensaje[len(mensaje)-1] == 10 {
				break
			}
		}
		enviaMensajes(mensaje, conexion)
		mensaje = make([]byte, 0)
	}
}

func enviaMensajes(mensaje []byte, cliente net.Conn){
	for _, conexion := range listaConexiones {
		conexion.Write(mensaje)
	}
}
