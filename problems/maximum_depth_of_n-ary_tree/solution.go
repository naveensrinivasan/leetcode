/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

type level struct{
    T *Node
    L int
}
func maxDepth(root *Node) int {
    max := 1
    if root == nil{
        return 0    
    }
    helper(level{root,1},&max)
    return max
}

func helper(l level,max *int){
    
    if l.L > *max{
        *max = l.L
    }
    
    for _,item := range l.T.Children{
        helper(level{item,l.L+1},max)
    }
    if len(l.T.Children) == 0{
        return
    }
}