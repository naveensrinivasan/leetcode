func maxScore(s string) int {
    max := 0
    for i:=0;i<len(s);i++{
        l,r := s[:i],s[i:]
        x, z := findZerosAndOnes(l)
        a,y := findZerosAndOnes(r)
        if x == 0 && z == 0{
            continue
        }else if a == 0 && y == 0{
            continue
        }
        sum := x+y
        if sum > max{
            max = sum
        }
    }
    return max
}
func findZerosAndOnes(s string)(int,int){
    zeros,ones := 0,0
    for _, item := range strings.Split(s,""){
        if item == "0"{
            zeros++
        }else{
            ones++
        }
    }
    return zeros,ones
}