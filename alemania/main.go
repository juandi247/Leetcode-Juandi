package main

import (
	"fmt"
	// "warmup/trees"
)

func main() {

// tesT:=[]int{0,0,1,1,1,2,2,3,3,4}

// rta:= RemoveDuplicates(tesT)
// fmt.Println("RTA ES ", rta)

 L := &Node{Value: 8}
    M := &Node{Value: 9}
    N := &Node{Value: 10}
    O := &Node{Value: 11}
    P := &Node{Value: 12}
    Q := &Node{Value: 13}
    R := &Node{Value: 14}
    S := &Node{Value: 15}

    // Nivel 2
    E := &Node{Value: 4, Children: []*Node{L, M}}
    F := &Node{Value: 5, Children: []*Node{N, O}}
    G := &Node{Value: 6, Children: []*Node{P, Q}}
    H := &Node{Value: 7, Children: []*Node{R, S}}

    // Nivel 1
    B := &Node{Value: 2, Children: []*Node{E, F}}
    C := &Node{Value: 3, Children: []*Node{G, H}}

    // Root
    A := &Node{Value: 1, Children: []*Node{B, C}}

// 	        	      A(1)
//           /                 \
//        B(2)                           C(3)
//       /    \                   /               \
//    E(4)   F(5)                G(6)             H(7)
//    / \    /      \           /  \              /  \
// L(8) M(9) N(10) O(11)      P(12) Q(13)    R(14) S(15)




Preorder(A)
// se deberia ver como -> 1,2,4,8,9,5,10,11,3,6,12,13,7,14,15
fmt.Print("\n")

PreorderInversed(A)
// se deberia ver como -> 1, 3,7,15,14,6,13,12,etc,etc


fmt.Print("\n ---- \n")


AA := &Node{Value: 8}
BB := &Node{Value: 9}
CC := &Node{Value: 10}
DD := &Node{Value: 11}
EE := &Node{Value: 12}
FF := &Node{Value: 13}
GG := &Node{Value: 14}
HH := &Node{Value: 15}

// Nivel 2
II := &Node{Value: 4, Left: AA, Right: BB}
JJ := &Node{Value: 5, Left: CC, Right: DD}
KK := &Node{Value: 6, Left: EE, Right: FF}
LL := &Node{Value: 7, Left: GG, Right: HH}

// Nivel 1
MM := &Node{Value: 2, Left: II, Right: JJ}
NN := &Node{Value: 3, Left: KK, Right: LL}

// Root
OO := &Node{Value: 1, Left: MM, Right: NN}


fmt.Println("AAA BINARY TREE IN ORDER")
InOrder(OO)
// listaVacia:= []int{}


// result:= mitrees.Pre_order_search(A, listaVacia)
fmt.Println("RESULTADOOOO: deberia ser el mismo del priemro", )
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

