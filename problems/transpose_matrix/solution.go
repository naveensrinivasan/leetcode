func transpose(A [][]int) [][]int {
	var result [][]int
	for col := 0; col < len(A[0]); col++ {
		var interim []int
		for i := 0; i < len(A); i++ {
			interim = append(interim, A[i][col])
		}
		result = append(result, interim)
	}
	return result

}
