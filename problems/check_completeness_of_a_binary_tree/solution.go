/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isCompleteTree(root *TreeNode) bool {
    if root == nil{
        return false
    }
    q := []*TreeNode{root}   
    wasRightTreeFilled := true
    wasHead := true
    for len(q) > 0{
        
        pop := q[0]
        q = q[1:]
        isLeftPresent := false
        headHasLeftAlone := false
        if wasHead{
            if pop.Right == nil && pop.Left == nil{
                return true
            }
            if pop.Left != nil && pop.Right == nil{
                headHasLeftAlone = true
            }
            wasHead = false
        }
        if pop.Right != nil && pop.Left == nil{
            return false
        }
        if pop.Left!= nil{
           
            q = append(q,pop.Left)
            if pop.Left.Left != nil{
                if headHasLeftAlone{
                    return false
                } 
                if !wasRightTreeFilled{
                    return false
                }
            }
            if pop.Left.Left != nil && pop.Left.Right != nil{
                isLeftPresent = true
            }
        }
        if pop.Right != nil{
            q = append(q,pop.Right)
            if pop.Right.Left != nil && pop.Right.Right != nil{
                wasRightTreeFilled = true
            }else{
                wasRightTreeFilled = false
            }
            if !isLeftPresent && (pop.Right.Left != nil || pop.Right.Right != nil){
                return false
            }
        }
    }
    return true
}

    