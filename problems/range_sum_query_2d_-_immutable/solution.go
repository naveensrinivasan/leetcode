type NumMatrix struct {
    L[][]int
}


func Constructor(matrix [][]int) NumMatrix {
    return NumMatrix{matrix}
}


func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
    sum := 0
    for i:=row1;i<=row2;i++{
        for j:=col1;j<=col2;j++{
            sum+= this.L[i][j]
        }
    }
    return sum
}


/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */