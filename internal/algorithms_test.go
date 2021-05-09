package internal

import (
	"github.com/fruiting/go-algorithms/internal/insertionsort"
	"github.com/fruiting/go-algorithms/internal/mergesort"
	"github.com/fruiting/go-algorithms/internal/quicksort"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

// TestSort_Ok tests sorting of all algorithms
func TestSort_Ok(t *testing.T) {
	unsortedArr := []int{5, 3, 1, 7, 6, 4, 9, 10, 2, 8}

	mergeSort := mergesort.NewMergeSort()
	insertionSort := insertionsort.NewInsertionSort()
	quickSort := quicksort.NewQuickSort()

	algorithms := []AlgorithmProcessor{
		mergeSort,
		insertionSort,
		quickSort,
	}

	reflects := []reflect.Value{
		reflect.ValueOf(mergeSort).Elem(),
		reflect.ValueOf(insertionSort).Elem(),
		reflect.ValueOf(quickSort).Elem(),
	}

	for key, algorithm := range algorithms {
		sortedArr := algorithm.Sort(unsortedArr)
		assert.Equal(t, sortedArr, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
		assert.Equal(t, algorithm.Name(), reflects[key].FieldByName("name").String())
		assert.Equal(t, algorithm.Complexity(), reflects[key].FieldByName("complexity").String())
	}
}
