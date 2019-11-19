package graph

func IsEleInSet(s map[int]int, e int) bool {
	return s[e] > 0
}

func AddEleToSet(s map[int]int, e, setIndex int) {
	s[e] = setIndex
}

func JoinTwoSet(s map[int]int, setIndex1, sIndex2 int) {
	if setIndex1 == sIndex2 {
		return
	}
	if setIndex1 > sIndex2 {
		setIndex1, sIndex2 = sIndex2, setIndex1
	}
	for key := range s {
		if s[key] == sIndex2 {
			s[key] = setIndex1
		}
	}
}

func findRedundantDirectedConnection(edges [][]int) []int {
	s := make(map[int]int, 1000)
	p := make(map[int]int, 1000)
	result := make([]int, 2)
	posibles := make([][]int, 0, 2)
	for _, e := range edges {
		child := e[1]
		parent := e[0]
		if p[child] == 0 {
			p[child] = parent
		} else {
			posibles = append(posibles, []int{p[child], child})
			posibles = append(posibles, e)
			break
		}
		if !IsEleInSet(s, parent) {
			AddEleToSet(s, parent, parent)
		}
		if s[child] == 0 {
			AddEleToSet(s, child, s[parent])
		} else if s[child] == s[parent] {
			// 存在环
			result = e
		} else {
			JoinTwoSet(s, s[child], s[parent])
		}
	}
	if len(posibles) > 0 {
		// 可能存在环，但三肯定存在这有两个父节点的点
		for i := len(posibles) - 1; i > -1; i-- {
			if connectedWithoutLink(edges, posibles[i]) {
				result = posibles[i]
				break
			}
		}
	} else {
		// 存在环,但肯定不存在这有两个父节点的点

	}
	return result
}

func connectedWithoutLink(edges [][]int, rLink []int) bool {
	s := make(map[int]int, 1000)
	for _, e := range edges {
		if e[0] == rLink[0] && e[1] == rLink[1] {
			continue
		}
		child := e[1]
		parent := e[0]
		if !IsEleInSet(s, parent) {
			AddEleToSet(s, parent, parent)
		}
		if s[child] == 0 {
			AddEleToSet(s, child, s[parent])
		} else if s[child] == s[parent] {
			// 存在环
			return false
		} else {
			JoinTwoSet(s, s[child], s[parent])
		}
	}
	for _, val := range s {
		if val != s[1] {
			return false
		}
	}
	return true
}
