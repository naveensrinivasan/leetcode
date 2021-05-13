func hammingDistance(x int, y int) int {
    result := 0
    for x > 0 || y > 0{
  
 if x&1 !=y &1{
   result++
 }

 x >>=1
 y >>=1
}
    return result
}