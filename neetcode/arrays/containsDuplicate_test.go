package arrays

import (
	"fmt"
	"strconv"
	"testing"
)

// Contains Duplicate
// Given an integer array nums, return true if any value appears more than once in the array, otherwise return false.

// Example 1:
// Input: nums = [1, 2, 3, 3]
// Output: true

// Example 2:
// Input: nums = [1, 2, 3, 4]
// Output: false

func repeatedValue(array []int) bool {

	miMap := make(map[int]int)
	for _, v := range array {
		_, exists := miMap[v]

		if exists {
			return true
		}
		miMap[v] = 1
	}
	return false
}

func TestDuplicate(t *testing.T) {

	mistruct := []struct {
		array  []int
		result bool
	}{
		{
			array:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 100000, 11111, 1},
			result: true,
		},
		{
			array:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			result: false,
		},
		{
			array:  []int{99, 100000, 1023223, 123323323, 99},
			result: true,
		},

		{
			array:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 100000, 11111, 1},
			result: true,
		},
	}

	
	for index, theStruct := range mistruct {
		t.Run("containsDuplicate test with array "+ strconv.Itoa(index), func(t *testing.T) {
			rv := repeatedValue(theStruct.array)
			if rv != theStruct.result {
				t.Errorf("No coincidieron las repuestas, mallllll")
			}
			fmt.Println("test PASSED")
		})
	}

}
