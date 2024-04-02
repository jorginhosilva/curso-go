package main

import (
	"fmt"
)

func main() {
	
	var igualdade bool

	val1 := 10
	val2 := 100

	igualdade = val1 == val2

	fmt.Println("Operador de Igualdade:", igualdade)

	var diferente bool

	diferente = val1 != val2

	fmt.Println("Operador Diferente:", diferente)

	var menorQue bool

	menorQue = val1 <= val2

	fmt.Println("Operador Menor Que: ", menorQue)

	var menor bool

	menor = val1 < val2

	fmt.Println("Operador Menor:", menor)

	var maiorQue bool

	maiorQue = val1 >= val2

	fmt.Println("Operador Maior Que:", maiorQue)

	var maior bool

	maior = val1 > val2

	fmt.Println("Operador Maior:", maior)

}