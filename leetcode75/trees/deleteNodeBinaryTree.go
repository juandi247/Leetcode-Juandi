package main

type TreeNode struct {
	Val   int
	Right *TreeNode
	Left  *TreeNode
}

/*Aproach:

The probelm is to delete it, when it has sons, so we need to append or get a new head in that case, and
conncet the subtrees below etc.

BEST APROACH:
Is to find the Node
When we find it check if both of the sons are not NIL
In that case we need to do a MERGE trees.

How?? to replace our current value in this case

			40
		30	       	50
	25    34     45   60

We need to replace 40. So we can take the BIGGEST value on the left subtree and put it into 40.
and we would have a valid binary Tree

The other way is to take the SMALLEST value on the RIGHT sibutree, so it will be 45 and put it
instead of 40.
then we have a valid subtree and finish

*/

func deleteNode(root *TreeNode, key int) *TreeNode {

	if root == nil {
		return root
	}

	//find the KEY

	if root.Val > key {
		//15> 3 then we should take the LEFT part of our tree
		root.Left= deleteNode(root.Left *TreeNode, key int)
	}else if root.Val < key{
		// current is for exmaple 2<5 so we should go to the right of the tree because there are the bigger values
	root.Right= 	deleteNode(root.Right *TreeNode, key int)
	}else{
	//WE FOUND OUR VALUE!!!

		//If one of the sons is NIL, we just swap for the other and thats it
		if root.Left==nil{
			root= root.Right	
			return root
		}

		if root.Right==nil{
			root= root.Left	
			return root
		}



		//We have two sons that are VALID and no nil, so we need to find now the MaxValue on left subtree or minValue on
		//rightSubTreeroot

		current:= root
		for current.Right!=nil{
			current= current.Left
		}
		
		root.Val= current.Val

		root.Right= deleteNode(root.Right, current.Val)


	}




return root

}
