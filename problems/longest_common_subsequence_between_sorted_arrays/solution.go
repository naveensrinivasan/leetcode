func longestCommomSubsequence(arrays [][]int) []int {
    m := make(map[int]int)
    for i,item := range arrays{
        for _,x := range item{
            if i == 0{
                m[x]++
            }else if _, ok := m[x];ok{
                m[x]++
            }
        } 
    }
    f := []int{}
    for k,v := range m{
        if v == len(arrays){
            f = append(f,k)
        }
    }
    sort.Ints(f)
    return f
}