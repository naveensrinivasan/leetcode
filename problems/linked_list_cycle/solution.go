// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

func hasCycle(head *ListNode) bool {
	first, second := head, head
	for second != nil && second.Next != nil {
		first = first.Next
		second = second.Next.Next

		if first == second {
			return true
		}
	}
	return false
}
