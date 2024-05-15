package variables

import (
	"fmt"
	"time"
	"strconv"
)

var Nombre string		// <- ""  tipo string
var Estado bool			// tipo boolean
var Sueldo float32		// tipo flotante 32
var Fecha time.Time


func RestoVariables() {
	Nombre = "Dani"
	Estado = true
	Sueldo = 1345.45
	Fecha = time.Now()
	
	fmt.Println(Nombre)
	fmt.Println(Estado)
	fmt.Println(Sueldo)
	fmt.Println(Fecha)

}

func ConviertoaTexto(numero int) (bool, string) {
	texto := strconv.Itoa(numero)
	return true, texto
}