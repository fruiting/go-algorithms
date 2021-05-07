package internal

import (
	"github.com/fruiting/go-algorithms/internal/mergesort"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSort_Ok(t *testing.T) {
	unsortedArr := []int{5, 3, 1, 7, 6, 4, 9, 10, 2, 8}
	algorithms := []AlgorithmProcessor{
		mergesort.NewMergeSort(),
	}
	algorithmsNames := []string{
		"merge sort",
		"insertion sort",
	}
	algorithmsComplexities := []string{
		"O(nLog(n))",
		"O(n^2)",
	}

	for key, algorithm := range algorithms {
		sortedArr := algorithm.Sort(unsortedArr)
		assert.Equal(t, sortedArr, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
		assert.Equal(t, algorithm.Name(), algorithmsNames[key])
		assert.Equal(t, algorithm.Complexity(), algorithmsComplexities[key])
	}
}
