/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func postorder(root *Node) []int {
    l := []int{}
    helper(root,&l)
    return l
}

func helper(root *Node,g *[]int){
    if root == nil{
        return
    }
    for _, item := range root.Children{
        helper(item,g)       
    }
    *g = append(*g,root.Val)
}