package internal

// AlgorithmProcessor processes with array
type AlgorithmProcessor interface {
	// Name returns algorithm name
	Name() string
	// Complexity returns algorithm complexity
	Complexity() string
	// Sort returns sorted array by specific algorithm
	Sort(arr []int) []int
}
