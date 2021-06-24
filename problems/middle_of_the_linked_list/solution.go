/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
    cur,counter := head,0
    m := make(map[int]*ListNode)
    for cur != nil{
        m[counter] = cur
        cur, counter = cur.Next, counter+1
    }
    return m[(counter/2)]
}