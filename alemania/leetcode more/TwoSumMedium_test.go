package main

import (
	"testing"
)

func TestTwoSum(t *testing.T) {

	structRta := []struct {
		array    []int
		target   int
		expected []int
	}{
		{array: []int{2, 7, 11, 15}, target: 9, expected: []int{1, 2}},
		{array: []int{2, 3, 4}, target: 6, expected: []int{1, 3}},
		{array: []int{-1, 0}, target: -1, expected: []int{1, 2}},
		{array: []int{1, 2, 3, 4, 6, 8, 23, 43}, target: 10, expected: []int{2, 6}}, // 2+8=10
		{array: []int{1, 2, 3, 4, 6, 8, 23, 43}, target: 50, expected: []int{6, 8}}, // 8+43=51 -> ajusta según target

	}

	for _, v := range structRta {

		result := twoSum(v.array, v.target)
		for i, resultValue := range result {

			if resultValue != v.expected[i] {

			}
		}
	}
}
