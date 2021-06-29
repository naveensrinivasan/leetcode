func countBalls(lowLimit int, highLimit int) int {
    m := make(map[int]int)
    max := 0
    for i:=lowLimit;i<=highLimit;i++{
        if i < 10{
            if v,ok := m[i];!ok{
                m[i]=1
                if max < 1{
                    max = 1
                }
            }else{
                v = v+1
                if max < v {
                    max = v
                }
                m[i] = v
            }
        }else{
            x :=i
            sum := 0
            for x > 0{
               sum += x%10
               x /= 10
            }
            if v,ok := m[sum];!ok{
                m[sum]=1
                if max < 1{
                    max = 1
                }
            }else{
                v = v+1
                if max < v {
                    max = v
                }
                m[sum] = v
            }
        }   
    }
    return max
}