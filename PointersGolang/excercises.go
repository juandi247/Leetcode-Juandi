package PointersGolang

import "fmt"



// ? EASY LEVEL

//! 1.
// Crea una variable x := 10.
// Crea un puntero p que apunte a x.
// Usa el puntero para cambiar el valor de x a 20.
// Imprime x y verifica que realmente cambió.

func FirstExcersice() {

	// asign value
	x := 10
	// create pointer to X
	p := &x
	// Modify the value of X, using the pointer (and the *p to look for the value, if we use p that is the memory addres
	*p = 20
	// print that changed
	fmt.Println(x)

}

//! 2.
// Intercambiar valores con punteros
// Crea una función swap(a, b *int) que intercambie los valores de a y b usando punteros.
// Prueba la función en el main().

func Swap(a, b *int) {

	valorA := *a
	*a = *b
	*b = valorA

	fmt.Printf("nuevo valor de a: %d, nuevo valor de b: %d", *a, *b)
}






//  ? MEDIUM LEVEL



// Nivel 2: Medio
// Sumar elementos de un array usando punteros
// Define un array [5]int{1, 2, 3, 4, 5}.
// Crea una función que reciba un puntero al primer elemento del array y sume todos los elementos.

func SumArray(arr *[5]int) int {
   // aquí va tu código
	sum:=0
   for i:= range arr{
	sum=sum+(*arr)[i]
   }
fmt.Println(sum)
   return sum
}




// Crear un nuevo struct usando punteros
// Define un struct llamado Car con campos brand y year.
// Crea una función que reciba los valores y retorne un puntero a un Car nuevo.

type Car struct{
	brand string
	year int
}

func RetornarPunteroStruct(brand string, year int) *Car{

	carrito:= Car{brand: brand, year: year}

	return &carrito
}