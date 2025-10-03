package main

/*
in this case we only just make a DFS

we need to have always the BIGGEST value so far
and we need just to return a +1 or +0 depending on
if we find a valid bigger node or node
*/

func findGoodNodes(root *TreeNode) int {

	if root == nil {
		return 0
	}

	return recursiveDfs(root, root.Val)
}

func recursiveDfs(node *TreeNode, maxVal int) int {

	// always start with base case
	if node == nil {
		return 0
	}

	wasBigger := 0
	if node.Val > maxVal {
		wasBigger++
		maxVal = node.Val
	}

	//return wasBigger + recursiveDfs(node.Left, maxVal) + recursiveDfs(node.Right, maxVal)
	left := recursiveDfs(node.Left, maxVal)
	right := recursiveDfs(node.Right, maxVal)

	//if the value was smaller thant our maxSoFar then the sum is 0 + nextFunction
	// but if we find a valid bigger node, we jsut return +1
	return wasBigger + left + right

}
