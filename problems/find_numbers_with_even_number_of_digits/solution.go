func findNumbers(nums []int) int {
    count := 0
    for _, num := range nums{
        counter := 0
        for num > 0{
            counter++
            num/=10
        }
        if counter%2==0{
            count++
        }
    }
    return count
}