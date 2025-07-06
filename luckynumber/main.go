package main

import "fmt"

// Given an array of integers arr, a lucky integer is an integer that has a frequency in the array equal to its value.

// Return the largest lucky integer in the array. If there is no lucky integer return -1.
//
func main() {

	testArray := []int{1, 1, 1, 2, 2, 33, 3, 4, 4, 4, 5, 5, 6, 6, 6, 6, 6, 6}

	luckyNumber := lucky(testArray)

	fmt.Println(luckyNumber)

}

func lucky(arr []int) int {
	if len(arr) == 0 {
		return -1
	}

	miMapa := make(map[int]int, len(arr))

	for _, value := range arr {
		mapValue, _ := miMapa[value]
		miMapa[value] = mapValue + 1
	}

	biggestNumber := -1
	for key, _ := range miMapa {
		if key == miMapa[key] {
			if key > biggestNumber {
				biggestNumber = key
			}
		}
	}

	return biggestNumber

}
