/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func longestZigZag(root *TreeNode) int {

    leftRecursion:= recurse(root.Left, "l", 0)
    rightRecursion:= recurse(root.Right, "r", 0)


return max(leftRecursion, rightRecursion)


}


func recurse(root *TreeNode, movement string, length int) int{

    if root==nil{
        return length
    }

    var leftRecursion int
    var rightRecursion int


    if movement=="l"{
        leftRecursion= recurse(root.Left, "l", 0)
        rightRecursion= recurse(root.Right, "r",length+1)
        return max(leftRecursion, rightRecursion)
    }




    leftRecursion= recurse(root.Left, "l", length+1)
    rightRecursion= recurse(root.Right, "r",0)
    return max(leftRecursion, rightRecursion)
}

/*
APROACH:


We will start it by 0 from the ROOT.

and calculate the recursive from Left and right.
We obtain or caclulate the max from both calls, from the LEFT side and the RIGHT side.




IF our movement was to the left, we now do the same call, compare the max from the LEFT, and from the RIGTH.
BUTTTT in this case we recurse from the contrary position with a +1
and with the current position with a 0 indicating that is a starting point.



IF movement==LEFT{

recurse(node.Right, "r", length+1) 
recurse(node.Left, "l", 0) -> here is zero because its not a zig zag, but we start it  from our current node.
and return again the MAX of it

}


*/
