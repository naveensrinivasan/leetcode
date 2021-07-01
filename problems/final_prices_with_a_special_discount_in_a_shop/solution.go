func finalPrices(prices []int) []int {
    result := []int{}
    s := []int{}
    for i:=len(prices)-1;i>=0;i--{
        for len(s) > 0 && prices[i] < s[len(s)-1]{
            s = s[:len(s)-1]
        }
        v := prices[i]
        if len(s) > 0{
            v = prices[i] - s[len(s)-1] 
        }
        result = append([]int{v},result...)
        s = append(s,prices[i])
    }
    return result
}