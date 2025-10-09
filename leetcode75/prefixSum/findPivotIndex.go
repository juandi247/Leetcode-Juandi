package main

func pivotIndex(nums []int) int {

	leftSum, rightSum := 0, 0

	for i := 1; i <= len(nums)-1; i++ {
		rightSum += nums[i]
	}

	for i := 0; i < len(nums)-1; i++ {
		if rightSum == leftSum {
			return i
		}
		leftSum += nums[i]
		rightSum -= nums[i+1]
	}

	if leftSum == 0 {
		return len(nums) - 1
	}

	return -1
}
