/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type F struct{
    V int
    L int
}
func deepestLeavesSum(root *TreeNode) int {
    stack := []F{}
    helper(root,0,&stack)
    sort.Slice(stack,func(i,j int)bool{
        return stack[i].L> stack[j].L
    })
    max := stack[0].L
    sum := 0
    for _,item := range stack{
        if item.L < max{
            break
        }
        sum += item.V
    }
    return sum
    
}
func helper(root *TreeNode,l int,s *[]F)[]F{
    l += 1
    f1 := F{V:root.Val,L:l,}
    *s = append(*s,f1)
    
    if root.Left != nil{
        *s =helper(root.Left,l,s)
    }
    if root.Right != nil{
        *s= helper(root.Right,l,s)
    }
    return *s
}