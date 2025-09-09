package main

import "fmt"

/*
Given five positive integers, find the minimum and maximum values that can be calculated by summing exactly four of the five
integers. Then print the respective minimum and maximum values as a single line of two space-separated long integers.

Example
arr [1,3,5,7,9]

The minimum sum is  and the maximum sum is 16 and max is 24
*/

func minMaxSum(arr []int) {

	/* we are given 5 numbers, and need to select four
	    to obtain the biggest sum and the smallest sum posible,
	   we can find the biggest, and the smallest element in the array.

	   We can do this in a O(N) and at the same time we are going to add  up all the numbers
	   So we obtain the sum of the 5 numbers, now we print the minSum that should be
	   minSum:= sumaCompleta - biggestNumber
	   maxSum:= sumaCompleta - smallestNumber
	*/
	maxNumber := arr[0]
	minNumber := arr[0]
	Total := 0
	for _, v := range arr {
		Total += v

		if v > maxNumber {
			maxNumber = v
		}

		if v < minNumber {
			minNumber = v
		}

	}

	fmt.Printf("%v %v", (Total - maxNumber), (Total - minNumber))

}
