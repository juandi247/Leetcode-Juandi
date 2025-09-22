package main

//https://leetcode.com/problems/symmetric-tree/

func SymetricTree(root *Node) bool {

	/*An initial though was to do it with BFS, and start comparating it over, eaiser.

	  BUT there is DFS APROACH: where we can do the same as compare a tree, but just compare left with right and right with left
	  Really simliar to comparing TWO trees. We are really comparing 2 subtrees. so its basically the same as SameTree execerise*/
	return recurse(root.Left, root.Right)
}

//just a helper func
func recurse(p, q *Node) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	left := recurse(p.Left, q.Right)  //Compare the left node of the LEFt tree, with the right node of the right Tree
	right := recurse(p.Right, q.Left) // do the same comparison

	return left && right
}
