/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
    if root == nil{
        return nil
    }
    
    node := root
    nodes := []*TreeNode{}
    nodes = append(nodes,root)
    for len(nodes) != 0 {
        node = nodes[0]
        nodes = nodes[1:]
        node.Left,node.Right = node.Right,node.Left
        
        if node.Left != nil{
            nodes = append(nodes,node.Left)
        }
        
        if node.Right != nil{
            nodes = append(nodes,node.Right)
        }
    }
    return root
}