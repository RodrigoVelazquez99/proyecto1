	package main

	import (
		"fmt"
		"net"
		"os"
		"github.com/RodrigoVelazquez99/proyecto1/src/Interfaz"
	)

	var conexion net.Conn

	func main() {
		if len(os.Args) != 3 {
	    fmt.Println("No agregaste correctamente los datos de la direccion y puerto")
	    os.Exit(1)
	  }
		var err error
		for {
			conexion, err = net.Dial("tcp", os.Args[1] + ":" + os.Args[2])
			if err == nil {
				break
			}
		}
		go Interfaz.IniciaInterfaz(conexion)
		defer conexion.Close()
		EnviaMensajes()
	}

	func EnviaMensajes(){
		for {
			mensaje := Interfaz.ObtenerMensaje()
			if mensaje != "" {
				conexion.Write([]byte(mensaje + "\n"))
			}
		}
	}
