package sort

type MergeSort struct {
	complexity string
}

func NewMergeSort() *MergeSort {
	return &MergeSort{
		complexity: "O(n*logn)",
	}
}

func (s *MergeSort) Complexity() string {
	return s.complexity
}

func (s *MergeSort) Sort(arr []int) []int {
	// Если один элемент, то нечего сортировать
	if len(arr) < 2 {
		return arr
	}

	// Находим середину массива и делим его на две части
	middle := len(arr) / 2
	leftRes := s.Sort(arr[0:middle])
	rightRes := s.Sort(arr[middle:])

	// Запускаем рекурсию, в которой постоянно делим массив пополам и потом сортируем
	return s.merge(leftRes, rightRes)
}

func (s *MergeSort) merge(left, right []int) []int {
	res := make([]int, 0, len(left)+len(right))

	// Перебираем пока есть хотя бы один элемент в обеих частях
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			res = append(res, left[0])
			left = left[1:]
		} else {
			res = append(res, right[0])
			right = right[1:]
		}
	}

	return append(res, append(left, right...)...)
}
