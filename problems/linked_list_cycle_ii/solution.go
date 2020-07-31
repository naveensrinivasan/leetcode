// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

func detectCycle(head *ListNode) *ListNode {
	first, second := head, head
	for second != nil && second.Next != nil {
		first = first.Next
		second = second.Next.Next

		if first == second {
			break
		}
	}
	if second == nil || second.Next == nil {
		return nil
	}

	for second = head; first != second; second, first = second.Next, first.Next {
	}
	return second
}
