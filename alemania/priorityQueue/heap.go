package main


/* 
HEAP
Always divided in MinHeap and MaxHeap
Just easier to get it the minHeap, means that the ROOT is going to be the minimum value of the tree

          5
        /   \
       10    15
      /  \   /  \
     20  30 25  40



The MAx heap, on the root is going to be the max value

          50
        /    \
      30      20
     /  \    /  \
   10   15  5    2



IMPORTANT -> The heap is not implemneted as a node based tree
Its really idficult to for example add a new node ,on the node that is 10 for example
How to we now to add to that? if we have a node based, we would need to make a DFS to reach the 
most left node, etc, etc. 

SO HEAPS ARE ALWAYS IMLPEMNETED AS Arrays (OR SLICES)
[50, 30, 20, 10, 15, 5, 2]
 0    1   2   3   4  5  6


-----IMPORTNATISISMOOOOO-----
LEFT CHILDREN:  2*indice + 1
RIGHT CHILDREN: 2* indice +1

PARENT:= (i-1)/2


*/




type MinHeap struct{
	length int
	data []int
}


func giveMeHeap() *MinHeap{
	return &MinHeap{}
}


func (m *MinHeap) insert(value int){
	/* Steps on the insert

	1. Add verification if legnth is 0 then just add it EZ.

	2. Append the new value to the end of the array. (a leaf)

	3. Bubble up or minheap it up
		3.1 Obtain the current position (would be easier with the length)
		3.2 Calculate the parent of my current position with ---->   (i-1)/2
		3.2 Compare the parent with my actual value, we will do this until we reach that the parent is bigger or equal than the 
			son. (minheap)
		3.2 When the value is not bigger, of the parent, means that we need to swap, so we swap both. 

		Thats it

	*/
	
	m.data= append(m.data, value)
	m.length++
	addedPosition:= m.length -1

	for addedPosition>0{ 
		parent:= (addedPosition-1)/2
		// minheap- if the parent is smaller than the son (RIGHT)
		if m.data[parent]<m.data[addedPosition]{
			break
		}

		// swap the values
		// temp:=m.data[parent]
		// m.data[parent]=value
		// m.data[addedPosition]=temp

		// another way to swap values
		m.data[parent], m.data[addedPosition]= m.data[addedPosition], m.data[parent] 
		// switch the position of our inserted value, to the parent one.
		addedPosition= parent
	}
}











/*
To delete the elemnt of the heap is not as easyu uas to pop the first one and thats it. The thing is to reorder
the heap becasue that deletion causes a reordee of the root

1. To delete the ROOT, we swap the value of it with the last element

2. We have a deleted root and a valid tree, but we need to minheap the value again.const

3. Compare both of the children of the root, get the MIN value of both of them, becasue we should have as the root the smallest vlue between both

4. Swap with the root, and continue until the comparison is not valid

*/
func (m *MinHeap) deleteRoot(){

	// swap the values
	m.data[0]= m.data[m.length-1]

	// delete now the last value
	m.data= m.data[:m.length-1]
	m.length--


	currentIndex:=0
	
	for{
		leftChildrenIndex:=currentIndex*2+1
		rightChildrenIndex:=currentIndex*2+2
		sonIndex, minValue:=0,0
	
		
		// get the value of the children with the smalles value and its index
		if m.data[leftChildrenIndex]< m.data[rightChildrenIndex]{
			minValue=m.data[leftChildrenIndex]
			sonIndex= leftChildrenIndex
		}else{
			minValue=m.data[rightChildrenIndex]
			sonIndex= rightChildrenIndex
		}
		



		if m.data[currentIndex]> minValue{
			// swap
			m.data[currentIndex], m.data[sonIndex]= m.data[sonIndex], m.data[currentIndex]
			currentIndex=sonIndex
		}else{
			break
		}
	}



}