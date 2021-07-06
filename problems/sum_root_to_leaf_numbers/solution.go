/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {
       type F struct{
        K *TreeNode
        V int
    }  
    result := 0
    cur := root
    l := []F{}
    l = append(l,F{cur,cur.Val})
    
    for cur!= nil && len(l) > 0{
        cur = l[0].K
        
        y := 0
        if cur != nil{
            y = l[0].V
        }
        l = l[1:]
        if cur.Left != nil{
            node := cur.Left
            l = append(l, F{node,y * 10 + node.Val})
            
        }
        if cur.Right != nil{
            node := cur.Right
            l = append(l, F{node,y * 10 + node.Val})
        }
        if cur.Left == nil && cur.Right == nil{
            result += y
        }
    }
    return result
} 
