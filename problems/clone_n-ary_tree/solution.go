/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func cloneTree(root *Node) *Node {
    if root==nil{
        return nil
    }
    c := &Node{Val:root.Val}
    helper(root,c)
    return c
}

func helper(r,c *Node){
    if len(r.Children) == 0{
        return
    }
    for i,item := range r.Children{
        c.Children = append(c.Children,&Node{Val:item.Val})
        helper(r.Children[i],c.Children[i])
    }
}