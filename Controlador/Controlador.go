package Controlador

import (
  "github.com/mattn/go-gtk/glib"
  "github.com/mattn/go-gtk/gtk"
  //"os"
)

/* Interfaz grafica del usuario */

func IniciaInterfaz() {
    gtk.Init(nil)
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

    // Paneles
    
    frame1 := gtk.NewFrame("")
    framebox1 := gtk.NewVBox(false, 1)
    frame1.Add(framebox1)
    frame2 := gtk.NewFrame("")
    framebox2 := gtk.NewVBox(false, 1)
    frame2.Add(framebox2)

    panel.Pack1(frame2, true, true)
    panel.Pack2(frame1, false, false)

    // Barra de entrada de texto

    entry := gtk.NewEntry()
    entry.SetText("Escribe aqui")
    framebox1.Add(entry)

    //Botones

    botones := gtk.NewHBox(false, 1)
    boton := gtk.NewButtonWithLabel("Enviar")
    boton.Clicked(func() { boton.GetLabel() })
    botones.Add(boton)
    framebox1.PackStart(botones, false, false, 0)

    // Texto de entrada

    entradaMensajes := gtk.NewScrolledWindow(nil, nil)
    entradaMensajes.SetPolicy(gtk.POLICY_AUTOMATIC, gtk.POLICY_AUTOMATIC)
    entradaMensajes.SetShadowType(gtk.SHADOW_IN)
    texto := gtk.NewTextView()
    var iniciaLinea, acabaLinea gtk.TextIter
    buffer := texto.GetBuffer()
    buffer.GetStartIter(&iniciaLinea)
    buffer.Insert(&iniciaLinea, "Bienvenido")
    buffer.GetEndIter(&acabaLinea)
    entradaMensajes.Add(texto)
    framebox2.Add(entradaMensajes)

    ventana.Add(caratula)
    ventana.SetSizeRequest(600, 600)
    ventana.ShowAll()
    gtk.Main()
}
