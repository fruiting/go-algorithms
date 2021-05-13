package main

import (
	"bufio"
	"fmt"
	"github.com/fruiting/go-algorithms/internal"
	"github.com/fruiting/go-algorithms/internal/insertionsort"
	"github.com/fruiting/go-algorithms/internal/mergesort"
	"github.com/fruiting/go-algorithms/internal/quicksort"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	// minSize of array
	minSize int = 5
	// maxPrintAllSize to print all elements of sorted array
	maxPrintAllSize int = 10
	// maxSize of array
	maxSize int = 100000000
)

// numArr array of unsorted integers
var numArr []int

// wg WaitGroup
var wg sync.WaitGroup

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter array size: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(fmt.Errorf("can't read string: %w", err))
		return
	}

	input = strings.TrimSuffix(input, "\n")
	input = strings.TrimSpace(input)
	size, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println(fmt.Errorf("can't convert string to int: %w", err))
		return
	}

	if size < minSize || size > maxSize {
		fmt.Println(
			fmt.Sprintf(
				"Min length of array can't be less than %d. Max length of array can't be more than %d",
				minSize,
				maxSize,
			),
		)
		return
	}

	for i := 1; i <= size; i++ {
		numArr = append(numArr, i)
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(numArr), func(i, j int) {
		numArr[i], numArr[j] = numArr[j], numArr[i]
	})

	algorithms := []internal.AlgorithmProcessor{
		mergesort.NewMergeSort(),
		insertionsort.NewInsertionSort(),
		quicksort.NewQuickSort(),
	}

	wg.Add(len(algorithms))
	algoChan := make(chan string)
	quit := make(chan bool)

	for _, algorithm := range algorithms {
		go func(algorithm internal.AlgorithmProcessor) {
			algoChan <- algorithmInfo(algorithm)
			wg.Done()
		}(algorithm)
	}

	go func() {
		wg.Wait()
		quit <- true
	}()

	for {
		select {
		case info := <-algoChan:
			fmt.Println(info)
		case <-quit:
			close(algoChan)
			close(quit)
			return
		}
	}
}

func algorithmInfo(algorithm internal.AlgorithmProcessor) string {
	start := time.Now()
	sortedArr := algorithm.Sort(numArr)
	finish := time.Since(start).Seconds()

	str := "==============\n"
	str = str + fmt.Sprintf("name: %s\n", algorithm.Name())
	str = str + fmt.Sprintf("complexity: %s\n", algorithm.Complexity())

	if len(sortedArr) <= maxPrintAllSize {
		str = str + fmt.Sprintf("result: %v\n", sortedArr)
	} else {
		str = str + fmt.Sprintf(
			"result: [%d, %d, %d, ..., %d, %d, %d]\n",
			sortedArr[0],
			sortedArr[1],
			sortedArr[2],
			sortedArr[len(sortedArr)-3],
			sortedArr[len(sortedArr)-2],
			sortedArr[len(sortedArr)-1],
		)
	}

	str = str + fmt.Sprintf("time: %f\n", finish)
	str = str + "==============\n"

	return str
}
