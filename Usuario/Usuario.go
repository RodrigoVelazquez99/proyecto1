  package Usuario

  import(
    "net"
    "strings"
  )

  /* Implementa las operaciones de usuarios */

  type Usuario struct{
    nombre string
    estado string
    salas map[string]string
    solicitudes []string
    conexion net.Conn
  }

  var Usuarios []Usuario

  func InicializaUsuarios() {
    Usuarios = make([]Usuario,0)
  }

  func ObtenerUsuarios() ([]Usuario) {
    return Usuarios
  }

  func ObtenerListaUsuarios() string {
    lista := "Usuarios conectados:      " + "\n"
      for _, usuario := range Usuarios{
          lista += usuario.nombre + "\n"
      }
    return lista
  }

  func ObtenerUsuariosIdentificados() []Usuario {
    usuariosIdentificados := make([]Usuario, 0)
    for _, usuario := range Usuarios {
      usuariosIdentificados = append(usuariosIdentificados, usuario)
    }
    return usuariosIdentificados
  }

  func ObtenerConexiones(usuarios []Usuario) []net.Conn{
    conexiones := make([]net.Conn, 0)
    for _,usuario := range Usuarios {
      conexiones = append(conexiones, usuario.conexion)
    }
    return conexiones
  }

  func UsuarioIdentificado(conexion net.Conn) bool {
      for _,usuario := range Usuarios {
        if usuario.conexion == conexion {
          return true
        }
      }
    return false
  }

  func ObtenerNombre(conexion net.Conn) string {
      for _,usuario := range Usuarios {
        if usuario.conexion == conexion {
          return usuario.nombre
        }
      }
    return "DEFAULT"
  }

  func obtenerPermisoDeSala(usuario Usuario, sala string) string {
    for sala,permiso := range usuario.salas {
      if sala == sala {
        return permiso
      }
    }
    return "NO_EXISTE_LA_SALA"
  }

  func BuscaUsuarioPorNombre(nombreActual string) Usuario{
    user := Usuario{nombre: "default",estado: "",salas:nil, solicitudes:nil, conexion: nil}
      for _, usuario := range Usuarios {
        if usuario.nombre == nombreActual {
          return usuario
        }
      }
    return user
  }

  func BuscaUsuarioPorConexion(conexionActual net.Conn) Usuario {
    user := Usuario{nombre: "default",estado: "",salas:nil, solicitudes:nil,conexion: nil}
      for _, usuario := range Usuarios {
        if usuario.conexion == conexionActual {
          return usuario
        }
      }
    return user
  }

  func RegistraUsuarioNuevo(conexion net.Conn, nombre string, sala string){
    nuevaSala := make(map[string]string, 0)
    solicitudes := make([]string, 0)
    nuevoUsuario := Usuario {
      nombre: nombre,
      estado: "ACTIVE",
      salas: nuevaSala,
      solicitudes: solicitudes,
      conexion: conexion,
    }
    if !salaRegistrada(sala) {
      cambiaPermiso(sala, nuevoUsuario, "admin")
    }
    Usuarios = append(Usuarios, nuevoUsuario)
	}


  func eliminaUsuario(user Usuario){
      for i := 0 ; i < len(Usuarios) ; i++ {
        if Usuarios[i].nombre == user.nombre {
          Usuarios[i] = Usuarios[len(Usuarios) - 1]
          Usuarios = Usuarios[:len(Usuarios)-1]
          break
        }
      }
  }

  func creaSala(conexion net.Conn, sala string)  {
    admin := BuscaUsuarioPorConexion(conexion)
    nuevaSala := admin.salas
    nuevaSala[sala] = "admin"
    admin.salas = nuevaSala
  }

  func salaRegistrada(salaRequerida string) bool {
    for _, usuario := range Usuarios {
      for sala,_ := range usuario.salas {
        if sala == salaRequerida {
          return true
        }
      }
    }
    return false
  }

  func getSala(salaRequerida string) []net.Conn {
    destinatarios := make([]net.Conn, 0)
    for _, usuario := range Usuarios {
      for sala,_ := range usuario.salas {
        if sala == salaRequerida {
          destinatarios = append(destinatarios, usuario.conexion)
        }
      }
    }
    return destinatarios
  }

  func cambiaEstado(conexion net.Conn, estado string)  {
    usuario := BuscaUsuarioPorConexion(conexion)
    usuario.estado = estado
  }

  func cambiaPermiso(sala string, usuario Usuario, permiso string)  {
    nuevaSala := usuario.salas
    nuevaSala[sala] = permiso
    usuario.salas = nuevaSala
  }

  func IdentificaUsuario(entrada string, conexion net.Conn) ([]net.Conn, string)  {
    palabras := separaPalabras(entrada)
    nombre := palabras[1]
    RegistraUsuarioNuevo(conexion,nombre,"SALA_GLOBAL")
    destinatarios := getSala("SALA_GLOBAL")
    mensaje := nombre + " se ha conectado"
    return destinatarios, mensaje
  }

  func CapturaEstado(entrada string, conexion net.Conn) ([]net.Conn, string) {
    usuario := BuscaUsuarioPorConexion(conexion)
    destinatarios := getSala("SALA_GLOBAL")
    palabras := separaPalabras(entrada)
    estado := palabras[1]
    cambiaEstado(conexion, estado)
    mensaje := usuario.nombre  + " ha cambiado su estado a: " + estado
    return destinatarios, mensaje
  }

  func DevuelveUsuarios(conexion net.Conn) ([]net.Conn, string) {
    destinatarios := make([]net.Conn, 0)
    destinatarios = append(destinatarios, conexion)
    mensaje := ObtenerListaUsuarios()
    return destinatarios, mensaje
  }

  func NuevaSala(entrada string, conexion net.Conn) ([]net.Conn, string) {
    palabras := separaPalabras(entrada)
    sala := palabras[1]
    creaSala(conexion, sala)
    mensaje := "Nueva sala: " + sala
    destinatario := make([]net.Conn, 0)
    destinatario = append(destinatario, conexion)
    return destinatario, mensaje
  }


  func Desconecta(conexion net.Conn) ([]net.Conn, string) {
    usuario := BuscaUsuarioPorConexion(conexion)
    destinatarios := ObtenerConexiones(Usuarios)
    nombre := usuario.nombre
    eliminaUsuario(usuario)
    mensaje := "Se ha desconectado " + nombre
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


  func CapturaMensaje(entrada string) ([]net.Conn, string) {
    destinatarios := make([]net.Conn, 0)
    palabras := separaPalabras(entrada)
    usuario := BuscaUsuarioPorNombre(palabras[1])
    mensaje := palabras[2]
    destinatarios = append(destinatarios, usuario.conexion)
    return destinatarios,mensaje
  }

  func CapturaMensajePublico(entrada string) ([]net.Conn, string) {
    destinatarios := getSala("SALA_GLOBAL")
    palabras := separaPalabras(entrada)
    mensaje := palabras[1]
    return destinatarios, mensaje
  }

  func InvitaUsuarios(conexion net.Conn, entrada string) ([]net.Conn, string) {
    receptor := BuscaUsuarioPorConexion(conexion)
    palabras := separaPalabras(entrada)
    sala := palabras[1]
    permiso := obtenerPermisoDeSala(receptor, sala)
    destinatarios := make([]net.Conn, 0)
    if permiso != "admin" {
      destinatarios = append(destinatarios, conexion)
      mensaje := "No tienes permiso para invitar usuarios"
      return destinatarios, mensaje
    }
    for i := 2; i < len(palabras); i++ {
      usuario := BuscaUsuarioPorNombre(palabras[i])
      solicitudes := usuario.solicitudes
      solicitudes = append(solicitudes, sala)
      usuario.solicitudes = solicitudes
      destinatarios = append(destinatarios, usuario.conexion)
    }
    mensaje := "Invitacion enviada"
    return destinatarios, mensaje
  }

  func AceptarSolicitud(conexion net.Conn, entrada string) ([]net.Conn, string) {
    var mensaje string
    usuario := BuscaUsuarioPorConexion(conexion)
    destinatarios := make([]net.Conn, 0)
    invitado := false
    palabras := separaPalabras(entrada)
    sala := palabras[1]
    for _,solicitud := range usuario.solicitudes {
        if solicitud == sala {
          invitado = true
        }
    }
    if !invitado {
      destinatarios = append(destinatarios, conexion)
      mensaje = "No tienes solicitud para unirte a la sala"
      return destinatarios, mensaje
    }
    mensaje = usuario.nombre + " ha ingresado a la sala"
    destinatarios = getSala(sala)
    return destinatarios, mensaje
  }

  func MensajeSala(entrada string) ([]net.Conn, string) {
    palabras := separaPalabras(entrada)
    sala := palabras[1]
    mensaje := palabras[2]
    destinatarios := getSala(sala)
    return destinatarios, mensaje
  }

  func IdentificaBandera(conexion net.Conn, entrada []byte) (string) {
    mensaje := string(entrada)
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
