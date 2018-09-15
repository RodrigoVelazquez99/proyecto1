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
		entrada := make([]byte,256)
		var tmp []byte
		for {
			for {
				cadena, err := conexion.Read(entrada)
				if err != nil{
					if err == io.EOF {
						break
					}
				}
				entrada = bytes.Trim(entrada[:cadena], "\x00")
				tmp = append(tmp, entrada...)
				if tmp[len(tmp) - 1 ] == 10 {
					break
				}
			}
			bandera := Usuario.identificaBandera(conexion, tmp)
			mensaje := string(tmp)
				switch bandera {
				case  "SIN_BANDERA":
					conexion.Write(byte[]("Ingresa un comando valido \n "))
				case  "USUARIO_NO_IDENTIFICADO":
					conexion.Write(byte[]("Usuario no identificado, identificate \n"))
				case "IDENTIFY":
					destinatarios, salida := Usuario.IdentificaUsuario()
				case "STATUS":
					destinatarios, salida := Usuario.CapturaEstado(mensaje, conexion)
				case "USERS":
					salida := Usuario.DevuelveUsuarios()
				case "MESSAGE":
					destinatarios, salida := Usuario.CapturaMensaje(mensaje)
				case "PUBLICMESSAGE":
					destinatarios, salida := Usuario.CapturaMensajePublico(mensaje)
				case "CREATEROOM":
					Usuario.NuevaSala(mensaje, conexion)
			  case "INVITE":
					destinatarios, salida := Usuario.InvitaUsuarios(mensaje)
				case "JOINROOM":
					Usuario.AceptarSolicitud(mensaje)
				case "ROOMESSAGE":
					destinatarios, salida := Usuario.MensajeSala(mensaje)
				case "DISCONNECT":
					destinatarios, salida := Usuario.Desconecta(conexion)
					conexion.Close()
				}
				go enviaMensaje([]byte(mensaje), conexion, destinatarios)
				tmp = make([]byte, 0)
			}

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
