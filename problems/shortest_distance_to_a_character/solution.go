func shortestToChar(s string, c byte) []int {
    occ := []int{}
    result := []int{}
    for i,j := range s{
        if byte(j) == c {
            occ = append(occ,i)
        }
    }
    for i:= 0;i<len(s);i++{
        min := 10000000
        for _,item := range occ{
            r := abs(i-item)
            if r < min{
                min = r
                
            }
           
        }
        result = append(result,min)
    }
    return result
}
func abs(i int) int{
    if i < 0 {
        return i * -1
    }
    return i
}