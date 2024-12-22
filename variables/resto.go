package variables

import (
	"fmt"
	"strconv"
	"time"
)

// var nombre string // variable unica del archivo
var Nombre string // variable publica, para todo el paquete o donde se importe
var Estado bool
var Sueldo float32
var Fecha time.Time

func RestoVariables() {
	Nombre = "pedro"
	Estado = true
	Sueldo = 1451.44
	Fecha = time.Now()

	fmt.Println(Nombre)
	fmt.Println(Estado)
	fmt.Println(Sueldo)
	fmt.Println(Fecha)
}

func ConviertoATexto(numero int) (bool, string) {
	texto := strconv.Itoa(numero)
	return true, texto
}
