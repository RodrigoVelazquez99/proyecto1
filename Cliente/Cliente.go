	package main

	import (
		"bufio"
		"fmt"
		"net"
		"os"
		"strings"
	)
	var nombre string

	func main() {

		if len(os.Args) != 3 {
			fmt.Println("No agregaste correctamente los datos de la direccion y puerto")
			os.Exit(1)
		}
		fmt.Println(" Introduce el nombre de usuario ")
		fmt.Scanln(& nombre)
		fmt.Println("Bienvenido " + nombre)
		fmt.Println("				Comandos del chat" + "\n" +
										"/close		Cerrar sesion	  	" + "\n" +
										"/" + "\n" )
		var conexion net.Conn
		var err error
		for {
			conexion, err = net.Dial("tcp", os.Args[1] + ":" + os.Args[2])
			if err == nil {
				break
			}
		}
		defer conexion.Close()
		go recibeMensajes(conexion)
		enviaMensajes(conexion)

	}

	func enviaMensajes(conexion net.Conn){
		online := true
		for {
			for  online {
				conexion.Write([]byte(nombre + " se ha conectado" + "\n"))
				online = false
			}
			fmt.Print(nombre + ">")
			lector := bufio.NewReader(os.Stdin)
			mensaje, err := lector.ReadString('\n')
			if strings.Contains(mensaje, "/close") {
				fmt.Fprintf(conexion, nombre + " se ha desconectado " + "\n")
				break
			}
			if err != nil {
				break
			}
			mensaje = nombre + ":" + mensaje
			fmt.Fprintf(conexion, mensaje)
		}
	}

	func recibeMensajes(conexion net.Conn) {
		for {
			informacionDevuelta, err := bufio.NewReader(conexion).ReadString('\n')
			if err != nil {
				break
			}
			fmt.Print("\n" + string(informacionDevuelta))
		}
	}

	func revisaError(err error) {
		if err != nil {
			fmt.Println("Fallo la conexion al servidor")
			os.Exit(1)
		}
	}
