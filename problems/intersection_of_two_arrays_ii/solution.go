func intersect(nums1 []int, nums2 []int) []int {
      m := make(map[int]int)
    result := []int{}
    for _,i := range nums1{
        m[i]++
    }
    
    for _, i := range nums2{
        if v,ok:= m[i];ok && v>0{
            result = append(result,i)
            m[i]--
        }
    }
  
    return result
}