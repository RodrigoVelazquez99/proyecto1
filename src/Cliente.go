package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("No agregaste correctamente los datos del puerto y direccion")
		os.Exit(1)
	}
	direccionIP := os.Args[1]
	puerto := os.Args[2]
	var nuevoUsuario Usuario
	fmt.Println("Introduce el nombre de usuario")
	fmt.Scanln(&nuevoUsuario.nombre)
	direccion, errl := net.ResolveTCPAddr("tcp", direccionIP+":"+puerto)
	if errl != nil {
		errl.Error()
	}
	user, err := net.DialTCP("tcp", nil, direccion)
	if err != nil {
		err.Error()
	}
	mensaje, _ := ioutil.ReadAll(user)
	fmt.Println(string(mensaje))
}
