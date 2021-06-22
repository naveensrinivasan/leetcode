func mySqrt(x int) int {
    if x ==0 || x==1{
        return x
    }
    left,right,ans:= 1,x,0
    for left <= right{
        mid := (left+right)/2
        y := mid * mid
        
        if y == x{
            return mid
        }
        
        if y < x{
            left = mid +1
            ans = mid
        }else{
            right = mid -1
        }
    }
    return ans
}