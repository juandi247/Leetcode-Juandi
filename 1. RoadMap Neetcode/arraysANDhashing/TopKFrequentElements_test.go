package main

import (
	"fmt"
	"testing"
)

/*
Top K Frequent Elements
Given an integer array nums and an integer k, return the k most frequent elements within the array.

The test cases are generated such that the answer is always unique.

You may return the output in any order.

Example 1:

Input: nums = [1,2,2,3,3,3], k = 2

Output: [2,3]
Example 2:

Input: nums = [7,7], k = 1

Output: [7]
*/

/*
Aproach:
In this case when we talk aboout TOPK elements, we should think about a Heap it could be usefull
we can put a maxheap, becasue we need to obtain the most appeared elements, and then we traverse the heap


SORTING AND ALL THAT
1. Save the amount of times that the elements appeared
2. Traverse the map and depending of the threshold, we obtain the elements with most appearances
3. Make a bubble sort or whatever to update

HEAP WOW:
The solution will be to create a minheap - just an array of size of the answer

We need to FILL a minHeap
When its already filled, we need to check the ROOT, if the new element is bigger thant the ROOT
we need to delete the root, by changing it (NOT with the leaf) but with the new element bigger thant the previous root and
bubble down to its position.




*/

func deleteFromMinHeap(heap []int) {

	currentIndex := 0
	for currentIndex < len(heap) -1 {
		leftC := currentIndex*2 + 1
		rightC := currentIndex*2 + 2

		sonIndex, minValue := 0, 0



		if leftC>len(heap)-1 || rightC>len(heap)-1{
			break
		}

		if heap[leftC] > heap[rightC] {
			minValue = heap[leftC]
			sonIndex = leftC
		} else {
			minValue = heap[rightC]
			sonIndex = rightC
		}

	



		
		if heap[currentIndex]<= minValue{
			break
		}


		// swap if its bigger our current value than the son value (minheap)
		heap[currentIndex], heap[sonIndex] = heap[sonIndex], heap[currentIndex]
		currentIndex = sonIndex
	
	}

}

func TopKFrequentElements(nums []int, k int) []int {

	miMap := make(map[int]int)
	for _, v := range nums {
		miMap[v]++
	}

	minHeap := []int{}

	for key, value := range miMap {

		if len(minHeap) < k {
			minHeap = append(minHeap, key)
			fmt.Println("we appned to the heap",key )
		} else {
			if value > minHeap[0] {
				minHeap[0] = value
				fmt.Println("heap current , ", minHeap[:])
				deleteFromMinHeap(minHeap)
			}

		}

	}

	return minHeap
}

func Test_Exercise(t *testing.T) {
	tests := []struct {
		name     string
		array    []int
		k        int
		expected []int
	}{
		{
			name:     "Caso básico",
			array:    []int{1, 2, 2, 3, 3, 3},
			k:        2,
			expected: []int{2, 3}, // orden puede variar
		},
		{
			name:     "Un solo elemento repetido",
			array:    []int{7, 7, 7, 7},
			k:        1,
			expected: []int{7},
		},
		{
			name:     "Todos con misma frecuencia",
			array:    []int{1, 2, 3, 4},
			k:        2,
			expected: []int{1, 2}, // o cualquier combinación de dos
		},
		{
			name:     "K igual al número de elementos únicos",
			array:    []int{1, 2, 2, 3},
			k:        3,
			expected: []int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := TopKFrequentElements(tt.array, tt.k)

			// Como el orden puede variar, usamos mapas para comparar
			expectedMap := make(map[int]bool)
			for _, v := range tt.expected {
				expectedMap[v] = true
			}

			for _, v := range result {
				if !expectedMap[v] {
					t.Errorf("para %v, esperado %v, pero obtuve %v", tt.array, tt.expected, result)
				}
			}
		})
	}
}
