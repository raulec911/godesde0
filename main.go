package main

import (
	"fmt"

	"github.com/raulec911/godesde0/ejercicios"
)

func main() {
	/*estado, texto := variables.ConviertoaTexto(100)
	fmt.Println(estado)
	fmt.Println(texto)
	os := runtime.GOOS

	if os == "linux" || os == "OS X." {
		fmt.Println("Esto no es Windows, es ", os)
	} else {
		fmt.Println("Esto es Windows")
	}

	switch os {
	case "linux":
		fmt.Println("Esto es linux")
	case "darwin":
		fmt.Println("Esto es darwin")
	default:
		fmt.Printf("%s", os)
	}*/

	numero, texto := ejercicios.ConvNumerico("500A")
	fmt.Println(numero)
	fmt.Println(texto)
}
