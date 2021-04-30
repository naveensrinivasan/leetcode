func countNegatives(grid [][]int) int {
    result :=0
    for _,n := range grid{
        
        x := getnegativenumbers(n)
        fmt.Println(x)
        result+=x
    }
    return result
}

func getnegativenumbers(s []int) int {
	 result := 0
    for i:= len(s)-1;i>=0;i--{
        if s[i] <0{
            result++
        }else{
            break
        }
    }
	
	return result
}
