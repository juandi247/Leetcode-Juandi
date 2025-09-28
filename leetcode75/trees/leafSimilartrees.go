package main

import "slices"

//https://leetcode.com/problems/leaf-similar-trees/editorial/?envType=study-plan-v2&envId=leetcode-75

/*
aproach:

We take two subtrees, or two new trees and we need to have an array or someway
to compare value per value.

We can do this in two ways:

1. Create two slices
	Make a recursion function that goes DFS and everytime appends a LEAF value into the array
	Do this for both of the Trees
	Make a comparison so it will be O(N) and thats it


2. Use the same recursive function, but make it so we compare every node, instead of creating
adiddtional space parts.


*/
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {

	list1, list2 := []int{}, []int{}

	traverse(&list1, root1)
	traverse(&list2, root2)

	return slices.Equal(list1, list2)
}

func traverse(array *[]int, node *TreeNode) {

	if node == nil {
		return
	}

	traverse(array, node.Left)
	traverse(array, node.Right)

	if node.Right == nil && node.Left == nil {
		*array = append(*array, node.Val)

	}
}
