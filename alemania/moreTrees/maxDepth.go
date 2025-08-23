package main

func searchMaxDepth(node *Node ) int{
	if node==nil{
		return 0
	}

	leftValue:= searchMaxDepth(node.left)
	righValue:= searchMaxDepth(node.right)

	if leftValue > righValue{
		return leftValue +1
	}

	return righValue+1
}



func EASY_Max_depth(node *Node) int{


// base casej
	if node==nil{
		return 0
	}

/* we just need to do this. Why?
We just return 0 when we find a null or the branch is terminated.
So when it pops up again its going to compare Max(0,0) = 0 and +1 is the node that has those children.
So when it pops again up, its going to compare now Max(1, if there is another node so 1) +1 = meaning 2
and if the root is up, so we compare again 1+ max(2, the value if there is another node) -> so 3 is the DEPTH EASY


*/
return 1 + max(EASY_Max_depth(node.left), EASY_Max_depth(node.right))

}


