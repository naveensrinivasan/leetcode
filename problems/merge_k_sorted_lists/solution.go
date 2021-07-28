/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
    l := []int{}
    for _,item := range lists{
        cur := item
        for cur != nil{
            l = append(l,cur.Val)
            cur = cur.Next
        }
    }
    sort.Ints(l)
    if len(l) == 0{
        return nil
    }
    head := &ListNode{Val:l[0]}
    cur := head
    for i:=1;i<len(l);i++{
        n := &ListNode{Val:l[i]}
        cur.Next = n
        cur = cur.Next
    }
    return head
}