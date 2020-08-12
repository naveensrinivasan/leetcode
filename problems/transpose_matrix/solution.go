func transpose(A [][]int) [][]int {
	result := make([][]int, len(A[0]))
	for i, _ := range result {
		result[i] = make([]int, len(A))
	}
	for row := 0; row < len(A); row++ {
		for col := 0; col < len(A[0]); col++ {
			result[col][row] = A[row][col]
		}
	}
	return result
}
