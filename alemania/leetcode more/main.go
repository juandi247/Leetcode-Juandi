package main

func main() {

}

/*
So given this examples what i understand is that:
1. I need to search for two numbers on the array that summed are the target
2.Need to return an array of length 2, with the positions of them (+1)

Questions:
Important questions to ask:
There is always a solution? meaning that we will always find the number, if we dont find it then what should be returend?
Can i use repeated numbers?
Are there repeated nmubers?

Example 1:

Input: numbers = [2,7,11,15], target = 9
Output: [1,2]
Explanation: The sum of 2 and 7 is 9. Therefore, index1 = 1, index2 = 2. We return [1, 2].
Example 2:

Input: numbers = [2,3,4], target = 6
Output: [1,3]
Explanation: The sum of 2 and 4 is 6. Therefore index1 = 1, index2 = 3. We return [1, 3].
Example 3:

Input: numbers = [-1,0], target = -1
Output: [1,2]
Explanation: The sum of -1 and 0 is -1. Therefore index1 = 1, index2 = 2. We return [1, 2].

Posible aproaches:

1. My first idea should be compare all the numbers from starting
Mening that i would have two pointers:
first one would be on nums[0] and the other one on firstone+1.
It should be moving till we find the solution or a value that is bigger than the target itself.
If we dont find anything then we move the pointer +1 and start the second pointer again in +1

# This solutoin sounds meh, because would be technoally a nesteed for loop, on the worst case scenario so its not worhtit probably

2. Another aproach sounds to me that could be to search the pair of each value. Again  pointer to [0] but we would do
nums[0] - target = match.

And to search the match we can use (as the list is ordered, we can use binary search).

3. We can start by spliting the array into the posible answers. With binary search we can obtain the values that should be on the left
of the array meaingn <= target
we now have a simplified array (altohuht in worst case scenario this  would be unnecesary)

I will probably will start with number 2. but first i want to go trough it first.

there is a beter aproach-
The aproach of the 2nd is not that good, becasue we can still go for search and then evaluate, giving us O(n log n)

Can WE DO IT BEETTER???!!!!!

In O(N) with a two pointer aproach can be done:

We can start by getting the first pointer on to 0 the other at the end
If the sum of both is bigger, then we should move the one on the right.const
If the sum is smaller, then we should move the one on the left.
*/
func twoSum(numbers []int, target int) []int {

	pointerA := 0
	pointerB := len(numbers) - 1
	arrayFinal := [2]int{}

	for pointerA < pointerB {

		if numbers[pointerA]+numbers[pointerB] == target {
			arrayFinal[0] = pointerA + 1
			arrayFinal[1] = pointerB + 1
			return arrayFinal[:]
		}

		if numbers[pointerA]+numbers[pointerB] > target {
			pointerB--
		} else {
			pointerA++
		}

	}
	return arrayFinal[:]

}

// return the po sition of the elment if exits, if not return -1
func anotherBinarySearchxd(nums []int, target int) int {
	// binary search consist on, two pointers, start, end.
	// calculate the range. If the middle point is bigger than the target then just search the first part, and thsts it

	startPointer := 0
	endPointner := len(nums) - 1

	for startPointer <= endPointner {

		middlePoint := (startPointer + endPointner) / 2

		if nums[middlePoint] == target {
			return middlePoint
		}

		if nums[middlePoint] > target {
			endPointner = middlePoint - 1
		}

		if nums[middlePoint] < target {
			startPointer = middlePoint +1
		}
	}

	return -1
}
