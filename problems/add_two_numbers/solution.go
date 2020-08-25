/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	var cur, prev, result *ListNode
	for l1 != nil || l2 != nil {
		var sum = 0

		if l1 != nil {
			sum += l1.Val
		}
		if l2 != nil {
			sum += l2.Val
		}

		sum += carry
		carry = sum / 10
		if result == nil {
			result = &ListNode{Val: sum % 10}
			prev = result
		} else {
			cur = &ListNode{Val: sum % 10}
			prev.Next, prev = cur, cur
		}

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	if carry != 0 {
		cur = &ListNode{Val: carry}
		prev.Next, prev = cur, cur
	}
	return result
}
