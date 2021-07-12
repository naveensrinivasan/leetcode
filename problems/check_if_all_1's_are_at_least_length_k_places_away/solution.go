func kLengthApart(nums []int, k int) bool {
    counter := 0
    prev := 0
    isFirst :=true
    for i,item := range nums{
        if item ==1 && !isFirst{
            prev = counter
            counter = i 
            if prev != counter && (counter - prev)-1 <k{
                return false
            }
        }
        if item == 1 && isFirst{
            isFirst = false
        }
    }
    return true
}