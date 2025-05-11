package main

import (
	// "LeetcodeJuandi/PointersGolang"
	"LeetcodeJuandi/excercises/arraysSlices"
	// "LeetcodeJuandi/excercises/factorialpackage"
	"fmt"
)


var mySlice []int  //here we define a slice, its NOT an array

func main() {
// append things to a slice
	mySlice = append(mySlice, 3,3,43,5,45,46,4)
	fmt.Println(mySlice)


	
	nuevoArray := []int{1,2,3,4,5}
	fmt.Println(arraysSlices.ModifyArray(&nuevoArray))

	arraysSlices.InvertArray(nuevoArray)


	OrganizedArray := []int{1,2,3,4,5,7,8,9,10,34,55,77,90,100}
	DisorganizedArray := [10]int{2,6,8,2,867,23,6575,12,23,34}


	fmt.Println("Nos dio que el numero: ", arraysSlices.SearchNumberNotOrderedList(DisorganizedArray[:],34))

	fmt.Println("Nos dio que el numero: ", arraysSlices.SearchNumberNotOrderedList(DisorganizedArray[:],90))

	fmt.Println("binary search fue: ",arraysSlices.SearchNumberOrderedList(OrganizedArray,3))


	// arraysSlices.HeyQithSlice()

	// PointersGolang.FirstExcersice()

	// a:=30
	// b:=54232
	// PointersGolang.Swap(&a,&b)

	// p := PointersGolang.RetornarPunteroStruct("mercedses", 2000)
	// fmt.Println(*p)

	// fmt.Printf("Factorial: %d", factorialpackage.Factorial(4))

	// MazeRecursion.Test()

}
