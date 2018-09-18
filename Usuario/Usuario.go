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
    Clientes := make([]Usuario, 0)
    Usuarios = Clientes
  }

  func ObtenerListaUsuarios() string {
    lista := "Usuarios conectados:      " + "\n"
      for i:= 0 ;i < len(Usuarios) ; i++ {
          lista += Usuarios[i].estado + " " + Usuarios[i].nombre + "\n"
      }
    return lista
  }

  func ObtenerConexiones() []net.Conn {
    conexiones := make([]net.Conn, 0)
    for i:= 0; i < len(Usuarios); i++ {
      if Usuarios[i].conexion != nil {
        conexiones = append(conexiones, Usuarios[i].conexion)
      }
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
    return "SIN_IDENTIFICAR"
  }

  func ObtenerConexion(nombre string) net.Conn {
    for _,usuario := range Usuarios {
      if usuario.nombre == nombre {
        return usuario.conexion
      }
    }
    return nil
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
    user := Usuario{nombre: "SIN_IDENTIFICAR",estado: "",salas:nil, solicitudes:nil, conexion: nil}
      for _, usuario := range Usuarios {
        if usuario.nombre == nombreActual {
          return usuario
        }
      }
    return user
  }

  func BuscaUsuarioPorConexion(conexionActual net.Conn) Usuario {
    user := Usuario{nombre: "SIN_IDENTIFICAR",estado: "",salas:nil, solicitudes:nil,conexion: nil}
      for _, usuario := range Usuarios {
        if usuario.conexion == conexionActual {
          return usuario
        }
      }
    return user
  }

  func RegistraUsuarioNuevo(conexion net.Conn, nombre string, sala string){
    nuevaSala := make(map[string]string)
    nuevaSala[sala] = "MIEMBRO"
    solicitudes := make([]string, 0)
    nuevoUsuario := Usuario {
      nombre: nombre,
      estado: "ACTIVE",
      salas: nuevaSala,
      solicitudes: solicitudes,
      conexion: conexion,
    }
    Usuarios = append(Usuarios, nuevoUsuario)
	}


  func eliminaUsuario(usuario Usuario){
      for i := 0 ; i < len(Usuarios) ; i++ {
        if len(Usuarios) == 1 {
          nuevosUsuarios := make([]Usuario, 0)
          Usuarios = nuevosUsuarios
          break
        }
        if Usuarios[i].nombre == usuario.nombre {
          Usuarios[i] = Usuarios[len(Usuarios) - 1]
          Usuarios = Usuarios[:len(Usuarios)-1]
          break
        }
      }
  }

  func creaSala(conexion net.Conn, sala string)  {
    admin := BuscaUsuarioPorConexion(conexion)
    nuevaSala := admin.salas
    nuevaSala[sala] = "ADMIN"
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
        conexion := usuario.conexion
        if sala == salaRequerida {
          destinatarios = append(destinatarios, conexion)
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
    if len(palabras) <= 1 {
      destinatario := make([]net.Conn, 0)
      destinatario = append(destinatario, conexion)
      return destinatario, "No agregaste el nombre de usuario"
    }
    nombre := string(palabras[1])
    RegistraUsuarioNuevo(conexion,nombre,"SALA_GLOBAL")
    destinatarios := getSala("SALA_GLOBAL")
    mensaje := nombre + " se ha conectado" + "\n"
    return destinatarios, mensaje
  }

  func CapturaEstado(entrada string, conexion net.Conn) ([]net.Conn, string) {
    usuario := BuscaUsuarioPorConexion(conexion)
    destinatarios := getSala("SALA_GLOBAL")
    palabras := separaPalabras(entrada)
    if len(palabras) <= 1 {
      destinatario := make([]net.Conn, 0)
      destinatario = append(destinatario, conexion)
      return destinatario, "No ingresaste el estado"
    }
    switch palabras[1] {
    case "ACTIVE":
      cambiaEstado(conexion, "ACTIVE")
    case "BUSY":
      cambiaEstado(conexion, "BUSY")
    case "FAR":
      cambiaEstado(conexion, "FAR")
    default:
      destinatario := make([]net.Conn, 0)
      destinatario = append(destinatario, conexion)
      mensaje := palabras[1] + " no es un estado valido"
      return destinatario, mensaje
    }
    mensaje := "[SERVIDOR] " + usuario.nombre  + " ha cambiado su estado a: " + palabras[1] + "\n"
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
    mensaje := "Nueva sala: " + sala + "\n"
    destinatario := make([]net.Conn, 0)
    destinatario = append(destinatario, conexion)
    return destinatario, mensaje
  }


  func Desconecta(conexion net.Conn) ([]net.Conn, map[string]net.Conn, string) {
    eliminado := BuscaUsuarioPorConexion(conexion)
    nombre := eliminado.nombre
    cambiosPermisos := cambiaPermisos(eliminado)
    eliminaUsuario(eliminado)
    if len(Usuarios) == 0 {
      return nil,nil,""
    }
    destinatarios := ObtenerConexiones()
    mensaje := "Se ha desconectado " + nombre + "\n"
    return destinatarios, cambiosPermisos, mensaje
  }

  func pop(arr []string) string {
    if len(arr) == 1 {
      return string(arr[0])
    }
    eliminado := arr[0]
    arr[0] = arr[len(arr) - 1 ]
    arr = arr[:len(arr) -1 ]
    return string(eliminado)
  }

  func cambiaPermisos(eliminado Usuario) map[string]net.Conn {
    var destinatarios map[string]net.Conn
    salasAdmin := obtenerSalasDeUsuarioAdmin(eliminado)
    if salasAdmin != nil {
      destinatarios = make(map[string]net.Conn)
      j := 0
      for i := len(salasAdmin) - 1; i >= 0 ; i-- {
          if len(salasAdmin) == 1 {
            for j < len(Usuarios) {
              salaEliminado := pop(salasAdmin)
              if compartenSala(salaEliminado, Usuarios[j]) {
                cambiaPermiso(salaEliminado,Usuarios[j],"ADMIN")
                mensaje := "Has cambiado a permiso de administrador de la sala de chat: " + salaEliminado + "\n"
                destinatarios[mensaje] = Usuarios[i].conexion
                break
              } else {
                j++
              }
            }
          }
          salaEliminado := pop(salasAdmin)
          if compartenSala(salaEliminado, Usuarios[j]) {
            cambiaPermiso(salaEliminado,Usuarios[j],"ADMIN")
            mensaje := "Has cambiado a permiso de administrador de la sala de chat: " + salaEliminado + "\n"
            destinatarios[mensaje] = Usuarios[i].conexion
          } else {
            j++
          }
      }
    }
    return destinatarios
  }

  func compartenSala(salaRequerida string, usuario Usuario) bool {
    for sala,_ := range usuario.salas {
          if sala == salaRequerida {
            return true
          }
    }
    return false
  }

  func obtenerSalasDeUsuarioAdmin(usuario Usuario) []string{
    salasAdmin := make([]string, 0)
    for sala,permiso := range usuario.salas {
        if permiso == "ADMIN" {
          salasAdmin = append(salasAdmin, sala)
        }
    }
    return salasAdmin
  }

  func separaPalabras(entrada string) []string {
    palabras := strings.Fields(entrada)
    return palabras
   }

  func juntaPalabras(palabras []string, indice int) string {
    var linea  string
    for i := indice ; i < len(palabras) ; i++ {
      linea += palabras[i] + " "
    }
    return linea
  }


  func CapturaMensaje(conexion net.Conn, entrada string) (string,[]net.Conn, string) {
    destinatarios := make([]net.Conn, 0)
    palabras := separaPalabras(entrada)
    if len(palabras) < 3  {
      destinatarios = append(destinatarios, conexion)
      return "[SERVIDOR]", destinatarios, "No agregaste los campos requeridos"
    }
    destinatario := BuscaUsuarioPorNombre(palabras[1])
    if destinatario.nombre == "SIN_IDENTIFICAR" {
      destinatarios = append(destinatarios, conexion)
      return "[SERVIDOR]", destinatarios, "No esta identificado: " + palabras[1]
    }
    mensaje := juntaPalabras(palabras, 2)
    destinatarios = append(destinatarios, destinatario.conexion)
    return "PRIVATE MESSAGE", destinatarios, mensaje
  }

  func CapturaMensajePublico(conexion net.Conn, entrada string) (string, []net.Conn, string) {
    destinatarios := make([]net.Conn, 0)
    if len(Usuarios) == 0 {
      destinatarios = append(destinatarios, conexion)
    } else {
      destinatarios = ObtenerConexiones()
    }
    palabras := separaPalabras(entrada)
    mensaje := juntaPalabras(palabras, 1)
    return "PUBLIC MESSAGE", destinatarios, mensaje
  }

  func InvitaUsuarios(conexion net.Conn, entrada string) ([]net.Conn, string) {
    receptor := BuscaUsuarioPorConexion(conexion)
    palabras := separaPalabras(entrada)
    sala := palabras[1]
    permiso := obtenerPermisoDeSala(receptor, sala)
    destinatarios := make([]net.Conn, 0)
    if permiso != "ADMIN" {
      destinatarios = append(destinatarios, conexion)
      mensaje := "No tienes permiso para invitar usuarios" + "\n"
      return destinatarios, mensaje
    }
    for i := 2; i < len(palabras); i++ {
      usuario := BuscaUsuarioPorNombre(palabras[i])
      solicitudes := usuario.solicitudes
      solicitudes = append(solicitudes, sala)
      usuario.solicitudes = solicitudes
      destinatarios = append(destinatarios, usuario.conexion)
    }
    mensaje := "Invitacion para unirse a la sala: " + sala + "\n"
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
      mensaje = "No tienes solicitud para unirte a la sala" + "\n"
      return destinatarios, mensaje
    }
    mensaje = usuario.nombre + " ha ingresado a la sala" + "\n"
    destinatarios = getSala(sala)
    return destinatarios, mensaje
  }

  func MensajeSala(conexion net.Conn, entrada string) (string, []net.Conn, string) {
    palabras := separaPalabras(entrada)
    sala := palabras[1]
    mensaje := juntaPalabras(palabras, 2)
    usuario := BuscaUsuarioPorConexion(conexion)
    esMiembro := false
    for sala,_ := range usuario.salas {
      if sala == sala {
        esMiembro = true
      }
    }
    if !esMiembro {
      mensaje := "No eres miembro de la sala de chat" + "\n"
      destinatarios := make([]net.Conn, 0)
      destinatarios = append(destinatarios, conexion)
      return "", destinatarios, mensaje
    }
    destinatarios := getSala(sala)
    return sala, destinatarios, mensaje
  }

  func IdentificaBandera(conexion net.Conn, entrada []byte) (string) {
    mensaje := string(entrada)
    nombre := ObtenerNombre(conexion)
  	if strings.Contains(mensaje, "IDENTIFY") {
  		return "IDENTIFY"
  	}
    if nombre == "SIN_IDENTIFICAR" {
      return "SIN_IDENTIFICAR"
    }
    if strings.Contains(mensaje, "STATUS") {
  		return "STATUS"
  	}
    if strings.Contains(mensaje, "USERS") {
  		return "USERS"
  	}
    if strings.Contains(mensaje, "MESSAGE") {
  		return "MESSAGE"
  	}
    if strings.Contains(mensaje, "PUBLICMESSAGE") {
  		return "PUBLICMESSAGE"
  	}
    if strings.Contains(mensaje, "CREATEROOM") {
  		return "CREATEROOM"
  	}
    if strings.Contains(mensaje, "INVITE") {
  		return "INVITE"
		}
    if strings.Contains(mensaje, "JOINROOM") {
			return "JOINROOM"
		}
    if strings.Contains(mensaje, "ROOMESSAGE") {
  		return "ROOMESSAGE"
  	}
    if strings.Contains(mensaje, "DISCONNECT") {
  		return "DISCONNECT"
  	}
  	return "SIN_BANDERA"
  }
