package arraysSlices

import "fmt"

// define an array (its not initialized yet)

func ModifyArray(arr *[]int) int {
	sum := 0
	for _, value := range *arr {
		sum = value + sum
	}
	return sum
}

func InvertArray(arr []int) {
	start := 0
	end := len(arr) - 1

	for start != end {
		valorFinal := arr[end]

		arr[end] = arr[start]
		arr[start] = valorFinal

		start++
		end--
	}
	fmt.Println("array invertido: ", arr)
}

func SearchNumberNotOrderedList(arr []int, target int) bool {
	for _, value := range arr {
		if value == target {
			return true
		}
	}
	return false
}

// binary search implmementation
func SearchNumberOrderedList(arr []int, target int) bool {

	startPointer := 0
	endPointer := len(arr) - 1

	for startPointer != endPointer {
		rango := (startPointer + endPointer) / 2

		if arr[rango] == target {
			return true
		}

		if target > arr[rango] {
			startPointer = rango + 1
		}

		if target < arr[rango] {
			endPointer = rango - 1
		}
	}
	return false

}

func Hey() {
	// !this wont work, we intialize an array with CERO length
	//! in order to moidify it or appneed more things, we need to crea te SLICE
	array := []int{}
	array[0] = 2
	fmt.Println(array)
}

func HeyQithSlice() {
	// !this wont work, we intialize an array with CERO length
	//! in order to moidify it or appneed more things, we need to crea te SLICE
	array := make([]int, 3, 10)
	array[5] = 2
	fmt.Println(array)
}
