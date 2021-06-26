/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
    var prev *ListNode
    cur := head
    counter := 0
    for cur != nil{
        if prev != nil && counter%2==1{
            prev.Val , cur.Val = cur.Val, prev.Val
            fmt.Println(prev,cur)
        }
        prev = cur
        cur = cur.Next
        counter++
    }
    return head
}