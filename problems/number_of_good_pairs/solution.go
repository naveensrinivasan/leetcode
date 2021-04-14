func numIdenticalPairs(nums []int) int {
    m := make(map[int]int)
    counter := 0
    for i:= 0;i<len(nums);i++{
        counter += m[nums[i]]
        m[nums[i]]++
    }
    return counter
}