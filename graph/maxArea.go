package graph

func maxArea(height []int) int {
	// iMap := make(map[int]int, len)
	// for i, val := range height  {
	// 	iMap[]
	// }
	// sort.Ints(height)
	maxVal := 0
	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			val := calculateArea(height, i, j)
			if val > maxVal {
				maxVal = val
			}
		}
	}
	return maxVal
}

func calculateArea(height []int, i, j int) int {
	h := height[i]
	if height[i] > height[j] {
		h = height[j]
	}
	return (j - i) * h
}
