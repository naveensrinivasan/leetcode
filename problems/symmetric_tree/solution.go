/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
    if root == nil{
        return false
    }
    

    nodesL := []*TreeNode{}
    nodesR := []*TreeNode{}
    nodesL = append(nodesL,root.Left)
    nodesR = append(nodesR,root.Right)
    if root.Left == nil && root.Right == nil{
        return true
    }
    if root.Left == nil{
        return false
    }
    if root.Right == nil{
        return false
    }
    if root.Left.Val != root.Right.Val {
        return false
    }
    for len(nodesL) != 0 || len(nodesR) != 0{
        if len(nodesL) != len(nodesR){
            return false
        }
        l := nodesL[0]
        r := nodesR[0]
      
        
        nodesL = nodesL[1:]
        nodesR = nodesR[1:]
        
        if l.Val != r.Val{
            return false
        }
        
        if l.Left != nil && r.Right == nil || l.Left == nil && r.Right != nil && l.Right != nil && r.Left == nil{
            return false
        }else if l.Right == nil && r.Left != nil{
            return false
        }
        
        if l.Left != nil{
            nodesL = append(nodesL,l.Left)
        } 
        if r.Right != nil{
            nodesR = append(nodesR,r.Right)
        } 
        
        if l.Right != nil{
            nodesL = append(nodesL,l.Right)
        }
        if r.Left != nil{
            nodesR = append(nodesR,r.Left)
        }
        
        
    }
    return true
}
