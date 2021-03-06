/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    if head == nil{
        return nil
    }
    var prev,next,cur *ListNode
    cur = head
    for cur != nil{
        next = cur.Next
        cur.Next = prev
        prev = cur
        cur = next
    }
    return prev
}