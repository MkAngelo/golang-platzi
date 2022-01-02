# GO/GOLANG

Golang es un lenguaje de programación creado con la eficiencia de **C** y la syntaxis de **Python**

Es un lenguaje estrictamente tipado creado en Google por Robert, Rob Pike y Ken Thompson. Anunciado en Noviembre de 2009 y su primera version fue en 2012. Gophers

La forma en que se organiza un proyecto en **GO** es la siguiente:
 * bin: guardar los ejecutables
 * pkg: guarda ciertos codigos de go modules
 * src: Aquí se escribe el codigo/ Se guarda el conjunto de codigo de librerias de terceros.

> Que es package?
 
    Un package es el nombre de la corpeta donde esta guardado

 La forma de ejecutar un archivo .go es con *build* de esta forma:
 * go build archivo.go

 Ó de la siguente forma (más rapida), usando *run*:
 * go run archivo.go

En Go se usan los slices (mutables), arrays (no mutables), los maps(diccionarios).

### TIPOS DE DATOS

> Numeros enteros
  
* int = Depende del OS (32 o 64 bits)
* int8 = 8bits = -128 a 127
* int16 = 16bits = -2^15 a 2^15-1
* int32 = 32bits = -2^31 a 2^31-1
* int64 = 64bits = -2^63 a 2^63-1

>Optimizar memoria cuando sabemos que el dato simpre va ser positivo

* uint = Depende del OS (32 o 64 bits)
* uint8 = 8bits = 0 a 127
* uint16 = 16bits = 0 a 2^15-1
* uint32 = 32bits = 0 a 2^31-1
* uint64 = 64bits = 0 a 2^63-1

>Números decimales
* float32 = 32 bits = +/- 1.18e^-38 +/- -3.4e^38
* float64 = 64 bits = +/- 2.23e^-308 +/- -1.8e^308

>Textos y booleanos
* string = ""
* bool = true or false

>Números complejos
* Complex64 = Real e Imaginario float32
* Complex128 = Real e Imaginario float64
* Ejemplo : c:=10 + 8i

### STRUCTS
Es la forma de hacer clases en Go.

### CREAR PAQUETES EN GOLANG
Al momento de crear paquetes el GO, se presentan algunos errores con respecto a la creacion de estos en carpetas separadas, es por ello que la solucion es usar lo siguiente:

* Estar en la carpeta raiz del proyecto y ejecutar
    
    `go mod init <nombreCarpetaRaiz>`

* Obtenemos un archivo con el nombre go.mod
* Posteriormente en nuestro `main.go` importamos nuestro modulo desde la carpeta raiz

    `import "platzi/src/modulo"`