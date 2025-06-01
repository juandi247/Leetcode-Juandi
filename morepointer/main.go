package main

import "fmt"

func main() {

	num := 1

	func(number *int) {
		*number += 1
	}(&num)

	fmt.Printf("the funciton changed the number from 1, to %d \n", num)


	// declare a pointer to num

	pNum:= &num
	fmt.Printf("this is the addres memory (or a pointer to the num) %d \n" , pNum)
	fmt.Printf("this is the value of the addres where the pointer is to %v its a derreferenece", *pNum)

	num2:= 43

	// pNum2:= &num2


fmt.Println("\nclaores NO CAMIADOS", num, num2)

Swap(&num, &num2)

fmt.Println("\nclaores camibados", num, num2)



// miSlice:= []int{2,2,34,3}
}

func Swap(num1, num2 *int) {
	temp:= *num2
	*num2= *num1
	*num1=temp
}


func mifuncion(miSlice []int) int{
	var promedio int
	for i := range(miSlice){
		promedio+= miSlice[i]
	}
	return promedio
}



// Ejercicios:
// Crea una función que reciba un slice de enteros y devuelva el promedio.

// Implementa una función que elimine un elemento por índice de un slice.

// Crea una función que combine dos slices de strings eliminando duplicados.

// Implementa una función que recorte (truncate) un slice a un tamaño máximo dado.


