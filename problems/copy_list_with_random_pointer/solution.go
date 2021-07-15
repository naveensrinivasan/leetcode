/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
    if head == nil{
        return nil
    }
    m := make(map[*Node]*Node)
    cur := head
    result := &Node{}
    resultCur := result
    
    for cur!= nil{
        if v,ok := m[cur];!ok{
            // new node isn't there in the map
            if cur!= head{
                x := &Node{}
                resultCur.Next = x
                resultCur = resultCur.Next
            }
            m[cur] = resultCur
        }else{
            resultCur.Next = v
            resultCur = resultCur.Next
        }
        resultCur.Val = cur.Val
        
        if cur.Random != nil {
            if v,ok := m[cur.Random];!ok{
                x := &Node{Val:cur.Random.Val}
                resultCur.Random =x
                m[cur.Random] = x
            }else{
                resultCur.Random = v
            }
        }
        cur = cur.Next
    }
    return result
}