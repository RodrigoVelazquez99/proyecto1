# proyecto1

Chat programado en Go
Modelado y Programacion
Go version 1.10.1

#Autor

Velázquez Cruz Rodrigo Fernando
Numero de cuenta UNAM: 315254565

## Dependencias

* libgtk-3-dev
* libglib2.0-dev
* libcairo2-dev

Estas bibliotecas van en el directorio $GOPATH

* github.com/mattn/go-gtk/glib
* github.com/mattn/go-gtk/gtk


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

Error con la interfaz grafica, se ingresa el texto por la interfaz, pero el texto recibido del servidor se imprime unicamente en terminal.

Faltan pruebas unitarias.
