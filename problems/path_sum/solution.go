/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, targetSum int) bool {
        type F struct{
    R *TreeNode
    arr  []int
    val int
}
    if root == nil{
        return false
    }
    stack :=[]F{}
    stack = append(stack,F{R:root})
    for len(stack) != 0{
        pop := stack[0] 
        stack = stack[1:]
        
        if pop.R.Left == nil && pop.R.Right == nil{
             pop.val += pop.R.Val
             if pop.val == targetSum{
                 return true
             }
        }
        
        if pop.R.Right != nil{
            val := pop.val + pop.R.Val
            x := append(pop.arr,pop.R.Val)
            stack = append(stack,F{pop.R.Right,x,val})     
        }
        if pop.R.Left != nil{
            val := pop.val + pop.R.Val
            temp := []int{}
            temp = append(temp,pop.arr...)
            x := append(temp,pop.R.Val)
            stack = append(stack,F{pop.R.Left,x,val})         
        }
    }
    return false
    
}