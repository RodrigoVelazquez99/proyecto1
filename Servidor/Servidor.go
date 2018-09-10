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
		listaConexiones = make([]net.Conn, 1)

		for {
			conexion, disponible := servidor.Accept()
			revisaError(disponible)
			if conexion != nil {
				listaConexiones = append(listaConexiones, conexion)
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
		var tmp []byte
		for {
			for {
				cadena, err := conexion.Read(mensaje)
				if err != nil{
					if err == io.EOF {
						break
					}
				}
				mensaje = bytes.Trim(mensaje[:cadena], "\x00")
				tmp = append(tmp, mensaje...)
				if tmp[len(tmp) - 1 ] == 10 {
					break
				}
			}
			enviaClientes(tmp, conexion)
			tmp = make([]byte, 0)
		}
	}

	func enviaClientes(mensaje []byte, conexionActual net.Conn){
		for _, conexion := range listaConexiones {
			if conexion != nil{
				if conexion == conexionActual{
					continue
				} else {
					conexion.Write(mensaje)
				}
			}
		}
	}
