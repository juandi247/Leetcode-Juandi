package main

import (
	"LeetcodeJuandi/PointersGolang"
	"LeetcodeJuandi/excercises/factorialpackage"
	"fmt"
)

func main() {

	PointersGolang.FirstExcersice()

	// a:=30
	// b:=54232
	// PointersGolang.Swap(&a,&b)

	p := PointersGolang.RetornarPunteroStruct("mercedses", 2000)
	fmt.Println(*p)

	fmt.Printf("Factorial: %d", factorialpackage.Factorial(4))

	// MazeRecursion.Test()

}
