// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
		if slow == fast {
			break
		}
	}
	if fast == nil || fast.Next == nil {
		return nil
	}
	for fast = head; slow != fast; slow, fast = slow.Next, fast.Next {
	}
	return slow
}
