/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rangeSumBST(root *TreeNode, low int, high int) int {
    result:= 0
    helper(root,low,high,&result)
    return result
}

func helper(root *TreeNode,low,high int, result *int){
    if root == nil{
        return
    }
    if root.Val >= low && root.Val<=high{
            *result+=root.Val
    } 
    helper(root.Left,low,high,result)
    helper(root.Right,low,high,result)
}