package main

import "fmt"

// ! this was created to write the bubble sort and the quick sort algotihsm, ewith somw excercises inside
func main() {

	testSlice := []int{2, 3, 4, 7, 1, 2, 8, 1, 9, 5, 5, 2, 1, 1, 4, 56, 1}
	// this final implementation (can be improved by checking if the array has no swaps, its already sorted, but it has the logic)
	BubbleSort(testSlice)

	fmt.Println(testSlice)

}

func BubbleSort(miArray []int) {

	/* orderedSlice:= make([]int, len(miArray))
	how does bubble sort work? its defined by poping the biggest element of the array to the last element.
	So there would be a ouble iteration kinda like a nested for loop, but the twist is that we dont need to
	check till the last element (because we know that already is the biggest, so we iterate till Len(slice)- 1(-i))
	*/
	// We compare the actual number to the next one, if the actual is bigger we SWAP, if not we conitnue normally wihtout swapping
	initialLength := len(miArray)
	for i := 0; i < len(miArray); i++ {
		fmt.Println("ejecutando algo otra vez")
		// on the first excecution, we dont need to get until the end, why? because the biggest nuumber will already be there by the last swap
		// we sustract the first i, because we need progresively avoid using the last number because its already th ebiggest checked
		for j := 0; j < initialLength-i-1; j++ {
			if miArray[j] > miArray[j+1] {
				// swap
				temp := miArray[j]
				miArray[j] = miArray[j+1]
				miArray[j+1] = temp
			}
		}

	}
}
