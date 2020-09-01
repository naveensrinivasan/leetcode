/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	prev := result
	for l1 != nil || l2 != nil {
		if l1 != nil {
			a := l1.Val
			if l2 != nil {
				b := l2.Val
				if a < b {
					prev.Next = &ListNode{Val: a}
					prev = prev.Next
					if l1 != nil {
						l1 = l1.Next
					}
				} else {
					prev.Next = &ListNode{Val: b}
					prev = prev.Next
					if l2 != nil {
						l2 = l2.Next
					}
				}
			} else {
				prev.Next = &ListNode{Val: a}
				prev = prev.Next
				if l1 != nil {
					l1 = l1.Next
				}
				continue
			}

		}
		if l2 != nil {
			b := l2.Val
			if l1 != nil {
				a := l1.Val
				if a < b {
					prev.Next = &ListNode{Val: a}
					prev = prev.Next
					if l1 != nil {
						l1 = l1.Next
					}
				} else {
					prev.Next = &ListNode{Val: b}
					prev = prev.Next
					if l2 != nil {
						l2 = l2.Next
					}
				}
			} else {
				prev.Next = &ListNode{Val: b}
				prev = prev.Next
				if l2 != nil {
					l2 = l2.Next
				}
				continue
			}
		}

	}
	return result.Next
}
