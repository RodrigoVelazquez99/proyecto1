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
			fmt.Println("No agregaste correctamente los datos de la direccion IP y puerto")
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
			bandera := Usuario.IdentificaBandera(conexion, tmp)
			mensaje := string(tmp)
			var destinatarios []net.Conn
			var salida string
			var tipo string
				switch bandera {
				case  "SIN_BANDERA":
					conexion.Write([]byte("Ingresa un comando valido \n "))
					tmp = make([]byte, 0)
					continue
				case  "USUARIO_NO_IDENTIFICADO":
					conexion.Write([]byte("Usuario no identificado, identificate \n"))
					tmp = make([]byte, 0)
					continue
				case "IDENTIFY":
					usuario := Usuario.ObtenerNombre(conexion)
					if usuario != "DEFAULT" {
							conexion.Write([]byte("Ya existe un usuario identificado con ese nombre \n"))
							tmp = make([]byte, 0)
							continue
					} else{
						tipo = "NOTIFICACION"
						destinatarios, salida = Usuario.IdentificaUsuario(mensaje, conexion)
					}
				case "STATUS":
					tipo = "NOTIFICACION"
					destinatarios, salida = Usuario.CapturaEstado(mensaje, conexion)
				case "USERS":
					tipo = "NOTIFICACION"
					destinatarios, salida = Usuario.DevuelveUsuarios(conexion)
				case "MESSAGE":
					tipo = "MENSAJE"
					destinatarios, salida = Usuario.CapturaMensaje(mensaje)
				case "PUBLICMESSAGE":
					tipo = "MENSAJE"
					destinatarios, salida = Usuario.CapturaMensajePublico(mensaje)
				case "CREATEROOM":
					tipo = "NOTIFICACION"
					destinatarios, salida = Usuario.NuevaSala(mensaje, conexion)
					tmp = make([]byte, 0)
					continue
			  case "INVITE":
					tipo = "NOTIFICACION"
					destinatarios, salida = Usuario.InvitaUsuarios(conexion, mensaje)
				case "JOINROOM":
					tipo = "NOTIFICACION"
					destinatarios, salida = Usuario.AceptarSolicitud(conexion, mensaje)
				case "ROOMESSAGE":
					tipo = "MENSAJE"
					destinatarios, salida = Usuario.MensajeSala(mensaje)
				case "DISCONNECT":
					tipo = "NOTIFICACION"
					destinatarios, salida = Usuario.Desconecta(conexion)
				}
				go enviaRespuesta(tipo,salida, destinatarios)
				if bandera == "DISCONNECT"{
					conexion.Close()
				}
				tmp = make([]byte, 0)
			}
	}

	func enviaRespuesta(tipo string, mensaje string, destinatarios []net.Conn){
		var cadena string
		var nombre string
		for _, conexion := range destinatarios {
					nombre = Usuario.ObtenerNombre(conexion)
					if tipo == "NOTIFICACION"{
						conexion.Write([]byte(mensaje))
					} else {
						cadena = nombre + ": " + mensaje + "\n"
						conexion.Write([]byte(cadena))
					}
				}
	}
