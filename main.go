package main

import (
	"fmt"

	"github.com/DaniJSH/godesde0/variables"
)

func main() {
	estado, texto := variables.ConviertoaTexto(1427)
	fmt.Println(estado)
	fmt.Println(texto)
}

