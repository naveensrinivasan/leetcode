func reverse(x int) int {
    y := int(math.Abs(float64(x)))
    result := 0
    for y > 0 {
        result = result * 10
        result += y %10
        y /= 10
    
       
    }
    if result < -2147483648 {
		return 0
	}else if result > 2147483647 {
		return 0
	} else if x < 0{         
     return -1*result   
    }
 
    return result
    
}