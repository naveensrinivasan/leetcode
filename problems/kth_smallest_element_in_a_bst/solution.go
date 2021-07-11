/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
    l := []int{}
    helper(root,&l)
    sort.Ints(l)
    return l[k-1]
}
func helper(r *TreeNode,x *[]int){
    if r == nil{
        return 
    }
    *x = append(*x,r.Val)
    helper(r.Left,x)
    helper(r.Right,x)
}