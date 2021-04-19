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
	lo, hi,  result := 0, len(s)-1, 0
	for lo <= hi {
		if s[lo] > 0 {
			lo++
		} else if (s[hi] < 0) {
      result++
			hi--
		}else if (s[hi] >= 0){
      hi--
    }
    
	}
	return result
}
