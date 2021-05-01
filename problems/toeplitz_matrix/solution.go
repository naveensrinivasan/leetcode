func isToeplitzMatrix(matrix [][]int) bool {
    cols := len(matrix[0])
    // [0,0],[1,1]
    // [0,1,[1,2]
    // [0,2][1,3]
    
    // [1,0],[2,1] 
    // [1,1],[2,2]
    // [1,2],[2,3]
    
    for r:=0;r<len(matrix)-1;r++{
        for c:=0;c<cols-1;c++{
            fmt.Println(r,c)
            if  matrix[r][c] != matrix[r+1][c+1]{
                return false
            }
        }
    }
    return true
}