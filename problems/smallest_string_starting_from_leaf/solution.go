import "fmt"
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
   stack []string
}

func smallestFromLeaf(root *TreeNode) string {
    r := []string{}
    r = helper(F{root,[]string{}},r)
    sort.Strings(r)
    return r[0]
}


func helper(pop F,result []string)[]string{
    
    if pop.R.Left == nil && pop.R.Right == nil{
        temp := []string{}
        temp = append(temp,pop.stack...)
        temp = append(temp,string(pop.R.Val+97))
        var sb strings.Builder
        for i:=len(temp)-1;i>=0;i--{
            sb.WriteString(temp[i])
        }
        
    return append(result,sb.String())
    }
    
    if pop.R.Right != nil{
        temp := []string{}
        temp = append(temp,pop.stack...)
        temp = append(temp,string(pop.R.Val+97))
        result = helper(F{pop.R.Right,temp},result)     
    }
    
    if pop.R.Left != nil{
        temp := []string{}
        temp = append(temp,pop.stack...)
        temp = append(temp,string(pop.R.Val+97))
        result = helper(F{pop.R.Left,temp},result)     
        
    }
    return result
    
}
