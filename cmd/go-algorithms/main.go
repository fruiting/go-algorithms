package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"time"

	"github.com/fruiting/go-algorithms/internal/algorithms"
	"github.com/fruiting/go-algorithms/internal/algorithms/search"
	sort2 "github.com/fruiting/go-algorithms/internal/algorithms/sort"
)

func main() {
	searchers := []algorithms.Searcher{
		search.NewBinarySearch(),
	}
	sorters := []algorithms.Sorter{
		sort2.NewBubbleSort(),
		sort2.NewQuickSort(),
		sort2.NewMergeSort(),
	}

	sortedArrTypes := make([][]int, 0, 4)
	unsortedArrTypes := make([][]int, 0, 4)
	for _, arrType := range []int{10, 100, 1000, 10000} {
		sortedArr := make([]int, 0, arrType)
		for i := 0; i < arrType; i++ {
			sortedArr = append(sortedArr, i)
		}
		sortedArrTypes = append(sortedArrTypes, sortedArr)

		rand.Seed(time.Now().Unix())
		unsortedArrTypes = append(unsortedArrTypes, rand.Perm(arrType))
	}

	for _, sorter := range sorters {
		fmt.Println(
			fmt.Sprintf(
				"====================\nname: %scomplexity: %s\n--------------------",
				reflect.TypeOf(sorter),
				sorter.Complexity(),
			),
		)
		for _, sortedArrType := range sortedArrTypes {
			now := time.Now()
			sorter.Sort(sortedArrType)
			elapsed := time.Since(now)

			fmt.Println(
				fmt.Sprintf(
					"array len: %d\nsort time: %d\n--------------------",
					len(sortedArrType),
					elapsed,
				),
			)
		}
		fmt.Println("====================")
	}

	for _, searcher := range searchers {
		fmt.Println(
			fmt.Sprintf(
				"====================\nname: %scomplexity: %s\n--------------------",
				reflect.TypeOf(searcher),
				searcher.Complexity(),
			),
		)
		for _, sortedArrType := range sortedArrTypes {
			now := time.Now()
			searcher.Search(sortedArrType, 18)
			elapsed := time.Since(now)

			fmt.Println(
				fmt.Sprintf(
					"array len: %d\nsort time: %d\n--------------------",
					len(sortedArrType),
					elapsed,
				),
			)
		}
		fmt.Println("====================")
	}
}
