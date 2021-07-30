func merge(arr [][]int) [][]int {
    if len(arr) == 1{
        return arr
    }
    sort.Slice(arr,func(i,j int)bool{
        return arr[i][0] < arr[j][0]
    })
    m := make(map[[2]int]bool)
    result := [][]int{}
    for i:=0;i<len(arr)-1;i++{
        if arr[i+1][0] <= arr[i][1]{
            y := arr[i+1][1]
            if arr[i][1] > y{
                y = arr[i][1]
            }
            result = append(result,[]int{arr[i][0],y})
            m[[2]int{arr[i][0],y}] = true
            i++
        }else{
            result = append(result,arr[i])
            m[[2]int{arr[i][0],arr[i+1][1]}] = true
        } 
    }
        
    if arr[len(arr)-1][0] <= arr[len(arr)-2][1]{
            y := arr[len(arr)-2][1]
            if arr[len(arr)-1][1] > y{
                y = arr[len(arr)-1][1]
            }
        if _, ok := m[[2]int{arr[len(arr)-2][0],y}];!ok{
            result = append(result,[]int{arr[len(arr)-2][0],y})
        } 
    }else{
        result = append(result,arr[len(arr)-1])
    } 
    
    if reflect.DeepEqual(result,arr){
        return result
    }else{
        result = merge(result)  
    }
    return result
}