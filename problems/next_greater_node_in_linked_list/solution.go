/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func nextLargerNodes(head *ListNode) []int {
    type F struct{
        K,P,V int
    }
    ms:= []*F{}
    l := []*F{}
    cur := head
    for counter:=0;cur != nil;cur,counter=cur.Next,counter+1{
        for len(ms) > 0 && cur.Val > ms[len(ms)-1].K{
            ms[len(ms)-1].V = cur.Val
            ms = ms[:len(ms)-1]
        }
        v := &F{cur.Val,counter,0}
        ms = append(ms,v)
        l = append(l,v)
    }
    sort.Slice(l,func(i,j int)bool{
        return l[i].P < l[j].P
    })
    r := make([]int,len(l))
    for i,f := range l{
        r[i]=f.V
    }
    return r
}