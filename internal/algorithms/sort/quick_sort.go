package sort

type QuickSort struct {
	complexity string
}

func NewQuickSort() *QuickSort {
	return &QuickSort{
		complexity: "O(n*logn)",
	}
}

func (s *QuickSort) Complexity() string {
	return s.complexity
}

func (s *QuickSort) Sort(arr []int) []int {
	// Если в массиве один элемент, то и нечего сортировать
	if len(arr) < 2 {
		return arr
	}

	pivot := arr[0]                   // Берем опорный элемент
	left := make([]int, 0, len(arr))  // Все что меньше опорного
	right := make([]int, 0, len(arr)) // Все что больше опорного
	same := make([]int, 0, len(arr))  // Все что равно опорному

	for i := 0; i < len(arr); i++ {
		if arr[i] < pivot {
			left = append(left, arr[i])
		} else if arr[i] > pivot {
			right = append(right, arr[i])
		} else if arr[i] == pivot {
			same = append(same, arr[i])
		}
	}

	// Рекурсивно повторяем для элементов со значениями менее и более опорного пока не останется по одному элементу
	leftRes := s.Sort(left)
	rightRes := s.Sort(right)

	// Конкатенируем и возвращаем
	return append(append(leftRes, same...), rightRes...)
}
