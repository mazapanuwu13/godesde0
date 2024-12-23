package main

import (
	"fmt"

	"github.com/mazapanuwu13/godesde0/ejercicios"
	// "runtime"
	// "github.com/mazapanuwu13/godesde0/variables"
)

func main() {
	fmt.Println("hola desde main")
	// estado, texto := variables.ConviertoATexto(1588)
	// fmt.Print("estado: ", estado, "texto: ", texto)

	// if os := runtime.GOOS; os == "linux" {
	// 	fmt.Println("esto es linux")
	// } else {
	// 	fmt.Println("esto es window")
	// }

	// switch os := runtime.GOOS; os {
	// case "linux":
	// 	fmt.Println("esto es linux")
	// case "drawin":
	// 	fmt.Println("esto es drawin")
	// default:
	// 	fmt.Printf("%s \n", os)
	// }

	num, txt := ejercicios.Ejercicio01("600")

	fmt.Println(num)
	fmt.Println(txt)
}
