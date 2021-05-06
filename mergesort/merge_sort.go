package mergesort

// Mergesort algorithm
func Mergesort(arr []int) []int{
	merge := func(left, right []int) []int {
		lst:=make([]int, 0)
		for len(left) > 0 && len(right) > 0 {
			if left[0] < right[0] {
				lst = append(lst,left[0])
				left = left[1:]
			} else {
				lst = append(lst,right[0])
				right = right[1:]
			}
		}
		if len(left) > 0 {
			lst = append(lst,left...)
		}
		if  len(right) > 0 {
			lst = append(lst,right...)
		}

		return lst
	}

	length := len(arr)
	if length >= 2 {
		mid := length / 2
		left := arr[:mid]
		right := arr[mid:]
		sortedLeft := Mergesort(left)
		sortedRight := Mergesort(right)
		arr = merge(sortedLeft, sortedRight)
	}

	return arr
}
