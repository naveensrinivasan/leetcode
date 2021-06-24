/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
    x := []int{}
    cur := head
    for cur!= nil{
        x = append(x,cur.Val)
        cur = cur.Next
    }
    if len(x) > 0{
        m := k % len(x)
        for i:=0;i<m;i++{
            x = append([]int{x[len(x)-1]},x[:len(x)-1]...)
        }
    }
    var prev *ListNode
    head = &ListNode{}
    curr := head
    isIn := false
    for i:=0;i<len(x);i++{
       isIn = true
       curr.Val = x[i] 
       l := &ListNode{}
       prev = curr
       curr.Next = l
       curr = curr.Next
    }
    if prev != nil{
        prev.Next = nil
    }
    if !isIn{
        return nil
    }
    return head
}