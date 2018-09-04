package main

import (
	"fmt"
	"io/ioutil"
	"net"
)

func main () {

	direccion, errl := net.ResolveTCPAddr("tcp", "localhost:8080")
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
