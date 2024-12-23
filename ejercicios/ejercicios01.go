package ejercicios

import (
	"strconv"
)

var Number int

func Ejercicio01(par1 string) (int, string) {

	num, err := strconv.Atoi(par1)

	if err != nil {
		return 0, "hubo un error " + err.Error()
	}

	if num > 100 {

		return num, "Es mayor a 100"
	} else {
		return num, "Es menor a 100"
	}
}
