package Interfaz

import (
  "github.com/mattn/go-gtk/glib"
  "github.com/mattn/go-gtk/gtk"
)

var mensajes []string
var salida []string

/* Interfaz grafica del usuario */

func IniciaInterfaz() {
    gtk.Init(nil)
    mensajes = make([]string, 0)
    salida = make([]string, 0)
    ventana := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
    ventana.SetPosition(gtk.WIN_POS_CENTER)
    ventana.SetTitle("Chat")
    ventana.Connect("destroy", func(ctx *glib.CallbackContext) {
        println("\n Cliente desconectado", ctx.Data().(string))
        gtk.MainQuit()
    },"")
    caratula := gtk.NewVBox(false, 1)
    panel := gtk.NewVPaned()
    caratula.Add(panel)
    frame1 := gtk.NewFrame("")
    framebox1 := gtk.NewVBox(false, 1)
    frame1.Add(framebox1)
    frame2 := gtk.NewFrame("")
    framebox2 := gtk.NewVBox(false, 1)
    frame2.Add(framebox2)
    panel.Pack1(frame2, true, true)
    panel.Pack2(frame1, false, false)
    entrada := gtk.NewEntry()
    entrada.SetEditable(true)
    entrada.SetText("Escribe aqui")
    framebox1.Add(entrada)
    entrada.SetEditable(true)
    botones := gtk.NewHBox(false, 1)
    boton := gtk.NewButtonWithLabel("Enviar")
    boton.Connect("clicked", func () {
        mensajes = append(mensajes, entrada.GetText())
    })
    botones.Add(boton)
    framebox1.PackStart(botones, false, false, 0)
    entradaMensajes := gtk.NewScrolledWindow(nil, nil)
    entradaMensajes.SetPolicy(gtk.POLICY_AUTOMATIC, gtk.POLICY_AUTOMATIC)
    entradaMensajes.SetShadowType(gtk.SHADOW_IN)
    texto := gtk.NewTextView()
    var iniciaLinea, acabaLinea gtk.TextIter
    buffer := texto.GetBuffer()
    buffer.GetStartIter(&iniciaLinea)
      if len(salida) != 0 {
        msg := ObtenerSalida()
        buffer.Insert(&iniciaLinea, msg)
      }
    buffer.GetEndIter(&acabaLinea)
    entradaMensajes.Add(texto)
    framebox2.Add(entradaMensajes)
    ventana.Add(caratula)
    ventana.SetSizeRequest(600, 600)
    ventana.ShowAll()
    gtk.Main()
}

func ObtenerMensaje() string {
  if len(mensajes) == 0 {
    return ""
  }
  mensaje := mensajes[0]
  if len(mensajes) == 1 {
    limpiaMensajes()
    return mensaje
  }
  mensajes = mensajes[1:]
  return mensaje
}

func ObtenerSalida() string {
  if len(salida) == 0 {
    return ""
  }
  mensaje := salida[0]
  if len(salida) == 1 {
    limpiaSalida()
    return mensaje
  }
  salida = salida[1:]
  return mensaje
}

func ImprimirMensaje(mensaje string) {
  salida = append(salida, mensaje)
}

func limpiaMensajes()  {
  nuevaLista := make([]string, 0)
  mensajes = nuevaLista
}

func limpiaSalida()  {
  nuevaSalida := make([]string, 0)
  salida = nuevaSalida
}
