/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
    var list2Tail *ListNode
    var prev *ListNode
    counter:=0
    cur := list2
    list2Tail = cur
    
    for cur != nil{
        list2Tail, cur = cur,cur.Next
    }
    cur = list1
    
    for cur != nil{
        if counter == a{
           prev.Next = list2
        }
        if counter == b{
            list2Tail.Next= cur.Next
            break
        }
        prev, cur = cur,cur.Next
        counter++
    }
    return list1
}