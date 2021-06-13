func numberOfSteps(num int) int {
    if num <1{
        return 0
    }
    counter := 0
    for num > 0{
        if num %2 == 0{
            counter++
        }else{
            counter+=2
        }
        num >>= 1
    }
    return counter -1
}