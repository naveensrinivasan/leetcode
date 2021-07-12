/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func preorder(root *Node) []int {
    l := []int{}
    helper(root,&l)
    return l
}

func helper(r *Node,l *[]int){
    if r == nil{
        return
    }
    *l = append(*l,r.Val)
    for _, item := range r.Children{
        helper(item,l)
    }
}