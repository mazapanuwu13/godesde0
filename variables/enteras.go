package variables

import (
	"fmt"
)

func MuestroEnteros() { // al tener una letra en mayuscula estamos declarando que la funcion se puede exportar
	var intComun int //declaracion comun, no se define como nula, si no con el valor minimo de ese tipo de dato

	intde32 := int32(10) //crear variable por asignacion, toma el tipo de dato
	intde64 := int64(100)

	fmt.Println("int comun = ", intComun)
	fmt.Println("int de 32 = ", intde32)
	fmt.Println("int de 64 = ", intde64)
}
