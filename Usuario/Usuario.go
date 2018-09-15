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

  func CreaSala(conexion net.Conn, sala string)  {
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



  /**
  * Metodo que
  * @return bandera la bandera del mensaje
  * @return destinatarios la lista de destinatarios
  * @return mensaje el mensaje a enviar
  */
  func opcionesMensaje(arr []byte, conexionActual net.Conn) (string, string, []Usuario, string) {
  	entrada := string(arr)
  	bandera, tipo := identificaBandera(conexionActual, entrada)
    switch bandera {
    case "USUARIO_NO_IDENTIFICADO":
      return "USUARIO_NO_IDENTIFICADO", tipo, nil, entrada
    case "SIN_BANDERA":
      return "SIN_BANDERA", tipo, nil, entrada
    }

    if tipo == "SOLICITUD" {
      return bandera, tipo, nil, entrada
    }

    destinatarios, mensaje := identificaDestinatario(bandera,entrada, conexionActual)
    return bandera, tipo, destinatarios, mensaje
  }

  func capturaNombres(bandera string, entrada string) (string, []net.Conn) {
    conexiones := make([]net.Conn, 0)
    mensaje := string(entrada)
    var nombre string
    var sala string
    if bandera == "IDENTIFY" {
      for i := 9; i < len(mensaje) ; i++ {
        nombre += string(mensaje[i])
      }
    } else if bandera == "INVITE" {
      sala := ""
      tmp := ""
      j := 7
      for j < len(mensaje) {
         tmp += string(mensaje[j])
         j++
        if string(mensaje[j]) == " " && string(mensaje[j]) != "\n" {
          if sala != "" {
            nombre = tmp
            usuario := BuscaUsuarioPorNombre(nombre)
            conexiones = append(conexiones, usuario.conexion)
            tmp = ""
            j++
          } else {
            sala = tmp
            tmp = ""
            j++
          }
        }
      }
    }
    return sala, conexiones
  }

  func cambiaEstado(conexion net.Conn, estado string)  {
    usuario := BuscaUsuarioPorConexion(conexion)
    usuario.estado = estado
  }

  func manejaSolicitud(bandera string, conexion net.Conn, entrada string) (string, []Usuario) {
    switch bandera {
    case "IDENTIFY":
      _, nombres := capturaNombres(bandera, entrada)
      nombre := nombres[0]
      usuario := BuscaUsuarioPorConexion(nombre)
      RegistraUsuarioNuevo(conexion,usuario.nombre,"SALA_GLOBAL")
      vecinos := BuscaUsuariosPorSala(conexion)
      mensaje := usuario.nombre + " se ha conectado"
      return mensaje, vecinos
    case "STATUS":
      vecinos := BuscaUsuariosPorSala(conexion)
      palabras := separaPalabras(entrada)
      estado := palabras[1]
      usuario := BuscaUsuarioPorConexion(conexion)
      cambiaEstado(conexion, estado)
      mensaje := usuario.nombre  + " ha cambiado su estado a: " + estado
      return mensaje, vecinos
    case "USERS":
      usuarios := ObtenerListaUsuarios()
      return usuarios, nil
    case "CREATEROOM":
      palabras := separaPalabras(entrada)
      sala := palabras[1]
      CreaSala(conexion, sala)
      break
    case "DISCONNECT":
      usuario := BuscaUsuarioPorConexion(conexion)
      vecinos := BuscaUsuariosPorSala(conexion)
      desconectaUsuario(conexion)
      mensaje := "Se ha desconectado " + usuario.nombre
      return mensaje, vecinos
    }
    return "", nil
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

  /***
  *@return destinatarios los usuarios que reciben el mensaje
  *@return mensaje el mensaje
  */
  func capturaMensaje(entrada string) ([]Usuario, string) {
    destinatarios := make([]Usuario, 0)
    palabras := separaPalabras(entrada)
    destinatario := BuscaUsuarioPorNombre(palabras[1])
    mensaje := palabras[2]
    destinatarios = append(destinatarios, destinatario)
    return destinatarios,mensaje
  }

  func capturaMensajePublico(entrada string) ([]Usuario, string) {
    destinatarios := make([]Usuario, 0)
    palabras := separaPalabras(entrada)
    mensaje := palabras[1]
    destinatarios = ObtenerUsuariosIdentificados()
    return destinatarios, mensaje
  }

  func invitaUsuarios(entrada string) (string, []Usuario) {
    destinatarios := make([]Usuario, 0)
    palabras := separaPalabras(entrada)
    sala := palabras[1]
    var usuario Usuario
    for i := 2; i < len(palabras); i++ {
      usuario = BuscaUsuarioPorNombre(palabras[i])
      destinatarios = append(destinatarios, usuario )
    }
    return sala, destinatarios
  }

  func aceptarSolicitud(entrada string)  {
    palabras := separaPalabras(entrada)
    sala := palabras[1]
  }

  func mensajeSala(entrada string) ([]Usuario, string) {
    palabras := separaPalabras(entrada)
    sala := palabras[1]
    mensaje := palabras[2]
    usuarios := getSala(sala)
    return usuarios, mensaje
  }


	func identificaDestinatario(bandera string, entrada string, conexionActual net.Conn) ([]Usuario, string) {
  	var destinatarios []Usuario
    var sala string
    var mensaje string
  	switch bandera {
  	case "MESSAGE":
      destinatarios, mensaje = capturaMensaje(entrada)
  		return destinatarios, mensaje
  	case "PUBLICMESSAGE":
      destinatarios, mensaje = capturaMensajePublico(entrada)
  		return destinatarios, mensaje
  	case "INVITE":
      sala, destinatarios = invitaUsuarios(entrada)
			break
		case "JOINROOM":
      aceptarSolicitud(entrada)
  		break
  	case "ROOMESSAGE":
      destinatarios, mensaje = mensajeSala(entrada)
  		break
  	default :
  		return nil,""
   	}
  		return  destinatarios, mensaje
  	}

  func identificaBandera(conexion net.Conn, mensaje string) (string,string) {
    usuario := BuscaUsuarioPorConexion(conexion)
    if usuario.nombre == "default" {
      return "USUARIO_NO_IDENTIFICADO", ""
    }
  	if strings.Contains(mensaje, "IDENTIFY") {
  		return "IDENTIFY", "SOLICITUD"
  	} else if strings.Contains(mensaje, "STATUS") {
  		return "STATUS", "SOLICITUD"
  	} else if strings.Contains(mensaje, "USERS") {
  		return "USERS", "SOLICITUD"
  	} else if strings.Contains(mensaje, "MESSAGE") {
  		return "MESSAGE", "ENVIO"
  	} else if strings.Contains(mensaje, "PUBLICMESSAGE"){
  		return "PUBLICMESSAGE", "ENVIO"
  	} else if strings.Contains(mensaje, "CREATEROOM"){
  		return "CREATEROOM", "SOLICITUD"
  	} else if strings.Contains(mensaje, "INVITE"){
  		return "INVITE", "ENVIO"
		} else if strings.Contains(mensaje, "JOINROOM") {
			return "JOINROOM", "ENVIO"
		} else if strings.Contains(mensaje, "ROOMESSAGE") {
  		return "ROOMESSAGE", "ENVIO"
  	} else if strings.Contains(mensaje, "DISCONNECT") {
  		return "DISCONNECT", "SOLICITUD"
  	}
  	return "SIN_BANDERA", ""
  }
