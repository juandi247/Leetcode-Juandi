package main

import "fmt"

func main() {

	testList:= []int{1,7,2,5,4,7,3,6}
	fmt.Println("RESULTADO: " , ContainerWithMostWater(testList))

	testList2:= []int{2,2,2}
	fmt.Println("RESULTADO 2: " , ContainerWithMostWater(testList2))


}

type Node struct {
	left  *Node
	right *Node
	value int
}

// binary Tree Preorder traversal
// go to left but print frist the Roots

func dfs(head *Node) {
	if head==nil{
		return
	}
	fmt.Println(head.value)
	dfs(head.left)
	dfs(head.right)	
}



/*we have now 2 Linked Lists.
We are going to have 2 Pointers. Pointing into the first Node of each List
1(P),1,5,6,7
1(P),1,2,3
What we should do is

prestep-> we can have 2 options here, a dummy Node that starts the result, or 
we can make an initial comparation to start it, the best here is the dummy for me.

1. Compare the two pointers
2. IF the list1 <= List2 node -> then we add to our result Current.Next= List1
else -> current.Next= List2

3. Now current should be updated too, so we make Current = current.Next. so we can continue the list


EDGE CASE: 
The is an edge case, for example, if we have nothing to compare anymore, we can just add the node
and thats it?
if one list is bigger than the other, one, nd we reach a comparation of 4<= nil, than we can just add all of hte elemnts.

The easy thing should be, evaluate the comparison, only if both are NOT NIL
if this isnt acommpilshed, means that one of them or both are like that and so we can add the other.


Return??? we should always return the DUMMY or head.NEXT, that is our first value

*/




