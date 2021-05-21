func kidsWithCandies(candies []int, extraCandies int) []bool {
    x := make([]int,len(candies))
    copy(x,candies)
    sort.Ints(x)
    greatest := x[len(x)-1]
    result := []bool{}
    for _ , i := range candies{
        r := false
        if i + extraCandies >= greatest{
            r = true
        }
        result = append(result,r)
    }
    return result
}