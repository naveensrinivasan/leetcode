/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
    l :=[]int{}
    r :=[]int{}
    helper(root1,root2,&l,&r)
    return reflect.DeepEqual(l,r)
    
}

func helper(r1,r2 *TreeNode,l,r *[]int){
    if r1 == nil && r2 == nil{
        return
    }
    if r1 != nil && r1.Left == nil && r1.Right == nil{
        *l = append(*l,r1.Val)
    }
    if r2 != nil && r2.Left == nil && r2.Right == nil{
        *r = append(*r,r2.Val)
    }
    if r1 != nil && r2 != nil{
        helper(r1.Left,r2.Left,l,r)
        helper(r1.Right,r2.Right,l,r)
    }
    if r1 == nil{
        helper(nil,r2.Left,l,r)
        helper(nil,r2.Right,l,r)
    }
    if r2 == nil{
        helper(r1.Left,nil,l,r)
        helper(r1.Right,nil,l,r)
    }
}