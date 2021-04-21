func validMountainArray(arr []int) bool {
    isDoneIncreasing ,dec,inc := false,false,false
    if len(arr) < 3{
        return false
    }
    for i:= 1;i<len(arr);i++{
        if arr[i-1] < arr[i]{ 
            if isDoneIncreasing{
                return false
            }
            inc = true
        }else if arr[i-1] == arr[i]{
            return false
        }else if arr[i-1] > arr[i]{
            isDoneIncreasing = true
            dec = true
        } 
        
    }
    return isDoneIncreasing && dec && inc
}