import "strconv"
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type F struct{
   R *TreeNode
   stack string
}

func binaryTreePaths(root *TreeNode) []string {
    result := []string{}
    return helper(F{root,""},result)
}

func helper(pop F,result []string)[]string{
    
    arrow := ""
    if len(pop.stack) != 0{
        arrow = "->"
    }
    
    if pop.R.Left == nil && pop.R.Right == nil{
       return append(result,pop.stack + arrow+ strconv.Itoa(pop.R.Val))
    }
    
    if pop.R.Right != nil{
        result = helper(F{pop.R.Right,pop.stack + arrow+ strconv.Itoa(pop.R.Val)},result)     
    }
    
    if pop.R.Left != nil{
        result = helper(F{pop.R.Left,pop.stack + arrow+ strconv.Itoa(pop.R.Val)},result)     
    }
    return result
    
}
