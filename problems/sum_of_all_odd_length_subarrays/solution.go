func sumOddLengthSubarrays(arr []int) int {
    sum := 0
    total := arr[0]
    result := make([]int,len(arr))
    result[0] = arr[0]
    for i:=1;i<len(arr);i++{
        total += arr[i]
        result[i] = total
    }
    //sliding window
    //1 , 3 ,5
    for i:= 1;i<=len(arr);i=i+2{
    window := i
        for i:=0;i<=len(arr)-window;i++{
            if window == 1{
                sum += arr[i]
            }else{
                if i == 0{
                    sum+= result[i+window-1]
                }else{
                    
                    sum+= result[i+window-1] - result[i-1]
                }
            }
        }
    }
    return sum
}
