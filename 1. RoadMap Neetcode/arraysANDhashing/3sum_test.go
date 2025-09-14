package main

import (
	"fmt"
	"slices"
	"testing"
)

/* this is supposedly an arrays and hashing but the best aproach is two poitners haha



BRUTE FORCE:
We are going to evaluate for each value I, the values I+1 and I+2, and within
every iteration, we update the i+1 and i+2, with an ++. so we evaluate everythinh
Repeating this with every part, would give us an O(N**3)

We evaluate
I -> with J, and K
We need to evaluate each value I, J, with each K

So evaluate all the array.
On this evaluate all the posible J combinations
Evaluate from J combiations the K combinations
TRIPLE EVALUATION O(n**2)


TWO POINTERS
This aproach requires first a sorted array
So we sort it with an O (n log n)

Then with the array sorted, we still need to check out every combination posisble
the thing is, in this case would be O(n**2) because we evaluate is with pointers, and not
with a double.

We first evaluate all the posible combinations for each position, so its O(N) for each value
Inside we traverse the array just 1, because we already now its sorted, so we dont need
to evaluate for each value inside, another value
*/

func ThreeSum(nums []int) [][]int {
    slices.Sort(nums)
    antwort := [][]int{}

    for i := 0; i < len(nums)-2; i++ {
        // evitar duplicados de i
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }

        start := i + 1
        end := len(nums) - 1

        for start < end {
            sum := nums[i] + nums[start] + nums[end]

            if sum == 0 {
                newArray := []int{nums[i], nums[start], nums[end]}
                antwort = append(antwort, newArray)

                start++
                end--

                // evitar duplicados en start
                for start < end && nums[start] == nums[start-1] {
                    start++
                }
                // evitar duplicados en end
                for start < end && nums[end] == nums[end+1] {
                    end--
                }

            } else if sum < 0 {
                start++
            } else {
                end--
            }
        }
    }

    return antwort
}




func Test_ThreeSum(t *testing.T) {
    tests := []struct {
        array    []int
        expected [][]int
    }{
        {
            array:    []int{-1, 0, 1, 2, -1, -4},
            expected: [][]int{{-1, -1, 2}, {-1, 0, 1}},
        },
        {
            array:    []int{0, 1, 1},
            expected: [][]int{},
        },
        {
            array:    []int{0, 0, 0},
            expected: [][]int{{0, 0, 0}},
        },
        {
            array:    []int{-2, 0, 1, 1, 2},
            expected: [][]int{{-2, 0, 2}, {-2, 1, 1}},
        },
        {
            array:    []int{},
            expected: [][]int{},
        },
    }

    for _, tt := range tests {
        result := ThreeSum(tt.array)
        if !compareTriplets(result, tt.expected) {
            fmt.Println("error en el input", tt.array, "obtuvimos fue", tt.expected, "\n ---------------------- ")
            t.Errorf("For input %v, expected %v but got %v", tt.array, tt.expected, result)
        }
    }
}




func compareTriplets(a, b [][]int) bool {
    if len(a) != len(b) {
        return false
    }

    aMap := make(map[string]bool)
    for _, trip := range a {
        key := fmt.Sprintf("%v", trip)
        aMap[key] = true
    }

    for _, trip := range b {
        key := fmt.Sprintf("%v", trip)
        if !aMap[key] {
            return false
        }
    }
    return true
}


112  78   101   119
56   98   114    49








for i:=0; i<len(arr)/2; i++{
    for j:=0; j<len(arr)/2; j++{
        max:= blalbal



    }



}