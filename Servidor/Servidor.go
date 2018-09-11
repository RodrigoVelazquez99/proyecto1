	package main

	import (
		"fmt"
		"net"
		"os"
		"io"
		"bytes"
		"github.com/RodrigoVelazquez99/proyecto1/Usuario"
	)

	func main() {

		if len(os.Args) != 3 {
			fmt.Println("No agregaste correctamente los datos de la direccion y puerto")
			os.Exit(1)
		}

		fmt.Println("  Iniciando el servidor .....  ")
		servidor, err := net.Listen("tcp", os.Args[1] + ":" + os.Args[2])
	 	revisaError(err)
		Usuario.InicializaUsuarios()
		for {
			conexion, disponible := servidor.Accept()
			revisaError(disponible)
			go manejaCliente(conexion)
			//fmt.Print(Usuario.ObtenerUsuarios())
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
				if !(Usuario.UsuarioRegistrado(conexion)) {
					Usuario.RegistraUsuarioNuevo(conexion, tmp)
				}
				if tmp[len(tmp) - 1 ] == 10 {
					break
				}
			}
			enviaClientes(tmp, conexion)
			tmp = make([]byte, 0)
		}
	}

	func enviaClientes(mensaje []byte, conexionActual net.Conn){
		for _, conexion := range Usuario.ObtenerUsuarios() {
			if conexion != nil{
				if conexion == conexionActual{
					continue
				} else {
					conexion.Write(mensaje)
				}
			}
		}
	}
