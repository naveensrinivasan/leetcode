func findBuildings(heights []int) []int {
    type K struct{
        K,V int
    }
    s := []K{}
    r := []int{}
    for i:=0;i<len(heights);i++{
        for len(s) > 0 && heights[i] >= s[len(s)-1].K{
            s = s[:len(s)-1]
        }
        s = append(s,K{heights[i],i})
    }
    for _,v := range s{
        r = append(r,v.V)
    }
    return r
}