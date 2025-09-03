package main

import "fmt"

/*

Container With Most Water
You are given an integer array heights where heights[i] represents the height of the
i
t
h
i
th
  bar.

You may choose any two bars to form a container. Return the maximum amount of water a container can store.

APROACH

We are going to have two poitners

One goes to the start of the array. The other to the end of it.
We are NOT multiplying the numbers of the array. We are mutliplying the smallest number of
both pointers with the length in between both.
[1,7,2,5,4,7,3,4,6]


So in this case 1 and 6. The smallest is one, so we mutlpy
1* distance between them so its 7.



2.
Now we need to move one of the Pointers.
To move the pointers the best aproach should be:
See the Next value of the first pointer and the -1 value of the second pointer

Compare the two of them, and which is BIGGER, then we move them.const
So we can compare the biggest one easier.





STEPS:
1. Set the two pointers
	start and End
	Set the variable to return (INT)

2. Start the cycle and compare the value, save the biggest area into a variable

3. Decide where to move next
	Compare the two next values of the pointers, +1 on the first, and -1 on the second
	and update the pointer depending on who is bigger.

4. Continue till, the two pointers cross each other
*/



func ContainerWithMostWater(list []int) int{
	Start:=0
	End:= len(list)-1

	var BiggestArea int
	
	// [1,2,4,3]
	for Start<End{
		// obtain the biggest area
		// just obtain the biggest value betwen the current biggest area and
		// the smallest value in the heights
		CurrentArea:=min(list[Start], list[End])* End-Start
		if CurrentArea>BiggestArea{
			BiggestArea= CurrentArea
			fmt.Println()
		}
		

	
		// Move Pointers
		NExtStartValue:= list[Start+1]
		BeforeEndValue:= list[End-1]

		if NExtStartValue > BeforeEndValue{
			Start++
		}else{
			End--
		}
	}


	return BiggestArea



}