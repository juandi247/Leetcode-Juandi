package main

//https://leetcode.com/problems/same-tree/description
func isSameTree(p, q *Node) bool {

	/*We should retunr true or false if our trees are the same, but not same data, SAME structure and data
	we should do a DFS, so we go to the depth, and evaluate the strcuture, and acomplish base cases
	*/

	//if both nodes are nil, means that the strcuture is equal, return TRUE (and we reached the most depth
	if p == nil && q == nil {
		return true
	}

	//if one of them is nil, means that its DIFERENT, so we return FALSE
	if p == nil || q == nil {
		return false
	}

	//if the values are different, return FALSE
	if p.Val != q.Val {
		return false
	}

	//If the values are EQUAL, and the node is NOT NIL, we continue recursing using DFS
	left := isSameTree(p.Left, q.Left)
	right := isSameTree(p.Right, q.Right)

	//this compares the left and right, if one of them is false, then means that it returns false
	return left && right

}
