# proyecto1

Chat programado en Go

Modelado y Programacion

Go version 1.10.1

# Autor

Velázquez Cruz Rodrigo Fernando

Numero de cuenta UNAM: 315254565

## Dependencias

* libgtk-3-dev
* libglib2.0-dev
* libcairo2-dev
* GTK+ 3.20

Estas bibliotecas van en el directorio $GOPATH

* github.com/mattn/go-gtk/glib
* github.com/mattn/go-gtk/gtk

Instalacion :

```bash
$ sudo apt-get install libgtk-3-dev libcairo2-dev libglib2.0-dev
```

```bash
$ go get github.com/mattn/go-gtk/gtk
```

```bash
$ go get github.com/mattn/go-gtk/glib
```

[GTK+3 3.20](ftp.gnome.org/pub/gnome/sources/gtk+/3.20/gtk+-3.20.0.tar.xz)

## Levantar el servidor

```bash
$ go run Servidor.go <ip> <puerto>
```

## Conectar un cliente

```bash
$ go run Cliente.go <ip> <puerto>
```

## Compilacion

```bash
$ go build
```

### Errores conocidos

Faltan pruebas unitarias.
