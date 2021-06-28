/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
    if root == nil{
        return []int{}
    }
    s := []*TreeNode{}
    nums := []int{}
    node := root
    for node != nil || len(s) > 0{
        if node != nil{
            s = append(s,node)
            node = node.Left
        }else{
            node = s[len(s) - 1]
            s = s[:len(s) - 1]
            nums = append(nums,node.Val)
            node = node.Right
        }
    }
    return nums
}