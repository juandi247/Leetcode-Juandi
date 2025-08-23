package main

/*
Search a 2D Matrix
You are given an m x n 2-D integer array matrix and an integer target.

Each row in matrix is sorted in non-decreasing order.
The first integer of every row is greater than the last integer of the previous row.
Return true if target exists within matrix or false otherwise.

Can you write a solution that runs in O(log(m * n)) time?



Example 1:
Input: matrix = [[1,2,4,8],[10,11,12,13],[14,20,30,40]], target = 10
Output: true


1. Understaindg the porblem:
	We have a matrix, so in order to acces to the elements, we should have a nested for loop. Go for each array, and each element.

	We are told that ech row is sorted. and itself its technically in order. So the only thing that we should be looking too is the array[0] element
	of the array, and then go into the array itself.


2. Questions

	What if we dont find any target??
	Are they some constrains regarding some limits, like desbording?
	Limit like int64 or just integers of 32bits.

	Requirement is to get an o log(n) so its binary search


3. Prbolem slution initial

	I think the initial aproach that im gonna take is.
	Start by looking for the first value of the arrays.
	So we are going to intiailice 2 pinters, start of the first array, and end of the other array.

	We are going to make a binary search between them, to find if our value is located on which array

	Then we need to update our pointers to the array itself, so we can again, make binary search and then  using oit o log (n)


	Pseudocode should be something like

	PointerStart
	PointerEnd
arrayTosee:=0
	for i:= arrayOfArrays{

	if arrayofarays[i]== target{
	return TRUE}

	if arrayOfArrays[i]>target{
	arrayTosee:=i-1
	break
	}


	Then we can make a binary search in the other array


	The probelm with this is that the solucion is O(N * log(n))




	How can we do it just on o log n???



	PointerA
	PinterB

	Could be

	Start on the first arrrayOfarrays[0]
	End on the last arrayOfArrsy[lastArray]

	Take the middle point, and now compare

	If middlePoint == target{
	return true}

	if middlepoint < PointerA{
	PointerA = middlePoint
	}

	if middlePOint > PointerB{
PointerB = middlePoint
	}


	if middlePoint > PointA && middlePoint< POintB{
	arrayToSearch=middlePoint
	}




Again now a binary search but inside of it. So there should be like a double binary search, first to select the arrray where we are looking

Then another inside


*/


func DoubleAproach(numbers [][]int, target int) bool{

	// this aproach is even the one from neetcode xd.
	// just need to binary search between the first arrays


	// binary search in the selected array and thats it




	return false
}





/*
BETTER SOLUTION:

This solution implies more Math than anything
We treat the array as 1 Dimension, but without doing that. So we just do a binary search and thats it

The trick is:

We can calculate the length of the array by doing the area
Rows x Columns = Area

The thing is, with the area now we have different indexes of the array. So we have por example 3x 4 = 12
So for bianry search we have the pointer from 0 to len(array) = 0 to 11


So we calculate and we have 5 now. Okay but how can we acces to the position on the 2D dimensional array with this position????

From 5 we want the [row, column] index.

To now the row we use the formula of ----> Index / Column =  Row
To now the Column the formula is ----> Index % Colum = Colum


Thats it. Now we can search or look the value and thats it
*/

func MathAproach(numbers [][]int, target int) bool{


	rangeOfArray:= len(numbers)* len(numbers[0])
	startPointer:=0
	endPointer:=rangeOfArray


	for startPointer <= endPointer{
		middlePoint:= endPointer + (startPointer - endPointer)/2

		row:=middlePoint/ len(numbers[0])
		column:=middlePoint% len(numbers)
		value:= numbers[row][column]
		if value==target{
			return true
		}

		if value > target{
			startPointer= middlePoint
		}else{
			endPointer= middlePoint
		}

	}
	return false

}