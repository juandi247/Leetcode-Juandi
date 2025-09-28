package main

//https://leetcode.com/problems/maximum-depth-of-binary-tree/


type *Node struct{
	Val int
	Left *Node
	Right *Node
}




func maxDepthBinaryTree(root *Node) int {

	// we do a DFS, because we need to see the depth

	if root == nil {
		return 0
	}

	leftSon := maxDepthBinaryTree(root.Left)
	rightson := maxDepthBinaryTree(root.Right)

	return 1 + max(leftSon, rightson)

}

/*
Callstack: it will start by the root
then goes to left
then on that node 9 we execute LEFT, because its nil, it returns 0.
Then the right again is nil so we have 0.
And then we should return the depth of the node 9. Because it was Zero the sons, we just do a
+1 and then return it. So that the ROOT, the HEAD, has the count of how much was on the left


ON the stack this would look like we executed already all the left depth and we return 1

So its left= 1
Right= execute the same thing on the right


It keeps going an thats it.


*/
