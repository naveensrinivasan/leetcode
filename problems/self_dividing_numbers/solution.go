func selfDividingNumbers(left int, right int) []int {
    result := []int{}
    
    for i:= left;i<=right;i++{
        if i%10 == 0{
                continue
            }
        wasDivisble := false
        //fmt.Println(getdigits(i))
        for _ , d := range  getdigits(i){
            
            if i%d != 0{
                wasDivisble = false
                break
            } 
            wasDivisble = true
        }
        if wasDivisble{
            result = append(result, i)
        }
    }
    return result
}

func getdigits(n int)[]int{
    if n< 10{
        return []int{n}
    }
    result := []int{}
    for n > 0{
        if n== 0 || n % 10 == 0{
            return  []int{}
        } 
        result =append(result, n %10)    
        n/=10
    }
    return result
}