/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
    l := []int{}
    if head == nil{
        return nil
    }
     node,result := head , &ListNode{}
    
    for node != nil{
         l,node = append(l,node.Val) , node.Next
    }
    sort.Ints(l)
    
    
    node = result
    for i:= 0;i<len(l);i++{
        node.Val = l[i]
        if i != len(l) -1{
            node.Next = &ListNode{}
            node = node.Next    
        } 
    }
    
    
    return result
}