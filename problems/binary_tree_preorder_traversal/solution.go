/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
    g := []int{}
    helper(root,&g)
    return g
}
func helper(root *TreeNode,result *[]int){
    if root== nil{
        return
    }
    *result = append(*result,root.Val)
    helper(root.Left,result)
    helper(root.Right,result)
}