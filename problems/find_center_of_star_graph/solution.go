func findCenter(edges [][]int) int {
	m := make(map[int]int)
	maxCount := 0
	max, key := -1, 0
    
	for _, node := range edges {
		m[node[1]]++
	}
    
	for k, v := range m {
		if v == max {
			maxCount++
		}
		if v > max {
			max, key = v, k
		}
	}
    
	if maxCount > 0 {
		m1 := make(map[int]int)
		for _, node := range edges {
			m1[node[0]]++
		}
		max = -1
		for k, v := range m1 {
			if v > max {

				max, key = v, k
			}

		}

	}
	return key
}