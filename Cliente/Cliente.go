package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {

	if len(os.Args) != 3 {
		fmt.Println("No agregaste correctamente los datos de la direccion y puerto")
		os.Exit(1)
	}

	var nombre string

	fmt.Println(" Introduce el nombre de usuario ")
	fmt.Scanln(& nombre)
	fmt.Println("Bienvenido " + nombre)
	var conexion net.Conn
	var err error
	for {
		conexion, err = net.Dial("tcp", os.Args[1] + ":" + os.Args[2])
		if err == nil{
			break
		}
	}

	defer conexion.Close()
	go recibeMensajes(conexion)
	enviaMensajes(conexion, nombre)

}

func enviaMensajes(conexion net.Conn, nombre string){
	for {
		fmt.Print(nombre + ">")
		mensaje, _ := bufio.NewReader(conexion).ReadString('\n')
		mensaje = nombre + ":" + mensaje
		conexion.Write([]byte(mensaje))
	}
}

func recibeMensajes(conexion net.Conn) {
	for {
		informacionDevuelta := bufio.NewReader(os.Stdin)
		lector, _ := informacionDevuelta.ReadString('\n')
		fmt.Fprintf(conexion, lector + "\n")
	}
}


func revisaError(err error) {
	if err != nil {
		fmt.Println("Fallo la conexion al servidor")
		os.Exit(1)
	}
}
