/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	s1 := listToSlice(l1)
	s2 := listToSlice(l2)

	s1len, s2len := len(s1), len(s2)
	if s2len > s1len {
		s1, s2 = s2, s1
		s2len, s1len = s1len, s2len
	}

	var cur *ListNode

	carry := 0
	for i, j := s1len-1, s2len-1; i >= 0; i, j = i-1, j-1 {
		a, b := 0, 0
		if j >= 0 {
			b = s2[j]
		}
		a = s1[i]
		carry, cur = addtwoListNodes(a, b, carry, cur)
	}
	if carry > 0 {
		_, cur = addtwoListNodes(0, 0, carry, cur)
	}
	return cur

}
func addtwoListNodes(a, b, carry int, prev *ListNode) (int, *ListNode) {
	sum := a + b + carry
	carry = sum / 10
	cur := &ListNode{Val: sum % 10, Next: prev}
	return carry, cur
}
func listToSlice(l *ListNode) []int {
	s := []int{}
	for iter := l; iter != nil; iter = iter.Next {
		s = append(s, iter.Val)
	}
	return s
}
