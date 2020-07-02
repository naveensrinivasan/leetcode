func twoSum(nums []int, target int) []int {
    m := make(map[int]int)
    
    for i,v := range nums{
        potentialMatch := target - v
        if location, ok := m[potentialMatch];ok{
            return []int{i,location}
        }
        m[v] = i
    }
    return nil
}