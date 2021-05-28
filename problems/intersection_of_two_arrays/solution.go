func intersection(nums1 []int, nums2 []int) []int {
    m := make(map[int]int)
    result := []int{}
    for _,i := range nums1{
        m[i] = 1
    }
    
    for _, i := range nums2{
        if v,ok:= m[i];ok && v ==1{
            result = append(result,i)
            m[i] = 0
        }
    }
  
    return result
}