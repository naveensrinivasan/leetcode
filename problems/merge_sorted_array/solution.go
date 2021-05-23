func merge(nums1 []int, m int, nums2 []int, n int)  {
    if len(nums2) ==0 {
        return 
    }
    for i:= 0;i<len(nums2);i++{
        nums1[len(nums1)-n+i] = nums2[i]
        
    }
    sort.Ints(nums1)
}