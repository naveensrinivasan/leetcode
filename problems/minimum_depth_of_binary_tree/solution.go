/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
    type level struct{
        T *TreeNode
        L int
    }
    if root == nil{
        return 0
    }  
    q := []*level{}
    q = append(q, &level{T :root, L :1})
    min := 10000000
    
    for len(q) != 0{
        node := q[len(q)-1]
        q = q[:len(q)-1]
        
        if node.T.Left == nil && node.T.Right == nil && node.L < min{
            min = node.L
            continue
        }
        
        if node.T.Left != nil{
            q = append(q,&level{T :node.T.Left, L :node.L+1})
            
        }
        if node.T.Right != nil{
            q = append(q,&level{T :node.T.Right, L :node.L+1})
        }  
        
    }
    return min
}