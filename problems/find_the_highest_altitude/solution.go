func largestAltitude(gain []int) int {
    max,pos := 0,0
    for i:=0;i<len(gain);i++{
        if gain[i] > 0{
            pos = gain[i] + pos
        }else{
            pos = gain[i] + pos
        }
        if pos > max{
            max = pos
        }
    }
    return max
}
