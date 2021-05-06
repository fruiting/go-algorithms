package mergesort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSort_Ok(t *testing.T) {
	algorithm := NewMergeSort()
	sortedArr := algorithm.Sort([]int{5, 3, 1, 7, 6, 4, 9, 10, 2, 8})

	assert.Equal(t, sortedArr, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	assert.Equal(t, algorithm.Name(), name)
	assert.Equal(t, algorithm.Complexity(), complexity)
}
