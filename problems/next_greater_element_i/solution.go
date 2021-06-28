func nextGreaterElement(nums1 []int, nums2 []int) []int {
    m := make(map[int]int,len(nums2))
    s:= []int{}
    result := make([]int,len(nums1))
    
    for i:=len(nums2)-1;i >=0;i--{
        for len(s)>0 && nums2[i] > s[len(s)-1]{
            s = s[:len(s)-1]
        }
        v := -1
        if len(s) > 0{
            v = s[len(s)-1]
        }
        m[nums2[i]] = v
        s = append(s,nums2[i])
    }
    
    for i, item := range nums1{
        result[i] = m[item]
    }
    
    return result
}