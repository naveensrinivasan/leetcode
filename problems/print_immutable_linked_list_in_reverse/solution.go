/*   Below is the interface for ImmutableListNode, which is already defined for you.
 *
 *   type ImmutableListNode struct {
 *       
 *   }
 *
 *   func (this *ImmutableListNode) getNext() ImmutableListNode {
 *		// return the next node.
 *   }
 *
 *   func (this *ImmutableListNode) printValue() {
 *		// print the value of this node.
 *   }
 */

func printLinkedListInReverse(head ImmutableListNode) {
    x := []ImmutableListNode{}
    next := head
    x = append(x,next)
    for next != nil{
        next = next.getNext()
        if next == nil{
            break
        }
        x = append(x,next)
    }
    
    for i := len(x)-1;i>=0;i--{
        x[i].printValue()
    }
    
}