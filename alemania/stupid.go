package main

import "fmt"

// Descripción:
// Imprime números del 1 al n, pero:

// Si es divisible por 3, imprime "Fizz".

// Si es divisible por 5, imprime "Buzz".

// Si es divisible por ambos, imprime "FizzBuzz".

func FBcoso(target int) {
	for i:=1; i<=target; i++{
		if i%3==0{
			fmt.Printf("number %v Fizz \n", i)
		}

		if i%5==0{
			fmt.Printf("number %v Buzz \n", i)

		}
		if i%3==0 && i%5==0{
			fmt.Printf("number %v Fizzbuzz???? \n", i)
		}


	}
}