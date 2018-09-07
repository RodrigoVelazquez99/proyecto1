package Usuario

import "net"

type User struct{
  nombre string
  conexion net.Conn
}

var listaUsuarios []User

func GetUsuarios() []User{
  return listaUsuarios
}

func IncializaUsuarios(){
  listaUsuarios = make([]User, 1)
}

func AgregaUsuario(nuevoNombre string, nuevaConexion net.Conn) {
  nuevoUsuario := User{ nombre: nuevoNombre, conexion: nuevaConexion }
  listaUsuarios = append(listaUsuarios, nuevoUsuario)
}
