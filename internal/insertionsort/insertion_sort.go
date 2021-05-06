package insertionsort

const (
	// algorithm name
	name string = "insertion sort"
	// algorithm complexity
	complexity string = "O(n^2)"
)

// InsertionSort algorithm
type InsertionSort struct {
	name       string
	complexity string
}

// NewInsertionSort inits and returns InsertionSort struct
func NewInsertionSort() *InsertionSort {
	return &InsertionSort{
		name:       name,
		complexity: complexity,
	}
}

// Name returns algorithm name
func (m *InsertionSort) Name() string {
	return m.name
}

// Complexity returns algorithm complexity
func (m *InsertionSort) Complexity() string {
	return m.complexity
}

// Sort algorithm
func (m *InsertionSort) Sort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		current := arr[i]
		j := i

		for j > 0 && arr[j - 1] > current {
			arr[j] = arr[j - 1]
			j--
		}
		arr[j] = current
	}

	return arr
}
