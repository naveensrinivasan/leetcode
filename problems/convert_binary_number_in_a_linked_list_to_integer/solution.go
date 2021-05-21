/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getDecimalValue(head *ListNode) int {
    result := 0
    
    for head != nil{
        result,head = (result *2 ) + head.Val , head.Next
    }
    
    return result
}