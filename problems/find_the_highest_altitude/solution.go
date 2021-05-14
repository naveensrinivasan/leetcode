func largestAltitude(gain []int) int {
    max,pos := 0,0
    for i:=0;i<len(gain);i++{
        pos = gain[i] + pos
        if pos > max{
            max = pos
        }
    }
    return max
}
