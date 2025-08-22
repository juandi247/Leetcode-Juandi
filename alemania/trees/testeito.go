package trees

import (
	"fmt"
)

// always shouyld be a stack
func DFS_Test(node *Node) {
	// We print Root, then Left

	if node == nil {
		return
	}
	fmt.Printf(node.value, ", ")
	for i, v := range node.children {
		DFS_Test(Node)
	}

}

/* Callstaclk
        A
      / | \
     B  C  D
    / \    |
   E   F   G



   Print A

   Starts the For of the cjildren

   PRint B

   Print E

   Return NIL because E doesnt have more children

   Go Back to B and continiue the for cycle (its alreayd printed so B doesnt gets printed anymore)

   Print F

   Return NIL because F doesnt have more children

	Finishes the For of B, then continues with the for of A


	Prints C

	Return NIL because no NODes

	Goes back again to the for of A to finish

	Prints D

	PRints G

	Return NIL because G deosnt have more

	Goes back to D
	Returns NIL because the is no more nodes

	Goes back to A, finishes the for loop

	END


*/



func invertBinaryTRee(node *Node){

	if node==nil{
		return
	}

	invertBinaryTRee(node.Left)
	invertBinaryTRee(node.Right)

	temp:= node.Left
	node.Left= node.Right
	node.Right=temp
}




// this always should be a queue
func BFS_Test(node *Node) {
	// this should be only for a case where its null but there is no recursion obviosuly
	if node == nil {
		return
	}

	/*
	   Steps:
	   Add the first one to the queue
	*/
	queue := []*Node{node}

	// Start iterating on the queue until its lenght is cero

	for len(queue) != 0 {
		//save the current just to print it (because we can not print the value (because its a Node Type struct))
		currentNOde := queue[0]
		fmt.Println(currentNOde.value)

		//pop the first element of the queue because its already printed
		// here we just have an slice from the 1st elemnt to the rest (inluding 1st)
		queue = queue[1:]

		// add the children of the node that we poped to the queue
		for _, v := range currentNode.children {
			queue = append(queue, v)
		}
	}

}
