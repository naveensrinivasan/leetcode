import "strconv"
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func binaryTreePaths(root *TreeNode) []string {
    type F struct{
       R *TreeNode
       stack string
    }
    stack :=[]F{}
    stack = append(stack,F{R:root})
    result := []string{}
    for len(stack) != 0{
        pop := stack[0] 
        stack = stack[1:]
        
        arrow := ""
        if len(pop.stack) != 0{
            arrow = "->"
        }
        if pop.R.Left == nil && pop.R.Right == nil{
            result = append(result,pop.stack + arrow+ strconv.Itoa(pop.R.Val))
            continue
        }
        
        if pop.R.Right != nil{
            stack = append(stack,F{pop.R.Right,pop.stack + arrow+ strconv.Itoa(pop.R.Val)})     
        }
        if pop.R.Left != nil{
            stack = append(stack,F{pop.R.Left,pop.stack + arrow +strconv.Itoa(pop.R.Val)})     
        }
    }
    
    return result
    
    
}
