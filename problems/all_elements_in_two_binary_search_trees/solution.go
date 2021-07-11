/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
    l := []int{}
    helper(root1,root2,&l)
    sort.Ints(l)
    return l
}

func helper(r1 *TreeNode,r2 *TreeNode,l*[]int){
    if r1 == nil && r2 == nil{
        return
    }
    if r1 != nil{
        *l = append(*l,r1.Val)
    }
    if r2 != nil{
        *l = append(*l,r2.Val)
    }
    
    if r1 != nil && r2 != nil{
        helper(r1.Left,r2.Left,l)
        helper(r1.Right,r2.Right,l)
    } 
    
    if r1 == nil{
        helper(nil,r2.Left,l)
        helper(nil,r2.Right,l)
    }
    if r2==nil{
        helper(r1.Left,nil,l)
        helper(r1.Right,nil,l)
    }
}
