

func frequencySort(nums []int) []int {
    m := make(map[int]int)
    type F struct{
        K int
        V int
    }
    for i:= 0;i<len(nums);i++{
        m[nums[i]]++
    }
    l := []F{}
    for k,v := range m{
        l = append(l,F{k,v})
    }
    
    sort.Slice(l,func(i,j int) bool{
        if l[i].V == l[j].V{
            return l[i].K > l[j].K
        }
        return l[i].V < l[j].V
    })
    result := []int{}
    for i:=0;i<len(l);i++{
        for j:=0;j<l[i].V;j++{
            result = append(result,l[i].K)
        }
    }
    return result
}