/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
    type level struct{
        T *TreeNode
        L int
    }
    if root == nil{
        return 0
    }  
    q := []*level{}
    q = append(q, &level{T :root, L :1})
    max := 0
    
    for len(q) != 0{
        node := q[len(q)-1]
        q = q[:len(q)-1]
        if node.L > max{
            max = node.L
        }
        if node.T.Left != nil{
            q = append(q,&level{T :node.T.Left, L :node.L+1})
            
        }
        if node.T.Right != nil{
            q = append(q,&level{T :node.T.Right, L :node.L+1})
        }
        
    }
    return max
}