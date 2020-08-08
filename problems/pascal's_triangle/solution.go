func generate(numRows int) [][]int {
	if numRows == 0 {
		return [][]int{}
	}
	if numRows == 1 {
		return [][]int{[]int{1}}
	}
	result := [][]int{[]int{1}, []int{1, 1}}

	if numRows == 2 {
		return result
	}

	for row := 2; row < numRows; row++ {
		colrow := make([]int, row+1)
		for col := 0; col < len(colrow); col++ {
			if col == 0 || col == len(colrow)-1 {
				colrow[col] = 1
				continue
			}
			colrow[col] = (result[row-1][col-1] + result[row-1][col])
		}
		result = append(result, colrow)
	}

	return result

}
