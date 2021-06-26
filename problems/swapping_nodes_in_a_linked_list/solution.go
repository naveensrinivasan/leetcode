/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapNodes(head *ListNode, k int) *ListNode {
    cur,counter:= head,0
    var startNode *ListNode
    for cur != nil && cur.Next != nil{
        counter+=2
        cur = cur.Next.Next
    }
    if cur != nil{
        counter++
    }
    
    cur = head
    x := 0
    
    for cur != nil{
        x++
        if x == k{
            startNode = cur
            break
        }
        cur = cur.Next
    }
    cur = head
    for cur != nil{
        if counter == k{
            startNode.Val , cur.Val = cur.Val,startNode.Val
            return head
        }
        counter--
        cur = cur.Next
    }
    
    return head
}