/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func searchBST(root *TreeNode, val int) *TreeNode {
    nodes := []*TreeNode{}
    
    nodes = append(nodes,root)
    
    for len(nodes) != 0{
        node := nodes[0]
        nodes = nodes[1:]
        if node.Val == val{
            return node     
        }
        if node.Left != nil{
            nodes = append(nodes,node.Left)
        }
        if node.Right != nil{
            nodes = append(nodes,node.Right)
        }
    }
    return nil 
}

