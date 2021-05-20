/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getLonelyNodes(root *TreeNode) []int {
    s := []*TreeNode{}
    result := []int{}
    s = append(s,root)
    
    for len(s) != 0{
        node := s[0]
        s = s[1:]
        count, val := 0,0
        
        if node.Left != nil{
            count++
            s = append(s,node.Left) 
            val = node.Left.Val
        }   
        
        if node.Right != nil{
            count++
            s = append(s,node.Right) 
            val = node.Right.Val
        }
        
        if count == 1{
            result = append(result,val)
        }
    }
    return result
}