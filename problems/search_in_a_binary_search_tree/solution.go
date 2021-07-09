/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func searchBST(root *TreeNode, val int) *TreeNode {
   return helper(root,val)
}

func helper(root *TreeNode,val int) *TreeNode{
    if root == nil{
        return nil
    }
    if root.Val == val{
        return root
    }
   if x := helper(root.Left,val); x != nil{
        return x
    }
    if x:= helper (root.Right,val); x != nil{
        return x
    }
    return nil
}

