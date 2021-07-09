/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
    g := []int{}
    helper(root,&g)
    return g
}
func helper(root *TreeNode,result *[]int){
    if root== nil{
        return
    }
    helper(root.Right,result)
    *result = append([]int{root.Val},*result...)
    helper(root.Left,result)
}