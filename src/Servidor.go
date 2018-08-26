package main

import (
	"fmt"
	"net"
	"os"
)

/* Controla el serivdor */

func creaServidor() {
	direccion, conexion := net.ResolveTCPAddr("tcp", "192.168.0.16")
	if conexion != nil {
		fmt.Fprintln("Ocurrio un error en el servidor")
		os.Exit(1)
	}
}
