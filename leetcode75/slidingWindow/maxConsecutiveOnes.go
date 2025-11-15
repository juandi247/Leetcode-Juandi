package main
/*
aproach:
count the nmuber of zeros, and if we are on the valid range, then we update our count.

If we find that countzeros bigger than valid

we just move our start pointer to the next found zero, so we avoid it and decrease the nmuber of seen and the window gets schrinked.

*/

func longestOnes(nums []int, k int) int {

    zeroCount:=0
    maxVal:=0

    start:=0
    end:=0

    for end<len(nums){
        if nums[end]==0{
            zeroCount++
        }

        // nos pasamos
        for zeroCount>k{
            if nums[start]==0{
                zeroCount--
            }
            start++
        }
        maxVal= max(maxVal, end-start+1)
        end++
    }


return maxVal

}
