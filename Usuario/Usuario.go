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

func buscaUsuario(conexion net.Conn) string {
  for _,usuario := range var {
    if conexion == usuario.conexion{
      return usuario.nombre
    }
  }
  return ""
}
