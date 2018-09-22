	package main

	import (
		"fmt"
		"net"
		"os"
		"io"
		"bytes"
		"github.com/RodrigoVelazquez99/proyecto1/src/Usuario"
	)

	/**
	* Modelado y programacion - Proyecto 1: Chat
	*	Velázquez Cruz Rodrigo Fernando
	*/

	func main() {
		if len(os.Args) != 3 {
			fmt.Println("No agregaste correctamente los datos de la direccion IP y puerto")
			os.Exit(1)
		}
		fmt.Println("  Iniciando el servidor .....  ")
		servidor, err := net.Listen("tcp", os.Args[1] + ":" + os.Args[2])
        fmt.Println("  Conectado en: " + os.Args[1])
        fmt.Println("  Puerto: " + os.Args[2])
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

	/* Recibe los mensajes de los clientes e identifica las palabras clave */
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
			var sala string
			var destinatarios []net.Conn
			var salida string
			var tipo string
				switch bandera {
				case  "SIN_BANDERA":
					conexion.Write([]byte("[SERVIDOR] Ingresa un comando valido \n "))
					tmp = make([]byte, 0)
					continue
				case  "SIN_IDENTIFICAR":
					conexion.Write([]byte("[SERVIDOR] Usuario no identificado, identificate \n"))
					tmp = make([]byte, 0)
					continue
				case "IDENTIFY":
					nombre := Usuario.ObtenerNombre(conexion)
					if nombre != "SIN_IDENTIFICAR" {
							conexion.Write([]byte("[SERVIDOR] Ya estas identificado \n"))
							tmp = make([]byte, 0)
							continue
					} else {
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
					sala, destinatarios, salida = Usuario.CapturaMensaje(conexion, mensaje)
				case "PUBLICMESSAGE":
					tipo = "MENSAJE"
					sala, destinatarios, salida = Usuario.CapturaMensajePublico(conexion, mensaje)
				case "CREATEROOM":
					tipo = "NOTIFICACION"
					destinatarios, salida = Usuario.NuevaSala(mensaje, conexion)
			  case "INVITE":
					tipo = "NOTIFICACION"
					destinatarios, salida = Usuario.InvitaUsuarios(conexion, mensaje)
				case "JOINROOM":
					tipo = "NOTIFICACION"
					destinatarios, salida = Usuario.AceptarSolicitud(conexion, mensaje)
				case "ROOMESSAGE":
					tipo = "MENSAJE"
					sala, destinatarios, salida = Usuario.MensajeSala(conexion, mensaje)
				case "DISCONNECT":
					var permisos map[string]net.Conn
					tipo = "NOTIFICACION"
					destinatarios, permisos, salida = Usuario.Desconecta(conexion)
					if destinatarios == nil {
						tmp = make([]byte, 0)
						conexion.Close()
						continue
					}
				  go enviaRespuesta(sala, conexion, tipo, salida, destinatarios)
					enviaNuevosPermisos(permisos)
					break
				}
				go enviaRespuesta(sala, conexion, tipo, salida, destinatarios)
				tmp = make([]byte, 0)
			}
	}

	/* Envia el mensaje a los destinatarios */
	func enviaRespuesta(sala string, conexionActual net.Conn, tipo string, mensaje string, destinatarios []net.Conn){
		var cadena string
		nombre := Usuario.ObtenerNombre(conexionActual)
		for _, conexion := range destinatarios {
				if conexion == nil {
					continue
				}
					if tipo == "NOTIFICACION"{
						conexion.Write([]byte(mensaje + "\n"))
					} else {
						if sala == "" {
						conexion.Write([]byte(mensaje + "\n"))
						} else {
						cadena = sala + " - " + nombre + ": " + mensaje + "\n"
						conexion.Write([]byte(cadena))
						}
					}
		}
	}

	/* Cambia los permisos de salas de chat de  un usario eliminado a un usuario conectado */
	func enviaNuevosPermisos(nuevosPermisos map[string]net.Conn)  {
		for mensaje,conexion := range nuevosPermisos {
			conexion.Write([]byte(mensaje + "\n"))
		}
	}
