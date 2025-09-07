package main

import "testing"

/*
Exercise 2 – Custom-Sorted Array
Problem Statement

You are given an array of positive integers.

In a single operation called a move, you may swap the elements at any two indices.

Your task is to determine the minimum number of moves required to sort the array so that:

All even elements appear at the beginning of the array.

All odd elements appear at the end of the array.

The order of the elements within the even group or within the odd group does not matter.

Example
arr = [16, 3, 4, 51]


Valid custom-sorted arrays include:

[16, 4, 3, 51]

[4, 16, 3, 51]

[16, 4, 51, 3]

[4, 16, 51, 3]

The most efficient solution requires 1 move (swap 3 and 4).

Output:

1

Function Signature
int moves(int[] arr)
*/


func minNumberSwaps(array []int) int{
	Start:=0
	End:=1

	numberOfSwaps:=0


	for End< len(array){
		// is ODD, IMPAR
		if array[Start]%2!=0{
			// verify if the END is even to swap
			if array[End]%2==0{
				// make a swap
				array[End], array[Start]= array[Start], array[End]
				numberOfSwaps++
			
			}
		}
		Start++
		End++
	}


return numberOfSwaps



}


/*
Aprach:

Since we shoudl have the even (pares) at the begginign, and at the end WITHOUT mattering the order the ODDS

My aproach shuold be a two pointer aproach. 

1. Have the first pointer to the cero value and the other to the first one. 


	Is the first pointer odd?? 
	If yes we need to check the other pointer, is the other pointer an Even?

*/

func Test_CustomSortedArray_Test(t *testing.T) {
    testCases := []struct {
        array    []int
        expected int
    }{
        {
            array:    []int{16, 3, 4, 51},
            expected: 1, // swap 3 and 4
        },
        {
            array:    []int{2, 4, 6, 1, 3, 5},
            expected: 0, // already separated
        },
        {
            array:    []int{1, 3, 5, 7, 2, 4, 6},
            expected: 3, // swap each odd with an even
        },
        {
            array:    []int{1, 2, 3, 4, 5, 6},
            expected: 3,
        },
        {
            array:    []int{2, 1, 4, 3, 6, 5},
            expected: 3,
        },
        {
            array:    []int{1, 3, 5, 7},
            expected: 0, // no evens, no swaps needed
        },
        {
            array:    []int{2, 4, 6, 8},
            expected: 0, // only evens, no swaps needed
        },
        {
            array:    []int{}, // empty array
            expected: 0,
        },
    }

    for i, tc := range testCases {
        result := minNumberSwaps(tc.array)
        if result != tc.expected {
            t.Errorf("Test case %d failed: got %d, expected %d", i+1, result, tc.expected)
        }
    }
}
