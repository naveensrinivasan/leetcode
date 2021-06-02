func kthSmallest(matrix [][]int, k int) int {
    x:=[]int{}
    for i:=0;i<len(matrix);i++{
        for j:=0;j<len(matrix[i]);j++{
            x = append(x,matrix[i][j])
        }
    }
    sort.Ints(x)
    return x[k-1]
}