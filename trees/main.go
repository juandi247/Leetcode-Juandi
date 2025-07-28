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
	// fmt.Println("resultaod", reverseString("abcdefg"))
	coso:= calcRange(10)
fmt.Println("rta de rango de 1, ", coso[:])
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



func reverseStrings(coso string) string {
	if coso==""{
		return ""
	}
return reverseStrings(coso[0:1]) + coso[1:]

}







func reverseString(myString string) string {

if myString==""{
	return ""
}
fmt.Println("vamos en ", myString)
return reverseString(myString[1:]+ myString[0:1])

}


// rangoHasta(n) -> Lista de números: dado un número "n", retorna la lista de números desde el 0 hasta el N incluído. Por ejemplo: rangoHasta(5) -> [0,1,2,3,4,5].

func calcRange (target int) []int{

	if target==0{
		return []int{10}
	}


a:= calcRange(target-(target+1))
return append(a, target)
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
