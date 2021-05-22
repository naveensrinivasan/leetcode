/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func countNodes(root *TreeNode) int {
    if root == nil{
        return 0
    }
    result := 0
    nodes := []*TreeNode{}
    nodes = append(nodes,root)
    
    for len(nodes) != 0{
        node := nodes[0]
        if node != nil{
            result++
        }
        nodes = nodes[1:]
        
        if node.Left != nil{
            nodes = append(nodes,node.Left)
        } 
        if node.Right != nil{
            nodes = append(nodes,node.Right)
        } 
    }
    return result
}