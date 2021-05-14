func heightChecker(heights []int) int {
    temp := make([]int,len(heights))
    copy(temp,heights)
    sort.Ints(temp)
    counter := 0
    fmt.Println(temp,heights)
    for i := range heights{
        if heights[i] != temp[i]{
            counter++
        }
    }
    return counter
}