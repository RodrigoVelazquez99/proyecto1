# proyecto1
Chat programado en Go

Modelado y Progrmacion

Velázquez Cruz Rodrigo Fernando

Go version 1.10.1

## Dependencias

* libgtk-3-dev
* libglib2.0-dev
* libcairo2-dev
* github.com/mattn/go-gtk/glib
* github.com/mattn/go-gtk/gtk
* GTK+ 3.20


## Compilacion

```bash
$ go build
```

## Levantar el servidor

```bash
$ go run Servidor.go <ip> <puerto>
```

## Conectar un cliente

```bash
$ go run Cliente.go <ip> <puerto>
