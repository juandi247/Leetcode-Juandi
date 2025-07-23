package main

import "fmt"

type Node struct {
	value     int
	childrenA *Node
	childrenB *Node
}

func main() {

	// fmt.Println("resultaod", sumNumbers(5))
	// fmt.Println("resultaod", sumNumbers(5))
	fmt.Println("resultaod", reverseString("abcdefg"))

}

func sumNumbers(upperbonud int) int {
	if upperbonud == 0 {
		return 0
	}
	fmt.Println("current number", upperbonud)
	return upperbonud + sumNumbers(upperbonud-1)

}



func factorial(number int )int {

	if number ==1 {
		return 1
	}
fmt.Println("number", number)
	return number * factorial(number -1 )
}





func reverseString(myString string) string {

if myString==""{
	return ""
}
fmt.Println("vamos en ", myString)
return reverseString(myString[1:]+ myString[0:1])

}

/*
Suma de 1 a N

Escribe una función recursiva que sume los números del 1 hasta n.

Factorial

Implementa una función que calcule el factorial de un número n.

Fibonacci

Implementa la serie de Fibonacci usando recursión simple.

Reversa de string

Haz una función que reciba un string y lo devuelva al revés usando recursión.

Conteo de dígitos

Cuenta cuántos dígitos tiene un número usando recursión.
*/
