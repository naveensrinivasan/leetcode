/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rangeSumBST(root *TreeNode, low int, high int) int {
    s := []*TreeNode{}
    s = append(s,root)
    result:= 0
    
    for len(s) != 0{
        node := s[0]
        s = s[1:]
        
        if node.Val >= low && node.Val<=high{
            result+=node.Val
        } 
        
        if node.Left != nil{
            s = append(s,node.Left)
        }
        
        if node.Right != nil{
            s = append(s,node.Right)
        }
    }
    return result
}