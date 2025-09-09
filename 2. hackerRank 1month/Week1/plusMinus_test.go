package main

import (
	"fmt"
	"strconv"
)

// Given an array of integers, calculate the ratios of its elements that are positive, negative, and zero. Print the decimal value of each fraction on a new line with  places after the decimal.

// Note: This challenge introduces precision problems. The test cases are scaled to six decimal places, though answers with absolute error of up to  are acceptable.

// Example

// There are  elements, two positive, two negative and one zero. Their ratios are ,  and . Results are printed as:

// 0.400000
// 0.400000
// 0.200000

func plusMinus(arr []int32){

miMap:= map[int8]float32{
	1:0,
	0:0,
	-1:0,
}
strconv

for _,v:=range arr{

	if v>0{
		miMap[1]++
	}else if v<0{
		miMap[-1]++
	}else{
		miMap[0]++
	}

}

length:= float32(len(arr))

// Print the decimal values (we need to make a FLOAT DIVISOIN, if not it wont work and with 6 decimals)

fmt.Printf("%.6f \n", miMap[1]/length)
fmt.Printf("%.6f \n", miMap[-1]/length)
fmt.Printf("%.6f \n", miMap[0]/length)

}
