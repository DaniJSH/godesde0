package variables

//función que muestra por consola tres variables enteras

import "fmt"

func MuestraEnteros() {
	
	//declaración
	var intcomun int

	//asignación
	intde32 := int32(10)
	intde64 := int64(100)

	fmt.Println("intcomun = ", intcomun)
	fmt.Println("intde32 = ", intde32)
	fmt.Println("intde64 = ", intde64)
}