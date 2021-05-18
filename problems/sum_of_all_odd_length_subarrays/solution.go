func sumOddLengthSubarrays(arr []int) int {
    sum := 0
    //sliding window
    //1 , 3 ,5
    for i:= 1;i<=len(arr);i=i+2{
    window := i
        for i:=0;i<=len(arr)-window;i++{
            sum+= sumAll(arr[i:i+window])
        }
    }
    return sum
}
func sumAll(x []int)int{
    sum := 0
    for _, i := range x{
        sum+=i
    }
    return sum
}