package firstgroup

import (
	"fmt"
	"testing"
)

func TestSortColors(t *testing.T) {
	fmt.Println("start")
	tests := [][]int{
		[]int{},
		[]int{0, 0, 0},
		[]int{1, 1, 1},
		[]int{2, 0, 0, 1, 0, 2, 1, 0},
	}

	okResults := [][]int{
		[]int{},
		[]int{0, 0, 0},
		[]int{1, 1, 1},
		[]int{0, 0, 0, 0, 1, 1, 2, 2},
	}

	for i, result := range tests {
		sortColors(result)
		okResult := okResults[i]
		if len(result) == len(okResult) {
			for in := range result {
				if result[in] != okResult[in] {
					t.Errorf("排序结果不同， index: %d; expected %v; got: %v", i, okResult, result)
					break
				}
			}
		} else {
			t.Errorf("结果和期望长度不同，index: %d; expected %v; got: %v", i, okResult, result)
		}
	}
}
