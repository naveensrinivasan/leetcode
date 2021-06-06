func isSumEqual(firstWord string, secondWord string, targetWord string) bool {
    x := ""
    for _ , i := range firstWord{
        x += strconv.Itoa(int(i-97))
    }
    y , _ := strconv.Atoi(x)
    
    x = ""
    for _ , i := range secondWord{
        x += strconv.Itoa(int(i-97))
    }
    z , _ := strconv.Atoi(x)
    
    x = ""
    for _ , i := range targetWord{
        x += strconv.Itoa(int(i-97))
    }
    a , _ := strconv.Atoi(x)
    
    return a == y+z
}