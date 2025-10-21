package main

type minHeap struct {
	size int
	arr  []int
}

func minHeapConstructor() *minHeap {
	return &minHeap{
		size: 0, arr: []int{},
	}
}

// to add to a minheap. we append to the latests
func (h *minHeap) add(number int) {

	h.arr = append(h.arr, number)
	h.size++

	//heapify up
	lastPosition := h.size - 1

	for lastPosition > 0 {

		parent := (lastPosition - 1) / 2
		//si el parent es mayor, significa que debemos swap, porque el mas chiquito debe ser el papá
		if h.arr[parent] < h.arr[lastPosition] {
			break
		}

		h.arr[parent], h.arr[lastPosition] = h.arr[lastPosition], h.arr[parent]
		lastPosition = parent
	}

}

func (h *minHeap) heapifyDown(number int) {

	h.arr[0] = number

	currIndex := 0

	for currIndex < len(h.arr) {

		left := currIndex*2 + 1
		right := currIndex*2 + 2

		if left >= len(h.arr) {
			break
		}

		if right >= len(h.arr) {
			if h.arr[index] < h.arr[left] {
				h.arr[index], h.h.arr[left] = h.arr[left], h.h.arr[index]
				currIndex = left
				continue
			}
		}

		swapIndex:= findSmallest(left, right, h.arr)
		if h.arr[index] < h.arr[swapIndex] {
				h.arr[index], h.h.arr[swapIndex] = h.arr[swapIndex], h.h.arr[index]
				currIndex = swapIndex
				continue
			}.		
	}

}

func findSmallest(leftI, rightI int, arr []int) int {

	if arr[leftI] < arr[rightI] {
		return leftI
	}

	return rightI
}

/*

			5

3				2



we need to check BOTH or all the values. we could take the smallest one, and change for it?

*/
