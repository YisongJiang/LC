package firstgroup

// 75th sort colors
func sortColors(nums []int) {
	frontIndex := 0
	currentIndex := 0
	endIndex := len(nums) - 1
	for endIndex > currentIndex {
		if nums[currentIndex] == 0 {
			nums[frontIndex], nums[currentIndex] = nums[currentIndex], nums[frontIndex]
			frontIndex++
		} else if nums[currentIndex] == 1 {
			currentIndex++
		} else if nums[currentIndex] == 2 {
			nums[endIndex], nums[currentIndex] = nums[currentIndex], nums[endIndex]
			endIndex--
		}
		if frontIndex >= currentIndex {
			currentIndex = frontIndex + 1
		}
	}
}
