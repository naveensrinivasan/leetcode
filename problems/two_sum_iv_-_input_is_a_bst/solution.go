/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findTarget(root *TreeNode, k int) bool {
    m := make(map[int]bool)
    x := false
    helper(root,k,m,&x)
    return x
}

func helper(root *TreeNode,k int,m map[int]bool, r *bool){
    if root == nil {
        return 
    }
    
    if _,ok := m[k-root.Val];ok{
        *r = true
        return
    }
    m[root.Val]=true
    helper(root.Left,k,m,r)
    helper(root.Right,k,m,r)
}
