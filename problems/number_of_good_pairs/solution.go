func numIdenticalPairs(nums []int) int {
    m := make(map[int]int)
    counter := 0
    for _,j := range nums{
      counter += m[j]
      m[j]++
    }
    return counter
}