package main

import (
	"fmt"
)

func main() {
rta:=reverseSame("Holamiamor")

rtaa:=easyAproach("pescaditopipelonciooooo")
fmt.Println("REPUESTA",rta)
fmt.Println("REPUESTA DOS",rtaa)

}

// reverse String

/* I need to reverse a string. "ABC" -> "CBA"
Questions:
	Are they all normal english Characters? this should not affect the result but still its important
	Constrains on the limit? 1 to some number

Aproach:
	There are some aproaches, in this case i would have two
	Creating a copy of the string
	Using the same string
*/

func reverseSame(text string) string {
	// we are going to have TWO pointers aproach
	start := 0
	end:= len(text)-1

	sliceText:= []rune(text)	

	for start< end{
		sliceText[start]= rune(text[end])
		sliceText[end]= rune(text[start])
		start++
		end--
	}
	return string(sliceText)
}


func easyAproach(texto string) string{
var mitexto string
for i:=len(texto)-1; i>0; i--{
mitexto= mitexto+string(texto[i])
}
return mitexto
}

