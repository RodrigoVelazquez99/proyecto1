  package Usuario

  import(
    "net"
    //"strings"
  )

  /* Implementa las operaciones de usuarios*/

  type Usuario struct{
    nombre string
    estado string
    conexion net.Conn
  }

  var SalasChat map[string]Usuario

  func InicializaUsuarios() {
    SalasChat = make(map[string][]Usuario)
  }

  func ObtenerUsuarios() (map[string][]Usuario) {
    return SalasChat
  }

  func ObtenerListaUsuarios() string {
    lista := "Usuarios conectados:      " + "\n"
      for nombre, _ := range Usuarios {
        lista += "    " + nombre + "\n"
      }
    return lista
  }

  func BuscaUsuario(nombreActual string) net.Conn {
    for nombre, conexion := range Usuarios{
      if nombre == nombreActual{
        return conexion
      }
    }
    return nil
  }

  func BuscaUsuarioPorConexion(conexionActual net.Conn) string {
    for nombre, conexion := range Usuarios {
      if conexion == conexionActual {
        return nombre
      }
    }
    return "";
  }


  func UsuarioRegistrado(conexionActual net.Conn) bool{
    for _, conexion := range Usuarios {
      if(conexion == conexionActual){
        return true
      }
    }
    return false
  }

  func RegistraUsuarioNuevo(conexion net.Conn, nombre string){
		Usuarios[nombre] = conexion
	}

  func desconectaUsuario(nombre string){

  }
