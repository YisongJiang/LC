package graph

import "fmt"

type couples struct {
	accessed bool
	members  [2]int
}

func getAnotherNum(num1 int) int {
	rest := num1 % 2
	if rest == 0 {
		return num1 + 1
	}
	return num1 - 1
}

func minSwapsCouples(row []int) int {
	totalCount := 0
	coupNums := len(row) / 2
	coups := make([]couples, 0, coupNums)
	indexMap := make(map[int]int, len(row))
	sameMap := make(map[int]int, len(row))
	for i := 0; i < coupNums; i++ {
		coups = append(coups, couples{
			false,
			[2]int{
				row[2*i],
				row[2*i+1],
			},
		})
		indexMap[row[2*i]] = i
		indexMap[row[2*i+1]] = i
		sameMap[row[2*i]] = row[2*i+1]
		sameMap[row[2*i+1]] = row[2*i]
	}
	var num1, num2 int
	for i := 0; i < coupNums; i++ {
		if !coups[i].accessed {
			num1 = getAnotherNum(coups[i].members[0])
			for num1 != coups[i].members[1] {
				totalCount++
				num2 = getAnotherNum(coups[i].members[1])
				coups[indexMap[num2]].accessed = true
				coups[i].members[1] = sameMap[num2]
			}
		}
	}
	fmt.Println("totalCount", totalCount)
	return totalCount
}
