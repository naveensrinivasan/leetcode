func numEquivDominoPairs(dominoes [][]int) int {
    type Key [2]int
    result := 0
    m:= make(map[Key]int)
    for i:=0;i<len(dominoes);i++{
        x := []int{dominoes[i][0],dominoes[i][1]}
        sort.Ints(x)
        f := Key{x[0],x[1]}
        if _,ok := m[f];!ok{
            m[f] = 0
        }else{
           m[f]++
        }
    }
    for _,v := range m{
        result+= (v*(v+1))/2
    }
    return result
}