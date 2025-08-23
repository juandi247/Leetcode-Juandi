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