/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil{
        return nil
    }
    
    m := make(map[int]int)
    cur := head
    
    for cur!= nil{
      m[cur.Val]++
      cur = cur.Next
    }
    
    x := []int{}
    for k,v := range m{
        if v == 1{
            x = append(x,k)
        }
    }
    if len(x) == 0{
        return nil
    }
    sort.Ints(x)
    result := &ListNode{Val:x[0]}
    cur = result
    for i:=1;i<len(x);i++{
        l := &ListNode{Val:x[i]}
        cur.Next = l
        cur = cur.Next
    }
    return result
}