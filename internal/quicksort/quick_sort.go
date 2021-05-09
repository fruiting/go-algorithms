package quicksort

const (
	// algorithm name
	name string = "quick sort"
	// algorithm complexity
	complexity string = "O(nLog(n))"
)

// QuickSort algorithm
type QuickSort struct {
	name       string
	complexity string
}

// NewQuickSort inits and returns QuickSort struct
func NewQuickSort() *QuickSort {
	return &QuickSort{
		name:       name,
		complexity: complexity,
	}
}

// Name returns algorithm name
func (s *QuickSort) Name() string {
	return s.name
}

// Complexity returns algorithm complexity
func (s *QuickSort) Complexity() string {
	return s.complexity
}

// Sort returns sorted array by specific algorithm
func (s *QuickSort) Sort(arr []int) []int {
	partition := func(arr []int, low int, high int) int {
		pivot := arr[high]
		i := low - 1
		for j := low; j <= high-1; j++ {
			if arr[j] < pivot {
				i++
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
		arr[i+1], arr[high] = arr[high], arr[i+1]
		return i + 1
	}

	var quickSort func(arr []int, low int, high int) []int
	quickSort = func(arr []int, low int, high int) []int {
		if low > high {
			return arr
		}

		pi := partition(arr, low, high)
		quickSort(arr, low, pi-1)
		quickSort(arr, pi+1, high)

		return arr
	}

	return quickSort(arr, 0, len(arr)-1)
}
