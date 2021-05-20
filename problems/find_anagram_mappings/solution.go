func anagramMappings(nums1 []int, nums2 []int) []int {
    s := make(map[int]int)
    result := []int{}
    for i,j := range nums2{
        s[j]=i
    }
    for _, i := range nums1{
        result = append(result,s[i])
    }
    return result
}