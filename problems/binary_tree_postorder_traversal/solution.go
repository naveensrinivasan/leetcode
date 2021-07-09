/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
    nums := []int{}
    helper(root,&nums)
    return nums
}
func helper(root *TreeNode,g *[]int){
    if root == nil{
        return
    }
    *g = append([]int{root.Val},*g...)
    helper(root.Right,g)
    helper(root.Left,g)
}