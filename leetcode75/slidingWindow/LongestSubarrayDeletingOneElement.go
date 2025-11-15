
package main

/*
really similar apraoch, if we find a zero just means that is the one that we can delete.

we used end -start +1, but in this case the 
+1 is not needed becasue we obtain the subarray by deleting that zero, so just avoid that +1 and thats it.

Schrink or update the starting value when we see another zero, and thats it. 

MEDIUM HOHOH

*/
func longestSubarray(nums []int) int {
    start:=0
    end:=0
    zerosCount:=0
    maxVal:=0

    for end<len(nums){
        if nums[end]==0{
            zerosCount++
        }

        for zerosCount>1{
            if nums[start]==0{
                zerosCount--
            }
             start++
        }

        maxVal= max(maxVal, end - start)
        end++
    }
    return maxVal
}
