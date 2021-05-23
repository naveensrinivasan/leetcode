/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
    if p == nil && q == nil{
        return true
    }
    if p == nil || q == nil{
        return false
    }
    
    firstList := []*TreeNode{}
    secondList := []*TreeNode{}
    
    firstList = append(firstList,p)
    secondList = append(secondList,q)
    
    for len(firstList) != 0 || len(secondList) != 0{
        first := firstList[0]
        firstList = firstList[1:]
        
        second := secondList[0]
        secondList = secondList[1:]
        
        
        f,s := -100000, -100000
        if first == second{
            continue
        }
        
        
        if first != nil{
            f = first.Val
        }
        if second != nil{
            s = second.Val
        }
        
        if f != s {
            return false
        }
        
        if first.Left != nil{
            firstList = append(firstList,first.Left)
        }else{
            firstList = append(firstList,nil)
        }
        
        if first.Right != nil{
            firstList = append(firstList,first.Right)
        }else{
            firstList = append(firstList,nil)
        }
        
        if second.Left != nil{
            secondList = append(secondList,second.Left)
        }else{
            secondList = append(secondList,nil)
        }
        
        if second.Right != nil{
            secondList = append(secondList,second.Right)
        }else{
            secondList = append(secondList,nil)
        }
        
    }
    return true
}