	package main

	import (
		"fmt"
		"net"
		"os"
		"io"
		"bytes"
		"github.com/RodrigoVelazquez99/proyecto1/Usuario"
		"strings"
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
			bandera, tipo, destinatarios, texto := Usuario.opcionesMensaje(tmp, conexion)
			if bandera == "SIN_BANDERA" {
				conexion.Write(byte[]("Ingresa un comando valido \n "))
			} else if bandera == "USUARIO_NO_IDENTIFICADO" {
				conexion.Write(byte[]("Usuario no identificado, identificate \n"))
			} else {
				if tipo == "SOLICITUD" {
					respuesta := Usuario.manejaSolicitud(bandera, conexion, tmp)
					enviaMensaje([]byte(respuesta),)
				} else {
					enviaMensaje([]byte(texto), conexion, destinatarios)
				}
			}
			tmp = make([]byte, 0)
		}
	}

	func enviaMensaje(mensaje []byte, conexionActual net.Conn, destinatarios []Usuario){
		for _, conexion := range Usuario.ObtenerUsuarios() {
			if conexion != nil{
				if conexion == conexionActual{
					continue
				} else {
					cadena := string(mensaje)
					cadena = Usuario.BuscaUsuarioPorConexion(conexionActual) + ": " + cadena + "\n"
					conexion.Write([]byte(cadena))
				}
			}
		}
	}
