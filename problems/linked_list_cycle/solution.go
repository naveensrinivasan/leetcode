/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
    if head == nil || head.Next == nil{
        return false
    }
    
    t,h := head.Next , head.Next.Next
    for t!=nil && h!=nil{
        if t == h{
            return true
        }
        t = t.Next
        if h.Next == nil{
            break
        }
        h = h.Next.Next
    } 
    return false
}