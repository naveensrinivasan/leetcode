/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNodes(head *ListNode, m int, n int) *ListNode {
    cur := head
    var result *ListNode
    var innerresult *ListNode
    isSkip,skipCounter := false,0
    for cur != nil{
        if !isSkip{
            if result == nil{
                result = &ListNode{}
                innerresult = result
                innerresult.Val = cur.Val
            }else{
                l := &ListNode{Val: cur.Val}
                innerresult.Next = l
                innerresult = innerresult.Next
            }
            skipCounter++
            if skipCounter == m{
                isSkip = true
                skipCounter = 0
            }
            cur = cur.Next
        }else{
            skipCounter++
            if skipCounter == n{
                isSkip = false
                skipCounter = 0
            }
            cur = cur.Next
        }
    }
    
    return result
}