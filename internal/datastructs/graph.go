package datastructs

type Graph [][]int

func NewGraph() Graph {
	//return [][]int{
	//	{1, 2},
	//	{2, 3},
	//	{4, 2},
	//}
	return [][]int{
		{1, 2},
		{5, 1},
		{1, 3},
		{1, 4},
	}
}

func (g Graph) FindCenter() int {
	m := make(map[int]int, 2)
	for k, v := range g {
		if k == 0 {
			m[v[0]]++
			m[v[1]]++
			continue
		}

		val, ok := m[v[0]]
		if ok {
			m[v[0]] = val + 1
		}
		val, ok = m[v[1]]
		if ok {
			m[v[1]] = val + 1
		}
	}

	v1 := m[g[0][0]]
	v2 := m[g[0][1]]
	if v1 == len(g) {
		return g[0][0]
	}
	if v2 == len(g) {
		return g[0][1]
	}

	return -1
}
