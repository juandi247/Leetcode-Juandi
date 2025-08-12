package main

import (
	"fmt"
	"strings"
	"testing"
	"unicode"
)

// 1️⃣ Valid Palindrome
// Descripción:
// Dada una cadena de caracteres s, devuelve true si es un palíndromo, o false en caso contrario.
// Debes ignorar mayúsculas/minúsculas y todos los caracteres que no sean alfanuméricos.

func ValidPal(texto string) bool {
	cleanString := strings.TrimSpace(texto)

	fmt.Println("timmed", cleanString)
	cleanMayus := strings.ToLower(cleanString)
	fmt.Println("clean", cleanMayus)

	runeSlice := []rune{}
	for _, letter := range cleanMayus {
		if unicode.IsLetter(letter) {
			runeSlice = append(runeSlice, letter)
		}
	}

	// two pointer aproach
	if len(runeSlice) < 1 {
		return true
	}

	startPointer := 0
	endPointer := len(runeSlice) - 1

	fmt.Println("COso es: ", string(runeSlice))

	for startPointer < endPointer {
		if runeSlice[startPointer] != runeSlice[endPointer] {
			return false
		}

		if runeSlice[startPointer] == runeSlice[endPointer] {
			fmt.Println("COICNIDE UNA LETRA")
		}
		fmt.Printf("vamos a comprar el %v coin el %v", startPointer, endPointer)
		startPointer++
		endPointer--
	}

	return true
}

func TestPalindrome(t *testing.T) {

	testStruct := []struct {
		Text     string
		Expected bool
	}{
		{"A man, a plan, a canal: Panama", true},
		{"No lemon, no melon", true},
		{" ", true},
		{"0P0", true},
		{"Madam", true},
		{"Was it a car or a cat I saw?", true},

		{"teamo miamor", false},
		{"race a car", false},
		{"hello", false},
		{"palindrome", false},
		{"abcba!",true},
		{"A Toyota", true},
	}

	for _, v := range testStruct {
		result := ValidPal(v.Text)
		if result != v.Expected {
			t.Errorf("No coincidio en el test  texto: %v  exp: %v res: %v \n \n", v.Text, v.Expected, result)
		} else {
			fmt.Println("SUCCES")
		}

	}

}
