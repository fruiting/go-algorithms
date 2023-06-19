package search

type BinarySearch struct {
	complexity string
}

func NewBinarySearch() *BinarySearch {
	return &BinarySearch{
		complexity: "O (log2 n)",
	}
}

func (s *BinarySearch) Run() int {
	sortedArr := []int{1, 3, 5, 7, 8, 9, 10, 13, 15, 16, 18, 20, 22, 25, 27}
	val := 18

	return s.search(sortedArr, val)
}

func (s *BinarySearch) search(sortedArr []int, val int) int {
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
