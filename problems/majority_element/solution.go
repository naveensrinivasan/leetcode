func majorityElement(nums []int) int {
    m := make(map[int]int)
    for i:=0;i<len(nums);i++{
        m[nums[i]]++
    }
    max,key := -1000000 , 0
    fmt.Println(m)
    for k,v := range m{
        
        if v> max {
            max ,key= v,k
        }
        fmt.Println(v,max)
    }
    
    return key
}