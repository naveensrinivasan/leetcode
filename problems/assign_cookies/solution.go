func findContentChildren(g []int, s []int) int {
    sort.Ints(g)
    sort.Ints(s)
    counter := 0
    for i:=0;i<len(g);i++{
        x := g[i]
        if len(s) == 0{
            break
        }
        for i:=0;i<len(s);i++{
            y := s[i]
            if y>=x{
                s = s[i+1:]
                counter++
                break
            }
        }
    }
    return counter
}