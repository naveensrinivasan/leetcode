func distributeCandies(candyType []int) int {
    m := make(map[int]int)
    for _,i := range candyType{
        m[i]++
    }
    types := len(m)
    max := len(candyType)/2
    if types < max{
        return types
    }
    return max
}