  package Usuario

  import(
    "net"
    "strings"
  )

  /* Implementa las operaciones de usuarios */

  type Usuario struct{
    nombre string
    estado string
    sala string
    conexion net.Conn
  }

  var SalasChat map[string][]Usuario

  func InicializaUsuarios() {
    SalasChat = make(map[string][]Usuario)
  }

  func ObtenerUsuarios() (map[string][]Usuario) {
    return SalasChat
  }

  func ObtenerListaUsuarios() string {
    lista := "Usuarios conectados:      " + "\n"
      for _, usuarios := range SalasChat {
        for _, usuario := range usuarios{
            lista += usuario.nombre + "\n"
        }
      }
    return lista
  }

  func ObtenerUsuariosIdentificados() []Usuario {
    usuariosIdentificados := make([]Usuario, 0)
    for _, usuarios := range SalasChat {
      for _, usuario := range usuarios {
        usuariosIdentificados = append(usuariosIdentificados, usuario)
      }
    }
    return usuariosIdentificados
  }

  func BuscaUsuarioPorNombre(nombreActual string) Usuario{
    user := Usuario{nombre: "default",estado: "",sala:"", conexion: nil}
    for _, usuarios := range SalasChat {
      for _, usuario := range usuarios {
        if usuario.nombre == nombreActual {
          return usuario
        }
      }
    }
    return user
  }

  func BuscaUsuarioPorConexion(conexionActual net.Conn) Usuario {
    user := Usuario{nombre: "default",estado: "",sala:"", conexion: nil}
    for _, usuarios := range SalasChat {
      for _, usuario := range usuarios {
        if usuario.conexion == conexionActual {
          return usuario
        }
      }
    }
    return user
  }

  func BuscaUsuariosPorSala(conexionActual net.Conn) []Usuario {
    for _, usuarios := range SalasChat {
      for _, usuario := range usuarios {
        if usuario.conexion == conexionActual{
          return usuarios
        }
      }
    }
    return nil
  }

  func RegistraUsuarioNuevo(conexion net.Conn, nombre string, sala string){
    nuevoUsuario := Usuario{
      nombre: nombre,
      estado: "ACTIVE",
      sala: sala,
      conexion: conexion,
    }
    if salaRegistrada(sala) {
      tmp := SalasChat[sala]
      tmp = append(tmp, nuevoUsuario)
      SalasChat[sala] = tmp
    } else {
      listaUsuarios := make([]Usuario,0)
      listaUsuarios = append(listaUsuarios, nuevoUsuario)
      SalasChat[sala] = listaUsuarios
    }
	}


  func eliminaUsuario(user Usuario){
    for _, usuarios := range SalasChat {
      for i := 0 ; i < len(usuarios) ; i++ {
        if usuarios[i].nombre == user.nombre {
          usuarios[i] = usuarios[len(usuarios) - 1]
          usuarios = usuarios[:len(usuarios)-1]
          break
        }
      }
    }
  }

  func creaSala(conexion net.Conn, sala string)  {
    admin := BuscaUsuarioPorConexion(conexion)
    usuarios := make([]Usuario, 0)
    usuarios = append(usuarios, admin)
    SalasChat[sala] = usuarios
  }


  func salaRegistrada(salaRequerida string) bool {
    for sala, _ := range SalasChat {
      if sala == salaRequerida{
        return true
      }
    }
    return false
  }

  func getSala(salaRequerida string) []Usuario{
    for sala, usuarios := range SalasChat {
      if sala == salaRequerida {
        return usuarios
      }
    }
    return nil
  }

  func desconectaUsuario(conexion net.Conn)  {
    conexion.Write([]byte("Se ha desconectado"))
    eliminaUsuario(BuscaUsuarioPorConexion(conexion))
    conexion.Close()
  }

  func cambiaEstado(conexion net.Conn, estado string)  {
    usuario := BuscaUsuarioPorConexion(conexion)
    usuario.estado = estado
  }


  func IdentificaUsuario(entrada string, conexion net.Conn) ([]Usuario, string)  {
    palabras := separaPalabras(entrada)
    nombre := palabras[1]
    RegistraUsuarioNuevo(conexion,nombre,"SALA_GLOBAL")
    destinatarios := BuscaUsuariosPorSala(conexion)
    mensaje := nombre + " se ha conectado"
    return destinatarios, mensaje
  }

  func CapturaEstado(entrada string, conexion net.Conn) ([]Usuario, string) {
    destinatarios := BuscaUsuariosPorSala(conexion)
    palabras := separaPalabras(entrada)
    estado := palabras[1]
    usuario := BuscaUsuarioPorConexion(conexion)
    cambiaEstado(conexion, estado)
    mensaje := usuario.nombre  + " ha cambiado su estado a: " + estado
    return destinatarios, mensaje
  }

  func DevuelveUsuarios() string {
    usuarios := ObtenerListaUsuarios()
    return usuarios
  }

  func NuevaSala(entrada string, conexion net.Conn)  {
    palabras := separaPalabras(entrada)
    sala := palabras[1]
    creaSala(conexion, sala)
  }


  func Desconecta(conexion net.Conn) ([]Usuario, string) {
    usuario := BuscaUsuarioPorConexion(conexion)
    destinatarios := BuscaUsuariosPorSala(conexion)
    desconectaUsuario(conexion)
    mensaje := "Se ha desconectado " + usuario.nombre
    return destinatarios, mensaje
  }


