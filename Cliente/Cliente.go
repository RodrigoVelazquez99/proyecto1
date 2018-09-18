	package main

	import (
		"bufio"
		"fmt"
		"net"
		"os"
		//"bytes"
		"io"
	)

	func main() {

		if len(os.Args) != 3 {
			fmt.Println("No agregaste correctamente los datos de la direccion y puerto")
			os.Exit(1)
		}

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
		/*var mensaje []byte
		buffer := make([]byte, 256)
		for  {
			for {
				cadena, err := conexion.Read(buffer)
				if err != nil {
					if err == io.EOF {
						break
					}
				}
				buffer = bytes.Trim(buffer[:cadena], "\x00")
				mensaje = append(mensaje, buffer...)
				if mensaje[len(mensaje)-1] == 10 {
					break
				}
			}
			fmt.Printf("%s\n",mensaje[:len(mensaje)-1])
			mensaje = make([]byte, 0)
		}*/
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

	func revisaError(err error) {
		if err != nil {
			fmt.Println("Fallo la conexion al servidor")
			os.Exit(1)
		}
	}
