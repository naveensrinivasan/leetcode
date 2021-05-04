func diagonalSum(mat [][]int) int {
    sum := 0
    for i,col :=0,len(mat) -1;i<len(mat);i,col = i+1,col-1{
        sum += mat[i][i] 
        if i != col{
            sum += mat[i][col]
        }
    }
    return sum
}