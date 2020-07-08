func numberOfSteps (num int) int {
    result := num
    counter:= 0
    for result > 0{
        counter ++
        if result % 2 == 0{
            result = result /2
            continue
        }
        result--
    }
    return counter
}