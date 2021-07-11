/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumOfLeftLeaves(root *TreeNode) int {
    sum := 0
    helper(root,false,&sum)
    return sum
}

func helper(root *TreeNode,isLeft bool,sum *int){
    if root == nil{
        return
    }
    
    if isLeft && root.Left == nil && root.Right == nil{
        *sum += root.Val
    }
    
    helper(root.Left,true,sum)
    helper(root.Right,false,sum)
}