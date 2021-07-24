/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func closestKValues(root *TreeNode, target float64, k int) []int {
    type F struct{
        K int
        V float64
    }
    s := []int{}
    s= helper(root,s)
    x := []F{}
    for _, item := range s{
        x = append(x,F{item,Abs(float64(item),target)})
    }
    sort.Slice(x,func(i,j int)bool{
        return x[i].V < x[j].V
    })
    g := make([]int,k)

    for i:=0;i<k;i++{
        g[i] = x[i].K    
    } 
    return g
}

func helper(root *TreeNode,x []int)[]int {
    
    if root == nil{
        return x
    }
    x = append(x,root.Val)
    x = helper(root.Left,x) 
    x = helper(root.Right,x) 
    return x
}

func Abs(i,j float64)float64{
    x := i-j
    if x < 0 {
        return x * -1
    }
    return x
}
