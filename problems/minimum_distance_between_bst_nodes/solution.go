/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDiffInBST(root *TreeNode) int {
g := []int{}
    helper(root,&g)
    sort.Slice(g,func(i,j int)bool{
        return g[i] > g[j]
    })
    min := 100000000
    for i:=1;i<len(g);i++{
        x:= g[i-1] - g[i]
        if x < min{
            min = x
            if min == 1{
                return 1
            }
        }
    }
    return min
}

func helper(root *TreeNode,x *[]int){
    if root == nil{
        return 
    }
    *x = append(*x,root.Val)
    helper(root.Left,x)
    helper(root.Right,x)
} 