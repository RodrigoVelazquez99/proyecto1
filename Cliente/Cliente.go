	package main

	import (
		"bufio"
		"fmt"
		"net"
		"os"
		"io"
		"github.com/RodrigoVelazquez99/proyecto1/Controlador"
	)

	func main() {

		if len(os.Args) != 3 {
			fmt.Println("No agregaste correctamente los datos de la direccion y puerto")
			os.Exit(1)
		}
		Controlador.IniciaInterfaz()

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
		for {
			lector := bufio.NewReader(os.Stdin)
			mensaje, err := lector.ReadString('\n')
			if err != nil {
				break
			}
			conexion.Write([]byte(mensaje))
		}
	}

	func recibeMensajes(conexion net.Conn) {
		for {
	    mensaje, err := bufio.NewReader(conexion).ReadString('\n')
			if err != nil {
				if err == io.EOF {
					break
				}
			}
			fmt.Print(mensaje + "\n")
		}
	}
