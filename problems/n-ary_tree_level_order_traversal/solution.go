/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func levelOrder(root *Node) [][]int {
    if root == nil{
        return nil
    }
    queue :=[]*Node{root}
    result := [][]int{}
    counter :=0
    for len(queue) != 0{
        r := []int{}
        temp := []*Node{}
        for len(queue) != 0{
            pop := queue[0]
            queue  = queue[1:]
            r = append(r,pop.Val)
            for _, item := range pop.Children{
                temp = append(temp,item)       
            }
        }
        result = append(result,r)
        counter++
        queue = temp
    }
    return result
}