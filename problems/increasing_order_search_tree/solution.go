/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func increasingBST(root *TreeNode) *TreeNode {
    nums := []int{}
    nodes := []*TreeNode{}
    
    nodes = append(nodes,root)
    
    for len(nodes) > 0{
        node := nodes[0]
        nodes = nodes[1:]
        nums = append(nums,node.Val)
        
        if node.Left != nil{
            nodes = append(nodes,node.Left)
        }
        
        if node.Right != nil{
            nodes = append(nodes,node.Right)
        }
    }
    
    sort.Ints(nums)
    head := &TreeNode{}
    t := head
    
    for i :=0;i<len(nums)-1;i++{
        t.Val = nums[i] 
        newTree := &TreeNode{}
        t.Right = newTree
        t = newTree
    }
    
    t.Val = nums[len(nums)-1]
    
    return head
}