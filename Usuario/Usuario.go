package Usuario

import "net"
import "fmt"

var listaUsuarios map[string]net.Conn

func InicializaUsuarios(){
  listaUsuarios = make(map[string]net.Conn)
}

func AgregaUsuario(nuevoNombre string, nuevaConexion net.Conn) {
  listaUsuarios[nuevoNombre] = nuevaConexion
  fmt.Println(listaUsuarios)
}

func BuscaUsuario(conexion net.Conn) string {
  for llave, valor := range listaUsuarios {
    if valor == conexion{
      return llave
    }
  }
  return ""
}

func GetUsuarios() map[string]net.Conn{
  return listaUsuarios
}
