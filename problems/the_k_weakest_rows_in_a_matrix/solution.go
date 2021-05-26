func kWeakestRows(mat [][]int, k int) []int {
    type r struct{
        row int
        count int
    }
    
    result := []int{}
    l := []r{}
    
    for i := 0;i<len(mat);i++{
        sum := 0
        for j:=0;j<len(mat[0]);j++{
            sum += mat[i][j]
        }
        l = append(l,r{i,sum})
    }
    
    sort.Slice(l,func(i,j int) bool{
        if l[i].count == l[j].count{
            return l[i].row < l[j].row
        } 
        return l[i].count < l[j].count
    })
    
    for i:= 0;i<k; i++{
        result = append(result,l[i].row)
    }
 
    return result
}