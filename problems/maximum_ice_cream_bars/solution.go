func maxIceCream(costs []int, coins int) int {
    sort.Ints(costs)
    result:= 0
    for _, i:= range costs{
        if i > coins{
            break
        }
        coins -=  i
        result++
    }
    return result 
}