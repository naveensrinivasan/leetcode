/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	if slow.Next == nil {
		return true
	}

	second := reverseLList(slow.Next)
	cur := head
	for cur != nil && second != nil {
		if cur.Val != second.Val {
			return false
		}
		cur = cur.Next
		second = second.Next
	}
	return true
}

func reverseLList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var prev *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next, prev, cur = prev, cur, next
	}

	head.Next = nil
	return prev
}
