package main
func moveZeroes(nums []int)  {
    if len(nums)<=1{
        return
    }

    left:=0

    for right:=0; right<=len(nums)-1; right++{
        if nums[right]!=0 {
            // swap
            nums[right], nums[left]= nums[left], nums[right]
            left++
        }
    
    }
