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
* libgtk+2.0
* libgtksourceview2.0-dev
* GTK+ 3.20

Esta biblioteca va en el directorio $GOPATH

* github.com/mattn/go-gtk/gtk

Instalacion :

```bash
$ go get github.com/mattn/go-gtk/gtk
```
```bash
$ sudo apt-get install libgtk-3-dev 
```
```bash 
$ sudo apt-get install libcairo2-dev 
```
```bash 
$ sudo apt-get install libglib2.0-dev
```
```bash
$ sudo apt-get install libgtk+2.0
```
```bash
$ sudo apt-get install libgtksourceview2.0-dev
```

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
