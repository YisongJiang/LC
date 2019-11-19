package graph

func threeSum(nums []int) [][]int {
	res := make([][]int, 0, 10)
	if len(nums) < 3 {
		return res
	}
	indexMap := make(map[int][]int, len(nums))
	deDuplicateMap := make(map[int][]int, len(nums))
	for i, val := range nums {
		if _, ok := indexMap[val]; !ok {
			indexMap[val] = []int{}
		}
		indexMap[val] = append(indexMap[val], i)
	}
	for i := 0; i < len(nums); i++ {
		firstNum := nums[i]
		for j := i + 1; j < len(nums); j++ {
			secondNum := nums[j]
			targetNum := 0 - firstNum - secondNum
			exist := isTargetNumExist(nums, indexMap, targetNum, i, j)
			if exist {
				duplicated := resultIsDuplicated(nums, deDuplicateMap, targetNum, i)
				if !duplicated {
					res = append(res, []int{firstNum, secondNum, targetNum})
					updateDupliMap(deDuplicateMap, firstNum, secondNum, targetNum)
				}
			}

		}
	}
	return res
}

func isTargetNumExist(nums []int, indexMap map[int][]int, targetNum, firstIndex, secondIndex int) bool {
	if indexs, ok := indexMap[targetNum]; ok {
		for _, val := range indexs {
			if val != firstIndex && val != secondIndex {
				return true
			}
		}
	}
	return false
}

func resultIsDuplicated(nums []int, deDuplicateMap map[int][]int, targetNum, firstIndex int) bool {
	firstNum := nums[firstIndex]
	if elements, ok := deDuplicateMap[targetNum]; ok {
		for _, e := range elements {
			if e == firstNum {
				return true
			}
		}
	}
	return false
}

func updateDupliMap(deDuplicateMap map[int][]int, firstNum, secondNum, targetNum int) {
	deDuplicateMap[firstNum] = append(deDuplicateMap[firstNum], secondNum, targetNum)
	deDuplicateMap[secondNum] = append(deDuplicateMap[secondNum], firstNum, targetNum)
	deDuplicateMap[targetNum] = append(deDuplicateMap[targetNum], firstNum, secondNum)
}
