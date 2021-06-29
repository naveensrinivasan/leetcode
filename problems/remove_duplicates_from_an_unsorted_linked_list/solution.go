/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicatesUnsorted(head *ListNode) *ListNode {
    type F struct{
        K,C,P int
    }
    cur,m,counter:= head, make(map[int]F),0
    
    for cur != nil{
        if v,ok := m[cur.Val];!ok{
            m[cur.Val] = F{cur.Val,1,counter}
        }else{
            v.C++
            m[cur.Val] = v
        } 
        counter++
        cur = cur.Next
    }
    l := []F{}
    for _,V := range m{
        if V.C == 1{
            l = append(l,V)
        } 
    }
    sort.Slice(l,func(i,j int)bool{
       return l[i].P < l[j].P
    })
    if len(l) == 0{
        return nil
    }
    head = &ListNode{Val:l[0].K}
    cur = head
    for i:=1;i<len(l);i++{
        x := &ListNode{Val:l[i].K}
        cur , cur.Next = x,x
    }
    return head
}