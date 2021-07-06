func maxNumberOfApples(arr []int) int {
    counter := 0
    sum := 0
    sort.Ints(arr)
    for _,x := range arr{
        sum += x
        if sum > 5000{
            return counter 
        }
        counter++
    }
    return counter
}