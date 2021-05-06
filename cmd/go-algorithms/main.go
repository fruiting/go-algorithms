package main

import (
	"bufio"
	"fmt"
	"github.com/fruiting/go-algorithms/internal"
	"github.com/fruiting/go-algorithms/internal/insertionsort"
	"github.com/fruiting/go-algorithms/internal/mergesort"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

// maxSize of array
const maxSize int = 10000000

// numArr array of unsorted ints
var numArr []int

// wg WaitGroup
var wg sync.WaitGroup

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter array size: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		panic(fmt.Errorf("can't read string: %w", err))
	}

	input = strings.TrimSuffix(input, "\n")
	input = strings.TrimSpace(input)
	size, err := strconv.Atoi(input)
	if err != nil {
		panic(fmt.Errorf("can't convert string to int: %w", err))
	}

	if size > maxSize {
		fmt.Println(fmt.Sprintf("Max length of array can't be more than %d", maxSize))
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
	}

	wg.Add(len(algorithms))
	for _, algorithm := range algorithms {
		go executeAlgorithm(algorithm)
	}

	wg.Wait()
}

func executeAlgorithm(algorithm internal.AlgorithmProcessor) {
	start := time.Now()
	sortedArr := algorithm.Sort(numArr)
	finish := time.Since(start).Seconds()

	str := "==============\n"
	str = str + fmt.Sprintf("name: %s\n", algorithm.Name())
	str = str + fmt.Sprintf("complexity: %s\n", algorithm.Complexity())
	str = str + fmt.Sprintf(
		"result: [%d, %d, %d, ..., %d, %d, %d]\n",
		sortedArr[0],
		sortedArr[1],
		sortedArr[2],
		sortedArr[len(sortedArr) - 3],
		sortedArr[len(sortedArr) - 2],
		sortedArr[len(sortedArr) - 1],
	)
	str = str + fmt.Sprintf("time: %f\n", finish)
	str = str + "==============\n"
	fmt.Println(str)

	wg.Done()
}
