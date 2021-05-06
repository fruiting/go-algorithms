package mergesort

const (
	// algorithm name
	name string = "merge sort"
	// algorithm complexity
	complexity string = "O(nLog(n))"
)

// MergeSort algorithm
type MergeSort struct {
	name       string
	complexity string
}

// NewMergeSort inits and returs MergeSort struct
func NewMergeSort() *MergeSort {
	return &MergeSort{
		name:       name,
		complexity: complexity,
	}
}

// Name returns algorithm name
func (m *MergeSort) Name() string {
	return m.name
}

// Complexity returns algorithm complexity
func (m *MergeSort) Complexity() string {
	return m.complexity
}

// Sort algorithm
func (m *MergeSort) Sort(arr []int) []int {
	merge := func(left, right []int) []int {
		lst := make([]int, 0)
		for len(left) > 0 && len(right) > 0 {
			if left[0] < right[0] {
				lst = append(lst, left[0])
				left = left[1:]
			} else {
				lst = append(lst, right[0])
				right = right[1:]
			}
		}
		if len(left) > 0 {
			lst = append(lst, left...)
		}
		if len(right) > 0 {
			lst = append(lst, right...)
		}

		return lst
	}

	length := len(arr)
	if length >= 2 {
		mid := length / 2
		left := arr[:mid]
		right := arr[mid:]
		sortedLeft := m.Sort(left)
		sortedRight := m.Sort(right)
		arr = merge(sortedLeft, sortedRight)
	}

	return arr
}
