/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isUnivalTree(root *TreeNode) bool {
    
    if root == nil{
        return false
    }
    
    x := root.Val
    nodes := []*TreeNode{}
    nodes = append(nodes,root)
    isInitialized := false
    for len(nodes) != 0{
        node := nodes[0]
        if !isInitialized{
            isInitialized = true
            
        }else{
            if x != node.Val{
                return false
            }
        }
        nodes = nodes[1:]
        
        if node.Left != nil{
            nodes = append(nodes,node.Left)
        } 
        if node.Right != nil{
            nodes = append(nodes,node.Right)
        } 
    }
    return true
}