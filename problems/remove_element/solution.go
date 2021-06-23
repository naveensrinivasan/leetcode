func removeElement(nums []int, val int) int {
    counter := 0
    for _, n := range nums{
        if n != val{
            nums[counter] = n
            counter++
        }
    }
    return counter
}