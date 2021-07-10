/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findMode(root *TreeNode) []int {
    m := make(map[int]int)
    r := []int{}
    max := 0
    helper(root,&m,&max)
    for k,v := range m {
        if v == max-1{
            r = append(r,k)
        }
    }
    return r
}

func helper(root *TreeNode,m *map[int]int, max *int){
    if root == nil{
        return
    }
    x := *m
    x[root.Val]++
    if x[root.Val]+1 > *max{
        *max = x[root.Val]+1
    }
    helper(root.Left,m,max)
    helper(root.Right,m,max)
}