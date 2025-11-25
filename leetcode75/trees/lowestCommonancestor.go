
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
  return recurse(root,p,q)
}


func recurse(node, p,q *TreeNode) *TreeNode{

// base cases
// we reached the bottom and we didnt find our node.
if node==nil{
    return nil
}

// we found one of the nodes
if node==p || node==q{
    return node
}



left:= recurse(node.Left, p,q)
right:= recurse(node.Right,p,q)

// means that we have one of the nodes its on the left side, and the other is on the right side, so the only ancesotr is the ROOT
if left!=nil && right!=nil{
    return node
}


// if our value on the left is NIL, means that both of the values are on the RIGHT part of the tree
// since we have our base case that takes the FIRST noticed node, means that below that node, is the other one,
// than we can just return it
if left==nil{
    return right
}

// the same but on the other side
return left

}
