func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
    type X struct{
        K,V,S int
    }
    l := []X{}
    for i:=0;i<len(nums1);i++{
        for j:=0;j<len(nums2);j++{
            l = append(l,X{nums1[i],nums2[j],nums1[i]+nums2[j]})
        }
    }
    sort.Slice(l,func(i,j int) bool{
        return l[i].S < l[j].S
    })
    result := [][]int{}
    max := k
    if len(l) < k{
        max = len(l)
    }
    for i:=0;i<max;i++{
        result = append(result,[]int{l[i].K,l[i].V})
    } 
    return result
}