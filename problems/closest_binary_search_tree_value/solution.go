/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func closestValue(root *TreeNode, target float64) int {
    y := math.Abs(float64(root.Val) - target)
    z := &root.Val
    helper(root,target,z,&y)
    return *z
}

func helper(root *TreeNode,target float64,closest *int,x *float64){
    if root == nil{
        return
    }
    helper(root.Left,target,closest,x)
    helper(root.Right,target,closest,x)
    y := math.Abs(float64(root.Val) - float64(target))
    if y < *x{
        *x = y
        *closest = root.Val
    }
}