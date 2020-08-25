/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Val: 0, Next: head}
	fast, slow := dummy, dummy

	for fast.Next != nil {
		fast = fast.Next
		n = n - 1
		if n < 0 {
			slow = slow.Next
		}
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}
