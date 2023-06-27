package search

type BinarySearch struct {
	complexity string
}

func NewBinarySearch() *BinarySearch {
	return &BinarySearch{
		complexity: "O(log n)",
	}
}

func (s *BinarySearch) Complexity() string {
	return s.complexity
}

func (s *BinarySearch) Search(sortedArr []int, val int) int {
	start := 0                // определяем первый ключ массива
	end := len(sortedArr) - 1 // определяем последний ключ массива

	for start <= end {
		middle := (start + end) / 2 // определяем ключ в середине массива в текущей итерации (не базового массива)

		if sortedArr[middle] == val {
			return middle // если переданное значение равно значению в центре массива, то возвращаем ключ
		} else if sortedArr[middle] < val {
			start = middle + 1 // если переданное значение больше значения в центре массива, то сдвигаем поиск вправо от центра массива (не базового, а в текущей итерации)
		} else {
			end = middle - 1 // если переданное значение меньше значения в центре массива, то сдвигаем поиск влево от центра массива (не базового, а в текущей итерации)
		}
	}

	return -1
}
