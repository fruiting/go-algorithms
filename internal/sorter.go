package internal

// AlgorithmProcessor processes with array
type AlgorithmProcessor interface {
	// Sort returns sorted array by specific algorithm
	Sort(arr []int) [] int
	// Name returns algorithm name
	Name() string
	// Complexity returns algorithm complexity
	Complexity() string
}
