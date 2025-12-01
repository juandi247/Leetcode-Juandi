

/* 
Just do an XOR with the current number and the next.

If we reached a repetede number fo rexmaple we have
4-> 101
2->10
2->10


Then the XOR will be 4 and 2 at start: 101
but on the next 2, will be 10 XOR 111 so it will be 101 
*/
func singleNumber(nums []int) int {


if len(nums)==1{
    return nums[0]
}

prevNumber:=nums[0]
for i:=1;i<=len(nums)-1; i++{
    // comparison XOR currNumber with the prev
    currNumber:=nums[i]
    prevNumber^=currNumber
}

return prevNumber
    
}


/*
We can NOt create a map to have the count

we must do it in o(N)

Two pointers aproach?

When we reach a new number
If the starting value



 */
