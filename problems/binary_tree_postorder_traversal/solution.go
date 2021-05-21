/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
    s := []*TreeNode{}
    nums := []int{}
    if root == nil{
        return []int{}
    }
    s = append(s,root)
    
    for len(s) != 0{
        node := s[len(s)-1]
        s = s[:len(s)-1]
        
        if node!= nil{
            nums = append([]int{node.Val},nums...)    
        }
        if node.Left != nil{
            s = append(s,node.Left)
        }
        if node.Right != nil{
            s = append(s,node.Right)
        }
        
              
        
    }
    return nums
}