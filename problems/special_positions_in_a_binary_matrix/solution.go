func numSpecial(mat [][]int) int {
    result := 0
    
    for i:=0;i<len(mat);i++{
        sum:= 0 
        colPos := -1
        for j:=0;j<len(mat[0]);j++{
            sum+= mat[i][j] 
            if mat[i][j] == 1{
                colPos = j
            }
        }
        if sum == 1{
            colSum := 0 
            for j:=0;j<len(mat);j++{
                colSum+= mat[j][colPos] 
            }
            if colSum == 1{
                result++
            }
        }   
    }
    return result
}