  package Usuario

  import(
    "net"
    "strings"
  )

  /* Implementa las operaciones de usuarios */

  type Usuario struct{
    nombre string
    estado string
    salas [string]string
    solicitudes []string
    conexion net.Conn
  }

  var Usuarios []Usuario

  func InicializaUsuarios() {
    Usuarios = make([]Usuarios,0)
  }

  func ObtenerUsuarios() (map[string][]Usuario) {
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
    for _, usuario := range usuarios {
      usuariosIdentificados = append(usuariosIdentificados, usuario)
    }
    return usuariosIdentificados
  }

  func ObtenerConexiones(usuarios []Usuario) []net.Conn{
    conexiones := make([]net.Conn, 0)
    for _,usuario := range usuarios {
      conexiones = append(conexiones, usuario.conexion)
    }
    return conexiones
  }

  func UsuarioIdentificado(conexion net.Conn) bool {
      for _,usuario := range usuarios {
        if usuario.conexion == conexion {
          return true
        }
      }
    return false
  }

  func ObtenerNombre(conexion net.Conn) string {
      for _,usuario := range usuarios {
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
    user := Usuario{nombre: "default",estado: "",salas:nil, permiso:nil, conexion: nil}
      for _, usuario := range usuarios {
        if usuario.nombre == nombreActual {
          return usuario
        }
      }
    return user
  }

  func BuscaUsuarioPorConexion(conexionActual net.Conn) Usuario {
    user := Usuario{nombre: "default",estado: "",salas:nil, permiso:nil, conexion: nil}
      for _, usuario := range usuarios {
        if usuario.conexion == conexionActual {
          return usuario
        }
      }
    return user
  }

  func BuscaUsuariosPorSala(conexionActual net.Conn) []Usuario {
      for _, usuario := range usuarios {
        if usuario.conexion == conexionActual{
          return usuarios
        }
      }
    return nil
  }

  func RegistraUsuarioNuevo(conexion net.Conn, nombre string, sala string){
    nuevaSala := make([]string, 0)
    nuevaSala = append(nuevaSala, sala)
    permisos := make(map[string]string, 0)
    permisos[nuevaSala] = "miembro"
    nuevoUsuario := Usuario {
      nombre: nombre,
      estado: "ACTIVE",
      salas: nuevaSala,
      permiso: permisos,
      conexion: conexion,
    }
    if salaRegistrada(sala) {
      tmp := SalasChat[sala]
      tmp = append(tmp, nuevoUsuario)
      SalasChat[sala] = tmp
    } else {
      listaUsuarios := make([]Usuario,0)
      nuevoUsuario.permiso = "admin"
      listaUsuarios = append(listaUsuarios, nuevoUsuario)
      SalasChat[sala] = listaUsuarios
    }
	}


  func eliminaUsuario(user Usuario){
      for i := 0 ; i < len(usuarios) ; i++ {
        if usuarios[i].nombre == user.nombre {
          usuarios[i] = usuarios[len(usuarios) - 1]
          usuarios = usuarios[:len(usuarios)-1]
          break
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
    for _, usuario := range usuarios {
      for sala,_ := range usuario.salas {
        if sala == salaRequerida {
          return true
        }
      }
    }
    return false
  }

  func getSala(salaRequerida string) []Usuario {
    destinatarios := make([]Usuario, 0)
    for _, usuario := range usuarios {
      for sala,_ := range usuario.salas {
        if sala == salaRequerida {
          destinatarios = append(destinatarios, usuario)
        }
      }
    }
    return destinatarios
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

  func cambiaPermiso(sala string, usuario Usuario, permiso string)  {
    nuevaSala := usuario.salas
    nuevaSala[sala] = permiso
    usuario.salas = nuevaSala
  }

  func IdentificaUsuario(entrada string, conexion net.Conn) ([]net.Conn, string)  {
    palabras := separaPalabras(entrada)
    nombre := palabras[1]
    RegistraUsuarioNuevo(conexion,nombre,"SALA_GLOBAL")
    destinatarios := ObtenerConexiones(BuscaUsuariosPorSala(conexion))
    mensaje := nombre + " se ha conectado"
    return destinatarios, mensaje
  }

  func CapturaEstado(entrada string, conexion net.Conn) ([]net.Conn, string) {
    destinatarios := ObtenerConexiones(BuscaUsuariosPorSala(conexion))
    palabras := separaPalabras(entrada)
    estado := palabras[1]
    usuario := BuscaUsuarioPorConexion(conexion)
    cambiaEstado(conexion, estado)
    mensaje := usuario.nombre  + " ha cambiado su estado a: " + estado
    return destinatarios, mensaje
  }

  func DevuelveUsuarios(conexion net.Conn) ([]net.Conn, string) {
    usuarios := make([]Usuario, 0)
    usuarios = append(usuarios, BuscaUsuarioPorConexion(conexion))
    destinatarios := ObtenerConexiones(usuarios)
    mensaje := ObtenerListaUsuarios()
    return destinatarios, mensaje
  }

  func NuevaSala(entrada string, conexion net.Conn) ([]net.Conn, string) {
    palabras := separaPalabras(entrada)
    sala := palabras[1]
    creaSala(conexion, sala)
    mensaje := "Nueva sala: " + sala
    usuario := BuscaUsuarioPorConexion(conexion)
    nuevaSala := usuario.salas
    nuevaSala[sala] = "admin"
    usuario.salas = nuevaSala
    destinatario := make([]net.Conn, 0)
    destinatario = append(destinatario, conexion)
    return destinatario, mensaje
  }


  func Desconecta(conexion net.Conn) ([]net.Conn, string) {
    usuario := BuscaUsuarioPorConexion(conexion)
    destinatarios := ObtenerConexiones(BuscaUsuariosPorSala(conexion))
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


  func CapturaMensaje(entrada string) ([]net.Conn, string) {
    usuarios := make([]Usuario, 0)
    palabras := separaPalabras(entrada)
    destinatario := BuscaUsuarioPorNombre(palabras[1])
    mensaje := palabras[2]
    usuarios = append(usuarios, destinatario)
    destinatarios := ObtenerConexiones(usuarios)
    return destinatarios,mensaje
  }

  func CapturaMensajePublico(entrada string) ([]net.Conn, string) {
    usuarios := make([]Usuario, 0)
    palabras := separaPalabras(entrada)
    mensaje := palabras[1]
    destinatarios := ObtenerConexiones(usuarios)
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
      return destinatario, mensaje
    }
    for i := 2; i < len(palabras); i++ {
      usuario := BuscaUsuarioPorNombre(palabras[i])
      destinatarios := append(destinatarios, usuario.conexion)
    }
    mensaje = "Invitacion enviada"
    return destinatarios, mensaje
  }

  func AceptarSolicitud(conexion net.Conn, entrada string) ([]net.Conn, string) {
    var mensaje string
    var destinatarios []net.Conn
    usuarios := make([]Usuario, 0)
    invitado := false
    usuario := BuscaUsuarioPorConexion(conexion)
    palabras := separaPalabras(entrada)
    sala := palabras[1]
    for _,invitacion := range usuario.salas {
        if invitacion == sala{
          invitado = true
        }
    }
    if !invitado {
      usuarios = append(usuarios, usuario)
      mensaje = "No tienes solicitud para unirte a la sala"
      destinatarios = ObtenerConexiones(usuarios)
      return destinatarios, mensaje
    }
    mensaje = usuario.nombre + " ha ingresado a la sala"
    usuarios = getSala(sala)
    destinatarios = ObtenerConexiones(usuarios)
    return destinatarios, mensaje
  }

  func MensajeSala(entrada string) ([]net.Conn, string) {
    palabras := separaPalabras(entrada)
    sala := palabras[1]
    mensaje := palabras[2]
    usuarios := getSala(sala)
    destinatarios := ObtenerConexiones(usuarios)
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
