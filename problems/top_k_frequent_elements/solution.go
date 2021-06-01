func topKFrequent(nums []int, k int) []int {
    type F struct{
        K int
        V int
    }
    m:= make(map[int]int)
    for i:=0;i<len(nums);i++{
        m[nums[i]]++
    }
    l := []F{}
    for k,v := range m{
        l = append(l,F{k,v})
    }
    
    sort.Slice(l,func(i,j int)bool{
        return l[i].V > l[j].V
    })
    
    result :=[]int{}
    for i:=0;i<k;i++{
        result = append(result,l[i].K)
    }
    return result
}