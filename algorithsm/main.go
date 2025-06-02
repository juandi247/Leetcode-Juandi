package main

import "fmt"

// ! this was created to write the bubble sort and the quick sort algotihsm, ewith somw excercises inside
func main() {

	testSlice := []int{2, 3, 4, 7, 1, 2, 8, 1, 9, 5, 5, 2, 1, 1, 4, 56, 1}

	QuickSort(testSlice, 0, len(testSlice)-1)


	
	// this final implementation (can be improved by checking if the array has no swaps, its already sorted, but it has the logic)
	// BubbleSort(testSlice)
	fmt.Println(testSlice)




}


func partition( miarray[]int, low, high int) int {//return the pivot index

	pivot:= miarray[high]
	startingIndex:= low -1  //because we need to add to the index, we need to start it in 0 -1, so on the funciotn we chck it and +1


	for i:=low; i< high; i++{
		// if our current value is less than the pivot, we need to swap it to the left
		if(miarray[i]<= pivot){
			startingIndex++
			tmp:= miarray[i]
			miarray[i]= miarray[startingIndex]
			miarray[startingIndex]= tmp
		}
	}

	startingIndex++
	miarray[high]= miarray[startingIndex]
	miarray[startingIndex]= pivot

return startingIndex
}
func QuickSort(miArray []int, low, high int){
	if low >= high{
		return
	}

	pivotIndex:= partition(miArray, low, high)
	QuickSort(miArray, low, pivotIndex-1)
	QuickSort(miArray, pivotIndex+1, high)

// quick sort works as divide and cqoncuer?? conquer? idk
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
