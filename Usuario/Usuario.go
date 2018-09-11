  package Usuario

  import(
    "net"
  )

  var Usuarios map[string]net.Conn

  /* Implementa las operaciones de usuarios*/

  func InicializaUsuarios()  {
    Usuarios = make(map[string]net.Conn)
  }

  func ObtenerUsuarios() (map[string]net.Conn) {
    return Usuarios
  }

  func UsuarioRegistrado(conexionActual net.Conn) (bool){
    for _, conexion := range Usuarios {
      if(conexion == conexionActual){
        return true
      }
    }
    return false
  }

  func RegistraUsuarioNuevo(conexion net.Conn, mensaje []byte){
    cadena := string(mensaje)
		var nombre string
		for i := 0; i < len(cadena) ; i++ {
			 if string(cadena[i]) != " " {
				 nombre += string(cadena[i])
			 } else {
				 break
			 }
		}
		Usuarios[nombre] = conexion
	}
