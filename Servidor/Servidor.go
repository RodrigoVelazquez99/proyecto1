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
			/*if !(Usuario.UsuarioRegistrado(conexion)) {
					solicitaNombre(conexion)
			}*/
			go manejaCliente(conexion)
		}

	}

	func revisaError(err error) {
		if err != nil {
			fmt.Println("Fallo la conexion al servidor")
			os.Exit(1)
		}
	}
/*
	func recibeNombre(conexion net.Conn) string {
		buffer := make([]byte,128)
		var tmp []byte
		for {
			nombre, err := conexion.Read(buffer)
			if err != nil{
				if err == io.EOF {
					break
				}
			}
			buffer = bytes.Trim(buffer[:nombre], "\x00")
			tmp = append(tmp, buffer...)
			if tmp[len(tmp) - 1 ] == 10 {
				break
			}
		}
		return string(tmp)
	}*/

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
			bandera, destinatario := opcionesMensaje(tmp)
			if(bandera != "sinBandera"){
			enviaMensaje(tmp, conexion)
			} else {
				conexion.Write(byte[]("Ingresa un comando valido \n "))
			}
			tmp = make([]byte, 0)
		}
	}

	func enviaMensaje(mensaje []byte, conexionActual net.Conn){
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
/*
	func solicitaNombre(conexion net.Conn) {
		conexion.Write([]byte(" Introduce el nombre de usuario \n "))
		nombre := recibeNombre(conexion)
		Usuario.RegistraUsuarioNuevo(conexion, nombre)
		nombre = "Bienvenido " + nombre
		conexion.Write([]byte(nombre))
	}*/

	func opcionesMensaje(arr []byte) (string, net.Conn, bool) {
		entrada := string(arr)
		bandera := identificaBandera(entrada)
		if()
		nombre, mensaje, destinatario, err := identificaDestinatario(bandera,entrada)

		return mensaje, destinatario
	}

	func capturaNombre(cadena string, lugar int) (string, string) {
		nombre := ""
		var mensaje string
		var tmp string
		for i := lugar; i < len(cadena); i++ {
			if(cadena[i] == " "){
				if nombre != ""{
				nombre = tmp
				continue
				} else {
				mensaje = tmp
				break
				}
			}
			tmp += cadena[i]
		}
		return nombre, mensaje
	}

	func identificaDestinatario(bandera string, entrada string) (string, string, net.Conn, err error) {
		var destinatario net.Conn
		var mensaje string
		var nombre string

		switch bandera {
		case "MESSAGE":
			nombre, mensaje = capturaNombre(entrada,8)
			destinatario = Usuario.BuscaUsuario(nombre)
			break
		case "PUBLICMESSAGE":
			nombre = capturaNombre(entrada, )
			destinatario = Usuario.BuscaUsuario(nombre)
			break
		case "CREATEROOM":
			nombre = capturaNombre(entrada, )
			destinatario = Usuario.BuscaUsuario(nombre)
			break
		case "INVITE":
			nombre = capturaNombre(entrada, )
			destinatario = Usuario.BuscaUsuario(nombre)
			break
		case "JOINROOM":
			nombre = capturaNombre(entrada, )
			destinatario = Usuario.BuscaUsuario(nombre)
			break
		case "ROOMESSAGE":
			nombre = capturaNombre(entrada, )
			destinatario = Usuario.BuscaUsuario(nombre)
			break
		default :
			return "","",nil,error
			break
		}
		return nombre, mensaje, nil, destinatario
	}

	func identificaBandera(mensaje string) string,string{
		if strings.Contains(mensaje, "IDENTIFY") {
			return "IDENTIFY", ""
		} else if strings.Contains(mensaje, "STATUS") {
			return "STATUS"
		} else if strings.Contains(mensaje, "USERS") {
			return "USERS"
		} else if strings.Contains(mensaje, "MESSAGE") {
			return "MESSAGE", ""
		} else if strings.Contains(mensaje, "PUBLICMESSAGE"){
			return "PUBLICMESSAGE"
		} else if strings.Contains(mensaje, "CREATEROOM"){
			return "CREATEROOM"
		} else if strings.Contains(mensaje, "INVITE"){
			return "INVITE"
		} else if strings.Contains(mensaje, "JOINROOM") {
			return "JOINROOM"
		} else if strings.Contains(mensaje, "ROOMESSAGE") {
			return "ROOMESSAGE"
		} else if strings.Contains(mensaje, "DISCONNECT") {
			return "DISCONNECT"
		}
		return "sinBandera"
	}
