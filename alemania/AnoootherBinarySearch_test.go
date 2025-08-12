package main

import (
	"fmt"
	"testing"
)

func BinarySearch(miArray []int, target int) int {

	start := 0
	end := len(miArray) - 1

	for start <= end {
		midPoint := (start + end) / 2
		if miArray[midPoint] == target {
			return midPoint
		}
		if miArray[midPoint] > target {
			end = midPoint - 1
		}
		if miArray[midPoint] < target {
			start = midPoint + 1
		}
	}

	return -1
}





func TestBinarySearch(t *testing.T) {
    testingStruct := []struct {
        array  []int
        target int
        want   int
    }{
        {[]int{3, 3, 4}, 2, -1},
        {[]int{1, 2, 3, 4, 5}, 3, 2},
        {[]int{1, 2, 3, 4, 5}, 5, 4},
    }

    for _, tt := range testingStruct {
        got := BinarySearch(tt.array, tt.target)
        if got != tt.want {
            t.Errorf("BinarySearch(%v, %d) = %d; want %d",
                tt.array, tt.target, got, tt.want)
        }else{
			fmt.Println("SUCESS")
		}
    }
}