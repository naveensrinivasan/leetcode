func numJewelsInStones(jewels string, stones string) int {
    j := make(map[rune]rune)
    counter := 0
    for _,i := range jewels{
        j[i] = i
    } 
    for _,item := range stones{
        if _,ok := j[item];ok{
            counter++
        }
    } 
    return counter
}