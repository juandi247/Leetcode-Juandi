
package main


func RemoveDuplicates(nums []int) int {
    PointerA:= 0
    PointerB:= 1
	deletedCounter:=0

    for PointerB<len(nums){
		if nums[PointerA]==nums[PointerB]{
			deletedCounter++
			PointerB++
		}else{
			PointerA=PointerB
			PointerB++
		}
    }

    return len(nums) - deletedCounter
}