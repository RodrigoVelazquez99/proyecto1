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
	direccionIP := os.Args[1]
	puerto := os.Args[2]
	var nombre string
	fmt.Println("Introduce el nombre de usuario")
	fmt.Scanln(&nombre)
	fmt.Println("Bienvenido" + nombre)
	//agregaUsuario(nombre)
	conexion, err := net.Dial("tcp", direccionIP+":"+puerto)
	if err != nil {
		err.Error()
	}
	for {
		informacionDevuelta := bufio.NewReader(os.Stdin)
		fmt.Print(nombre + ">")
		lector, _ := informacionDevuelta.ReadString('\n')
		fmt.Fprintf(conexion, lector+"\n")
		mensaje, _ := bufio.NewReader(conexion).ReadString('\n')
		fmt.Print(string(mensaje))
	}
}
