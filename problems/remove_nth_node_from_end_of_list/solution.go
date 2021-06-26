/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    cur,cur2,counter:= head,head,0
    if head == nil{
        return head
    }
    
    for cur!= nil{
        counter++
        cur = cur.Next
    }
    if counter == 1{
        return nil
    }
    
    pos := counter -n
    prev := cur2
    
    for cur2!= nil{
        if pos == 0{
            if prev == cur2{
                return prev.Next
            }
            prev.Next = cur2.Next
            break
        } 
        pos--
        prev = cur2
        cur2 = cur2.Next
    }
    return head
}