/**
* Definition for singly-linked list.
 type ListNode struct {
*     Val int
*     Next *ListNode
* }
*/
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	prev := head
	cur, next := head.Next, head.Next
	for cur != nil {
		next = cur.Next
		next, cur.Next, prev, cur = cur.Next, prev, cur, next
	}

	head.Next = nil
	return prev
}
