package main

func main() {

}

type Node struct {
	Value int
	left  *Node
	right *Node
}

func CompareTRee(a *Node, b *Node) bool {
	// they are technically equal value
	if a == nil && b == nil {
		return true
	}

	// means that one of this is NULL, so they are not equal
	if a == nil && b == nil {
		return false
	}

	// both have an inside value so now we evaluate the value
	if a.Value != b.Value {
		return false
	}

	return CompareTRee(a.left, b.left) && CompareTRee(a.right, b.right)
}

func CompareTrees(a *Node, b *Node) bool {

	// base case to end the recursion if there is a match
	if a == nil && b == nil {
		return true
	}

	// base case to see if there is NO match
	if a == nil || b == nil {
		return false
	}

	// now evalute the value

	if a.Value != b.Value {
		return false
	}

	return CompareTrees(a.left, b.left) && CompareTRee(a.right, b.right)
}




func bst(node *Node, target int) bool {

	for node != nil {
		if node.Value == target {
			return true
		}

		if target > node.Value {
			node = node.right
		} else {
			node = node.left
		}
	}
	return false
}
