func decrypt(code []int, k int) []int {
    if k == 0{
        return make([]int,len(code))
    }
    max := len(code)
    result := []int{}
    if k > 0{
        for i:=0;i<len(code);i++{
            sum := 0
            for j:=0;j<k;j++{
                sum +=  code[clockArthmetic(max,i+1+j)]
            } 
            result = append(result,sum)
        }        
    }else{
        for i:=0;i<len(code);i++{
            sum := 0
            for j:= (k * (-1));j>0;j--{
                r := i - j
                if r < 0{
                    sum += code[len(code)+r]
                }else{
                    sum += code[r]
                }
            }
            result = append(result,sum)
        }
    }
    return result
}

func clockArthmetic(max,sum int)int{
    if sum >= max {
        return sum - max
    }
    return sum
}