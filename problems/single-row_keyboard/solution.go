func calculateTime(keyboard string, word string) int {
    m,prev,result := make(map[rune]int),0,0
    for i,j:= range keyboard{
        m[j]=i
    }
    for _,w := range word{
        move := m[w]
        result += abs(prev-move)
        prev= move
        
    }
    return result
}
func abs(i int)int{
    if i<0 {
        return i*-1
    }
    return i
}