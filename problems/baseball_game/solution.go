func calPoints(ops []string) int {
    res :=[]int{}
    sum := 0
    
    for _,i := range ops{
        
        if i == "C"{
            res = res[:len(res)-1]  
        }else if i == "D"{
            
            res = append(res,res[len(res)-1]*2)        
            
        }else if i == "+"{
            res = append(res,res[len(res)-1]+res[len(res)-2])        
        }else{
            num , _ := strconv.Atoi(i)
            res = append(res,num)   
        }
    }
    for _, i:= range res{
        sum+=i
    }
    return sum
}