func separaPalabras(entrada string) []string {
  palabra := ""
  palabras := make([]string, 0)
  for i := 0 ; i < len(entrada) ; i++ {
    if string(entrada[i]) == " "{
      continue
    }
    palabra += string(entrada[i])
    if i + 1 < len(entrada) && string(entrada[i + 1]) == " " {
      palabras = append(palabras,palabra)
      palabra = ""
    }
  }
  return palabras
 }


  func CapturaMensaje(entrada string) ([]Usuario, string) {
    destinatarios := make([]Usuario, 0)
    palabras := separaPalabras(entrada)
    destinatario := BuscaUsuarioPorNombre(palabras[1])
    mensaje := palabras[2]
    destinatarios = append(destinatarios, destinatario)
    return destinatarios,mensaje
  }

  func CapturaMensajePublico(entrada string) ([]Usuario, string) {
    destinatarios := make([]Usuario, 0)
    palabras := separaPalabras(entrada)
    mensaje := palabras[1]
    destinatarios = ObtenerUsuariosIdentificados()
    return destinatarios, mensaje
  }

  func InvitaUsuarios(entrada string) ([]Usuario, string) {
    destinatarios := make([]Usuario, 0)
    palabras := separaPalabras(entrada)
    sala := palabras[1]
    var usuario Usuario
    for i := 2; i < len(palabras); i++ {
      usuario = BuscaUsuarioPorNombre(palabras[i])
      destinatarios = append(destinatarios, usuario )
    }
    return destinatarios, sala
  }

  func AceptarSolicitud(entrada string)  {
    palabras := separaPalabras(entrada)
    sala := palabras[1]
    //  BUG:
  }

  func MensajeSala(entrada string) ([]Usuario, string) {
    palabras := separaPalabras(entrada)
    sala := palabras[1]
    mensaje := palabras[2]
    usuarios := getSala(sala)
    return usuarios, mensaje
  }

  func identificaBandera(conexion net.Conn, mensaje string) (string) {
    usuario := BuscaUsuarioPorConexion(conexion)
    if usuario.nombre == "default" {
      return "USUARIO_NO_IDENTIFICADO"
    }
  	if strings.Contains(mensaje, "IDENTIFY") {
  		return "IDENTIFY"
  	} else if strings.Contains(mensaje, "STATUS") {
  		return "STATUS"
  	} else if strings.Contains(mensaje, "USERS") {
  		return "USERS"
  	} else if strings.Contains(mensaje, "MESSAGE") {
  		return "MESSAGE"
  	} else if strings.Contains(mensaje, "PUBLICMESSAGE"){
  		return "PUBLICMESSAGE"
  	} else if strings.Contains(mensaje, "CREATEROOM"){
  		return "CREATEROOM"
  	} else if strings.Contains(mensaje, "INVITE"){
  		return "INVITE"
		} else if strings.Contains(mensaje, "JOINROOM") {
			return "JOINROOM"
		} else if strings.Contains(mensaje, "ROOMESSAGE") {
  		return "ROOMESSAGE"
  	} else if strings.Contains(mensaje, "DISCONNECT") {
  		return "DISCONNECT"
  	}
  	return "SIN_BANDERA"
  }
