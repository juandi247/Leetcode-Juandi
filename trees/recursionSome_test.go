package main

import (
	"fmt"
	"testing"
)

func Factorial(num int) int {
	if num == 1 {
		return 1
	}
	return num * Factorial(num-1)
}

func Test_Factorial(t *testing.T) {

	miStruct := []struct {
		number int
		result int
	}{
		{5, 5 * 4 * 3 * 2 * 1},
		{10, 10 * 9 * 8 * 7 * 6 * 5 * 4 * 3 * 2 * 1},
		{2, 2},
		{1, 1},
		{3, 3 * 2 * 1},
	}

	for i, v := range miStruct {

		t.Run("test", func(t *testing.T) {

			resultado := Factorial(v.number)

			if resultado != v.result {
				t.Error("AAA ERROR en ", i)
			}
		})
	}
}




// 2. Suma de los primeros N números
// Suma todos los números del 1 hasta n.

func Suma(number int) int{
if number ==0{
	return 0
}
return number + Suma(number-1)
}

func Test_Suma(t *testing.T) {

	miStruct := []struct {
		number int
		result int
	}{
		{5, 5 + 4 + 3 + 2 + 1},
		{10, 10 + 9 + 8 + 7 + 6 + 5 + 4 + 3 + 2 + 1},
		{2, 2+1},
		{1, 1},
		{3, 3 + 2 + 1},
	}

	for i, v := range miStruct {

		t.Run("test", func(t *testing.T) {

			resultado := Suma(v.number)

			if resultado != v.result {
				t.Error("AAA ERROR en ", i)
			}
		})
	}
}


func sumarArreglo(array []int) int{

	if len(array)== 0{
		return 0
	}
	return array[0] + sumarArreglo(array[1:])
}



func Test_SumarArreglo(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{input: []int{}, expected: 0},
		{input: []int{5}, expected: 5},
		{input: []int{1, 2, 3}, expected: 6},
		{input: []int{-1, -2, -3}, expected: -6},
		{input: []int{0, 0, 0}, expected: 0},
		{input: []int{10, -5, 15}, expected: 20},
	}

	for _, tt := range tests {
		result := sumarArreglo(tt.input)
	
		if result != tt.expected {
			t.Errorf("sumarArreglo(%v) = %d; expected %d", tt.input, result, tt.expected)
		}
		fmt.Println("EXITO ")
	}
}





func buscarEnArregloNoOrdenado(array []int, target int) bool{

if len(array)== 0{
	return false
}

if target != array[0]{
	return buscarEnArregloNoOrdenado(array[1:], target)
}

return true

}


func TestContains(t *testing.T) {
	tests := []struct {
		arr      []int
		target   int
		expected bool
	}{
		{arr: []int{}, target: 5, expected: false},
		{arr: []int{1, 2, 3, 4}, target: 3, expected: true},
		{arr: []int{1, 2, 3, 4}, target: 5, expected: false},
		{arr: []int{5}, target: 5, expected: true},
		{arr: []int{10, 20, 30, 40}, target: 10, expected: true},
		{arr: []int{10, 20, 30, 40}, target: 40, expected: true},
		{arr: []int{10, 20, 30, 40}, target: 25, expected: false},
	}

	for _, tt := range tests {
		result := buscarEnArregloNoOrdenado(tt.arr, tt.target)
		if result != tt.expected {
			t.Errorf("contains(%v, %d) = %v; expected %v", tt.arr, tt.target, result, tt.expected)
		}
		fmt.Println("EXITOooo")
	}
}



func INVERSION(cadena string) string{
//base case
//vamos 


return INVERSION(cadena[:1])

}





func invertirCadena(texto string)string{

if texto=="" || len(texto)==0{
	return ""
}
length:= len(texto)

return texto[length-1:] + invertirCadena(texto[:length-1])
}

func TestReverseString(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "", expected: ""},
		{input: "a", expected: "a"},
		{input: "ab", expected: "ba"},
		{input: "hello", expected: "olleh"},
		{input: "GoLang", expected: "gnaLoG"},
		{input: "12345", expected: "54321"},
		{input: "anita lava la tina", expected: "anit al aval atina"},
	}

	for _, tt := range tests {
		result := invertirCadena(tt.input)
		if result != tt.expected {
			t.Errorf("reverseString(%q) = %q; expected %q", tt.input, result, tt.expected)
		}
	}
}