/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottom(root *TreeNode) [][]int {
     if root == nil{
        return nil
    }
    queue :=[]*TreeNode{root}
    result := [][]int{}
    counter :=0
    for len(queue) != 0{
        r := []int{}
        temp := []*TreeNode{}
        for len(queue) != 0{
            pop := queue[0]
            queue  = queue[1:]
            r = append(r,pop.Val)
            if pop.Left != nil{
                temp = append(temp,pop.Left) 
            }
            if pop.Right != nil{
                temp = append(temp,pop.Right) 
            }
        }
        result = append([][]int{r},result...)
        counter++
        queue = temp
    }
    return result
}