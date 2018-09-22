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
    Usuarios = make([]Usuario, 0)
  }

  func ObtenerListaUsuarios() string {
    lista := "[SERVIDOR] Usuarios conectados: "
      for i:= 0 ;i < len(Usuarios) ; i++ {
          lista += Usuarios[i].nombre + " "
      }
    return lista
  }

  /* Regresa las conexiones a los sockets de todos los usuarios conectados */
  func ObtenerConexiones() []net.Conn {
    conexiones := make([]net.Conn, 0)
    for i:= 0; i < len(Usuarios); i++ {
      if Usuarios[i].conexion != nil {
        conexiones = append(conexiones, Usuarios[i].conexion)
      }
    }
    return conexiones
  }

  /* Busca un usuario por nombre */
  func ObtenerNombre(conexion net.Conn) string {
      for _,usuario := range Usuarios {
        if usuario.conexion == conexion {
          return usuario.nombre
        }
      }
    return "SIN_IDENTIFICAR"
  }

  /* Busca un Usuario por conexion del socket */
  func ObtenerConexion(nombre string) net.Conn {
    for _,usuario := range Usuarios {
      if usuario.nombre == nombre {
        return usuario.conexion
      }
    }
    return nil
  }

  /* Obtiene el permiso del usuario en la sala */
  func obtenerPermisoDeSala(usuario Usuario, salaRequerida string) string {
    for sala,permiso := range usuario.salas {
      if sala == salaRequerida {
        return permiso
      }
    }
    return "NO_EXISTE_LA_SALA"
  }

  /* Regresa las conexiones de los sockets de los miembros de una sala */
  func obtenerMiembrosSala(salaRequerida string) []net.Conn {
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

  /* Revisa que la conexion del socket pertenezca a algun usuario */
  func UsuarioIdentificado(conexion net.Conn) bool {
    for _,usuario := range Usuarios {
      if usuario.conexion == conexion {
        return true
      }
    }
    return false
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

  func limpiaListaUsuarios()  {
    nuevaLista := make([]Usuario, 0)
    Usuarios = nuevaLista
  }

  func eliminaUsuario(usuario Usuario){
    if len(Usuarios) == 1 {
      limpiaListaUsuarios()
      return
    }
    for i := 0 ; i < len(Usuarios) ; i++ {
      if Usuarios[i].nombre == usuario.nombre {
        Usuarios[i] = Usuarios[len(Usuarios) - 1]
        Usuarios = Usuarios[:len(Usuarios)-1]
        break
      }
    }
  }

  func registraUsuarioSala(usuario Usuario, salaRequerida string)  {
      for i := 0; i < len(Usuarios); i++ {
        if Usuarios[i].nombre == usuario.nombre {
          salasInscrito := Usuarios[i].salas
          salasInscrito[salaRequerida] = "MIEMBRO"
          Usuarios[i].salas = salasInscrito
          break
        }
      }
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

  func cambiaEstado(conexionActual net.Conn, estado string)  {
    for  i:= 0; i < len(Usuarios) ; i++ {
      if Usuarios[i].conexion == conexionActual {
        Usuarios[i].estado = estado
        break
      }
    }
  }

  func cambiaPermiso(sala string, usuario Usuario, permiso string)  {
    for i := 0;  i < len(Usuarios) ; i++ {
      if Usuarios[i].nombre == usuario.nombre  {
        nuevoPermiso := Usuarios[i].salas
        nuevoPermiso[sala] = permiso
        Usuarios[i].salas = nuevoPermiso
        break
      }
    }
  }

  func IdentificaUsuario(entrada string, conexion net.Conn) ([]net.Conn, string)  {
    palabras := separaPalabras(entrada)
    if len(palabras) <= 1 {
      destinatario := make([]net.Conn, 0)
      destinatario = append(destinatario, conexion)
      return destinatario, "[SERVIDOR] No agregaste el nombre de usuario"
    }
    nombre := string(palabras[1])
    RegistraUsuarioNuevo(conexion,nombre,"SALA_GLOBAL")
    destinatarios := obtenerMiembrosSala("SALA_GLOBAL")
    mensaje := "[SERVIDOR] " + nombre + " se ha conectado" + "\n"
    return destinatarios, mensaje
  }

  func DevuelveUsuarios(conexion net.Conn) ([]net.Conn, string) {
    destinatarios := make([]net.Conn, 0)
    destinatarios = append(destinatarios, conexion)
    mensaje := ObtenerListaUsuarios()
    return destinatarios, mensaje
  }

  func NuevaSala(entrada string, conexion net.Conn) ([]net.Conn, string) {
    destinatario := make([]net.Conn, 0)
    destinatario = append(destinatario, conexion)
    palabras := separaPalabras(entrada)
    if len(palabras) == 1 {
      return destinatario, "[SERVIDOR] No ingresaste el nombre de la sala"
    }
    nuevaSala := palabras[1]
    if salaRegistrada(nuevaSala) {
      return destinatario, "[SERVIDOR] Ya existe una sala con ese nombre, elige otro"
    }
    for i := 0; i < len(Usuarios) ; i++ {
      if Usuarios[i].conexion == conexion {
        salas := Usuarios[i].salas
        salas[nuevaSala] = "ADMIN"
        Usuarios[i].salas = salas
        break
      }
    }
    mensaje := "[SERVIDOR] Nueva sala: " + nuevaSala + "\n"
    return destinatario, mensaje
  }

  func Desconecta(conexion net.Conn) ([]net.Conn, map[string]net.Conn, string) {
    eliminado := BuscaUsuarioPorConexion(conexion)
    nombre := eliminado.nombre
    cambiosPermisos := cambiaPermisos(eliminado)
    eliminaUsuario(eliminado)
    destinatarios := ObtenerConexiones()
    mensaje := "[SERVIDOR] Se ha desconectado " + nombre + "\n"
    return destinatarios, cambiosPermisos, mensaje
  }

  func cambiaPermisos(eliminado Usuario) map[string]net.Conn {
    var mensaje string
    destinatarios := make(map[string]net.Conn, 0)
    salasAdmin := obtenerSalasDeUsuarioAdmin(eliminado)
    if salasAdmin != nil{
      j := 0
      for i := 0; i < len(Usuarios); i++ {
        if Usuarios[i].nombre != eliminado.nombre {
          if j < len(salasAdmin) && perteneceSala(salasAdmin[j], Usuarios[i]) {
            cambiaPermiso(salasAdmin[j], Usuarios[i],"ADMIN")
            mensaje = "[SERVIDOR] Has cambiado a permiso de administrador de la sala de chat: " + salasAdmin[j]
            destinatarios[mensaje] = Usuarios[i].conexion
            j++
          }
        }
      }
    }
    return destinatarios
  }

  func perteneceSala(salaRequerida string, usuario Usuario) bool {
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

  func CapturaEstado(entrada string, conexion net.Conn) ([]net.Conn, string) {
    usuario := BuscaUsuarioPorConexion(conexion)
    destinatarios := obtenerMiembrosSala("SALA_GLOBAL")
    palabras := separaPalabras(entrada)
    if len(palabras) <= 1 {
      destinatario := make([]net.Conn, 0)
      destinatario = append(destinatario, conexion)
      return destinatario, "[SERVIDOR] No ingresaste el estado"
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
      mensaje := "[SERVIDOR] " + palabras[1] + " no es un estado valido"
      return destinatario, mensaje
    }
    mensaje := "[SERVIDOR] " + usuario.nombre  + " ha cambiado su estado a: " + palabras[1] + "\n"
    return destinatarios, mensaje
  }

  func CapturaMensaje(conexion net.Conn, entrada string) (string,[]net.Conn, string) {
    destinatarios := make([]net.Conn, 0)
    palabras := separaPalabras(entrada)
    if len(palabras) < 3  {
      destinatarios = append(destinatarios, conexion)
      return "", destinatarios, "[SERVIDOR] No agregaste los campos requeridos"
    }
    destinatario := BuscaUsuarioPorNombre(palabras[1])
    if destinatario.nombre == "SIN_IDENTIFICAR" {
      destinatarios = append(destinatarios, conexion)
      return "", destinatarios, " [SERVIDOR] No esta identificado: " + palabras[1]
    }
    mensaje := juntaPalabras(palabras, 2)
    destinatarios = append(destinatarios, destinatario.conexion)
    return "PRIVATE MESSAGE", destinatarios, mensaje
  }

  func CapturaMensajePublico(conexion net.Conn, entrada string) (string, []net.Conn, string) {
    destinatarios := make([]net.Conn, 0)
    if len(Usuarios) == 1 {
      destinatarios = append(destinatarios, conexion)
    } else {
      destinatarios = ObtenerConexiones()
    }
    palabras := separaPalabras(entrada)
    mensaje := juntaPalabras(palabras, 1)
    return "PUBLIC MESSAGE", destinatarios, mensaje
  }

  func InvitaUsuarios(conexion net.Conn, entrada string) ([]net.Conn, string) {
    destinatarios := make([]net.Conn, 0)
    receptor := BuscaUsuarioPorConexion(conexion)
    palabras := separaPalabras(entrada)
    if len(palabras) == 1 {
      destinatarios = append(destinatarios, conexion)
      return destinatarios, "[SERVIDOR] Escribe la sala de chat"
    }
    sala := palabras[1]
    if !salaRegistrada(sala) {
      destinatarios = append(destinatarios, conexion)
      return destinatarios, "[SERVIDOR] La sala no existe"
    }
    permiso := obtenerPermisoDeSala(receptor, sala)
    if permiso != "ADMIN" {
      destinatarios = append(destinatarios, conexion)
      return destinatarios, "[SERVIDOR] No tienes permiso para invitar usuarios" + "\n"
    }
    for i := 2 ; i < len(palabras) ; i++ {
      usuario := BuscaUsuarioPorNombre(palabras[i])
      if usuario.nombre == "SIN_IDENTIFICAR" {
        continue
      }
      enviaSolicitud(usuario, sala)
      destinatarios = append(destinatarios, usuario.conexion)
    }
    mensaje := "[SERVIDOR] " + receptor.nombre + " te ha invitado a unirse a la sala: " + sala + "\n"
    return destinatarios, mensaje
  }

  func enviaSolicitud(usuario Usuario, solicitud string) {
    for i := 0 ; i < len(Usuarios) ; i++{
      if Usuarios[i].nombre == usuario.nombre {
        nuevaSolicitud := Usuarios[i].solicitudes
        nuevaSolicitud = append(nuevaSolicitud, solicitud)
        Usuarios[i].solicitudes = nuevaSolicitud
        break
      }
    }
  }

  /* Acepta la solicitud para unirse a la sala */
  func AceptarSolicitud(conexion net.Conn, entrada string) ([]net.Conn, string) {
    var mensaje string
    destinatarios := make([]net.Conn, 0)
    usuario := BuscaUsuarioPorConexion(conexion)
    invitado := false
    palabras := separaPalabras(entrada)
    if len(palabras) == 1 {
      destinatarios = append(destinatarios, conexion)
      return destinatarios, "[SERVIDOR] Ingresa el nombre de la sala"
    }
    sala := palabras[1]
    if !salaRegistrada(sala) {
      destinatarios = append(destinatarios, conexion)
      return destinatarios, "[SERVIDOR] La sala no existe"
    }
    for _,solicitud := range usuario.solicitudes {
        if solicitud == sala {
          invitado = true
          break
        }
    }
    if !invitado {
      destinatarios = append(destinatarios, conexion)
      mensaje = "[SERVIDOR] No tienes solicitud para unirte a la sala" + "\n"
      return destinatarios, mensaje
    }
    registraUsuarioSala(usuario, sala)
    mensaje = "[SERVIDOR] " + usuario.nombre + " ha ingresado a la sala: " + sala  + "\n"
    destinatarios = obtenerMiembrosSala(sala)
    return destinatarios, mensaje
  }

  /* Envia el mensaje a los miembros de una sala */
  func MensajeSala(conexion net.Conn, entrada string) (string, []net.Conn, string) {
    destinatarios := make([]net.Conn, 0)
    palabras := separaPalabras(entrada)
    if len(palabras) == 1 {
      destinatarios = append(destinatarios, conexion)
      return "", destinatarios, "[SERVIDOR] No ingresaste el nombre de la sala"
    }
    salaDestino := palabras[1]
    if !salaRegistrada(salaDestino) {
      destinatarios = append(destinatarios, conexion)
      return "", destinatarios, "[SERVIDOR] La sala no existe"
    }
    mensaje := juntaPalabras(palabras, 2)
    usuario := BuscaUsuarioPorConexion(conexion)
    esMiembro := false
    for sala,_ := range usuario.salas {
      if sala == salaDestino {
        esMiembro = true
      }
    }
    if !esMiembro {
      mensaje := "[SERVIDOR] No eres miembro de la sala de chat" + "\n"
      destinatarios = append(destinatarios, conexion)
      return "", destinatarios, mensaje
    }
    destinatarios = obtenerMiembrosSala(salaDestino)
    return salaDestino, destinatarios, mensaje
  }

  /* Identica la bandera en el mensaje*/
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
    if strings.Contains(mensaje, "PUBLICMESSAGE") {
  		return "PUBLICMESSAGE"
  	}
    if strings.Contains(mensaje, "ROOMESSAGE") {
      return "ROOMESSAGE"
    }
    if strings.Contains(mensaje, "MESSAGE") {
  		return "MESSAGE"
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
    if strings.Contains(mensaje, "DISCONNECT") {
  		return "DISCONNECT"
  	}
  	return "SIN_BANDERA"
  }
