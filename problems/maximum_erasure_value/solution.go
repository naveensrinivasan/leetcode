func maximumUniqueSubarray(nums []int) int {
    l,r,max,mapsize :=0,0,0,0
    m := make(map[int]int)
    for (r<len(nums)){
        if _, ok := m[nums[r]];!ok{
            m[nums[r]] = nums[r]
            mapsize += nums[r]
            r++
            max = Max(max,mapsize)
        }else{
            delete(m,nums[l])
            mapsize -= nums[l]
            l++
        }
    }
    return max
}

func calcSum(m map[int]int)int{
    sum := 0
    for i := range m{
        sum+=i
    }
    return sum
}

func lengthOfLongestSubstring(s string) int {
   a,b,max,mapsize := 0,0,0,0
    m := make(map[byte]byte)
    for (b<len(s)){
        if _,ok := m[s[b]];!ok{
            m[s[b]]=s[b]
            b++
            mapsize++
            max = Max(max,mapsize)
        }else{
            delete(m,s[a])
            a++
            mapsize--
        }
    }
    
    return max
}

// Max returns the larger of x or y.
func Max(x, y int) int {
    if x < y {
        return y
    }
    return x
}