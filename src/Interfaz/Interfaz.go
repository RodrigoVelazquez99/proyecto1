package Interfaz

import (
	"bufio"
	"errors"
	"io"
	"net"

	"github.com/gotk3/gotk3/gtk"
)

var mensajes []string

const (
	path = "../GUI/GUI.glade"
)

/* Interfaz grafica del usuario */

func IniciaInterfaz(conexion net.Conn) {
	gtk.Init(nil)
	builder, err := build(path)
	if err != nil {
		print("el path")
		panic(err)
	}
	mensajes = make([]string, 0)
	ventana, err := window(builder, "window")
	if err != nil {
		panic(err)
	}
	ventana.SetTitle("Chat")
	ventana.Connect("destroy", func() {
		println("\n Cliente desconectado")
		gtk.MainQuit()
	}, "")
	entrada, err := entry(builder)
	if err != nil {
		panic(err)
	}
	entrada.SetEditable(true)
	entrada.SetText("Escribe aqui")
	boton, err := button(builder, "button1")
	if err != nil {
		panic(err)
	}
	boton.Connect("clicked", func() {
		mensaje, _ := entrada.GetText()
		mensajes = append(mensajes, mensaje)
		entrada.SetText("")
	})
	entradaMensajes, err := scrolledWindow(builder)
	if err != nil {
		panic(err)
	}
	entradaMensajes.SetPolicy(gtk.POLICY_AUTOMATIC, gtk.POLICY_AUTOMATIC)
	entradaMensajes.SetShadowType(gtk.SHADOW_IN)
	texto, err := textView(builder)
	if err != nil {
		panic(err)
	}
  texto.SetEditable(false)
  texto.SetOverwrite(false)
	buffer, err := texto.GetBuffer()
	if err != nil {
		panic(err)
	}
	iniciaLinea := buffer.GetStartIter()
	acabaLinea := buffer.GetEndIter()
	go RecibeMensajes(buffer, iniciaLinea, acabaLinea, conexion)
	ventana.SetDefaultSize(600, 600)
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

func limpiaMensajes() {
	nuevaLista := make([]string, 0)
	mensajes = nuevaLista
}

func RecibeMensajes(buffer *gtk.TextBuffer, inicia *gtk.TextIter, acaba *gtk.TextIter, conexion net.Conn) {
	for {
		mensaje, err := bufio.NewReader(conexion).ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
		}
		buffer.Insert(acaba, mensaje)
	}
}

func build(ruta string) (*gtk.Builder, error) {
	builder, err := gtk.BuilderNew()
	if err != nil {
		return nil, err
	}
	if ruta != "" {
		err = builder.AddFromFile(ruta)
		if err != nil {
			return nil, errors.New("Error")
		}
	}
	return builder, nil
}

func window(builder *gtk.Builder, tipo string) (*gtk.Window, error) {
	object, err := builder.GetObject(tipo)
	if err != nil {
		return nil, err
	}
	ventana, ok := object.(*gtk.Window)
	if !ok {
		return nil, err
	}
	return ventana, nil
}

func textView(builder *gtk.Builder) (*gtk.TextView, error) {
	object, err := builder.GetObject("textview1")
	if err != nil {
		return nil, err
	}
	text, ok := object.(*gtk.TextView)
	if !ok {
		return nil, err
	}
	return text, nil
}

func scrolledWindow(builder *gtk.Builder) (*gtk.ScrolledWindow, error) {
	object, err := builder.GetObject("scrolledwindow2")
	if err != nil {
		return nil, err
	}
	ventana, ok := object.(*gtk.ScrolledWindow)
	if !ok {
		return nil, err
	}
	return ventana, nil
}

func button(builder *gtk.Builder, tipo string) (*gtk.Button, error) {
	object, err := builder.GetObject(tipo)
	if err != nil {
		panic(err)
	}
	boton, ok := object.(*gtk.Button)
	if !ok {
		return nil, err
	}
	return boton, nil
}

func grid(builder *gtk.Builder) (*gtk.Grid, error) {
	object, err := builder.GetObject("grid1")
	if err != nil {
		panic(err)
	}
	grid, ok := object.(*gtk.Grid)
	if !ok {
		return nil, err
	}
	return grid, nil
}

func entry(builder *gtk.Builder) (*gtk.Entry, error) {
	object, err := builder.GetObject("entry1")
	if err != nil {
		panic(err)
	}
	entry, ok := object.(*gtk.Entry)
	if !ok {
		return nil, err
	}
	return entry, nil
}

func creaColumna(nombre string, id int) *gtk.TreeViewColumn {
	cellRenderer, err := gtk.CellRendererTextNew()
	if err != nil {
		panic(err)
	}
	columna, err := gtk.TreeViewColumnNewWithAttribute(nombre, cellRenderer, "text", id)
	if err != nil {
		panic(err)
	}
	return columna
}

/*
func creaTreeView() (*gtk.TreeView, *gtk.ListStore) {
	treeView, err := gtk.TreeViewNew()
	if err != nil {
		panic(err)
	}
	treeView.AppendColumn(creaColumna("Title", COLUMN_TITLE))
	treeView.AppendColumn(creaColumna("Performer", COLUMN_PERFORMER))
	treeView.AppendColumn(creaColumna("Album", COLUMN_ALBUM))
	treeView.AppendColumn(creaColumna("Genre", COLUMN_GENRE))
	Paths := creaColumna("Path", COLUMN_PATH)
	Paths.SetVisible(false)
	treeView.AppendColumn(Paths)
	listStore, err := gtk.ListStoreNew(glib.TYPE_STRING, glib.TYPE_STRING, glib.TYPE_STRING, glib.TYPE_STRING, glib.TYPE_STRING)
	if err != nil {
		panic(err)
	}
	treeView.SetModel(listStore)
	return treeView, listStore
}

func actualizaLista(listStore *gtk.ListStore, canciones []string) {
	i := 0
	for i < len(canciones) {
		nuevoRenglon(listStore, canciones[i], canciones[i+1], canciones[i+2], canciones[i+3], canciones[i+4])
		i += 5
	}
}

func nuevoRenglon(listStore *gtk.ListStore, titulo string, interprete string, album string, genero string, ruta string) {
	iter := listStore.Append()
	err := listStore.Set(iter,
		[]int{COLUMN_TITLE, COLUMN_PERFORMER, COLUMN_ALBUM, COLUMN_GENRE, COLUMN_PATH},
		[]interface{}{titulo, interprete, album, genero, ruta})
	if err != nil {
		panic(err)
	}
}
*/
