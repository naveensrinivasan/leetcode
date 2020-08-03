

func twoSum(nums []int, target int) []int {
    m := make(map[int]int)
    for i , value := range nums{
        n := target - value
        if pos,ok := m[n];ok{
            return[]int {i,pos}
        }
        m[value] = i
       
    }
    
    return nil
}