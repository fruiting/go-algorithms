package sort

type BubbleSort struct {
	complexity string
}

func NewBubbleSort() *BubbleSort {
	return &BubbleSort{
		complexity: "O(n²)",
	}
}

func (s *BubbleSort) Complexity() string {
	return s.complexity
}

func (s *BubbleSort) Sort(arr []int) []int {
	// определяем количество элементов в массиве и идем от последнего к первому
	// таким образом мы понимаем сколько элементов надо будет сравнить в каждой итерации
	for j := len(arr) - 1; j > 0; j-- {
		// в каждой итерации делаем еще один цикл, в котором уже сравниваем нужное количество элементов
		for i := 0; i < j; i++ {
			// если первый элемент больше последующего, то меняем их местами
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
		}
	}

	return arr
}
