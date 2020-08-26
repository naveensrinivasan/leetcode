/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	remainder, carry := 0, 0
	var prev, result *ListNode
	for l1 != nil || l2 != nil {
		sum := carry

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		carry, remainder = sum/10, sum%10

		if result == nil {
			result = &ListNode{Val: remainder}
			prev = result
		} else {
			prev.Next = &ListNode{Val: remainder}
			prev = prev.Next
		}
	}
	if carry > 0 {
		prev.Next = &ListNode{Val: carry}
	}
	return result
}
