func singleNumber(nums []int) int {
    sum := 0
    for _, i := range nums{
        sum = sum ^ i
    }
    return sum
}