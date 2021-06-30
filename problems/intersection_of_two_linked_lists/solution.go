/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    m := make(map[*ListNode]int)
    cur,cur2 := headA,headB
    for cur!= nil || cur2 != nil{
        if cur!= nil {
            if _,ok := m[cur];!ok{
                m[cur] = 1
            }else{
                return cur
            }
            cur = cur.Next
        }
        if cur2!= nil {
            if _,ok := m[cur2];!ok{
                m[cur2] = 1
            }else{
                return cur2
            }
            cur2 = cur2.Next
        }
    }
    return nil
}