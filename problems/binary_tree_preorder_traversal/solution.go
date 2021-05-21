/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
    s := []*TreeNode{}
    nums := []int{}
    if root == nil{
        return []int{}
    }
    s = append(s,root)
    
    for len(s) != 0{
        node := s[len(s)-1]
        s = s[:len(s)-1]
        //fmt.Println(node)
        if node!= nil{
            nums = append(nums,node.Val)    
        }
        
        if node.Right != nil{
            s = append(s,node.Right)
        }
        if node.Left != nil{
            s = append(s,node.Left)
        }
        
        
        
        
        
        
    }
    return nums
}