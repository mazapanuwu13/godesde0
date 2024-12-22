package main

import (
	"fmt"

	"github.com/mazapanuwu13/godesde0/variables"
)

func main() {
	fmt.Println("hola desde main")
	estado, texto := variables.ConviertoATexto(1588)

	fmt.Print("estado: ", estado, "texto: ", texto)
}
