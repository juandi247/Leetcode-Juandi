package mitrees


type Node struct{
	Value int
	Children []*Node
}


func Pre_order_search(node *Node, listNodes []int) []int{

	// base case: when the node that we are in is NIL
	if node==nil{
		return listNodes
	}

	
	// recursive path
	listNodes = append(listNodes, node.Value)
	for _, value:= range node.Children{
		Pre_order_search(value, listNodes)
	}

return listNodes

}