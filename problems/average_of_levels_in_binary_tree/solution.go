/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func averageOfLevels(root *TreeNode) []float64 {
   if root == nil{
        return nil
    }
    queue :=[]*TreeNode{root}
    result := []float64{}
    counter :=0
    for len(queue) != 0{
        temp := []*TreeNode{}
        sum := 0
        tempCounter := 0
        for len(queue) != 0{
            pop := queue[0]
            queue  = queue[1:]
            sum += pop.Val
            tempCounter++
            if pop.Left != nil{
                temp = append(temp,pop.Left) 
            }
            if pop.Right != nil{
                temp = append(temp,pop.Right) 
            }
        }
        result = append(result,float64(sum)/float64(tempCounter))
        counter++
        queue = temp
    }
    return result
}