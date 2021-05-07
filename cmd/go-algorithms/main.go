package main

import (
	"bufio"
	"fmt"
	"github.com/fruiting/go-algorithms/internal"
	"github.com/fruiting/go-algorithms/internal/mergesort"
	"github.com/fruiting/go-algorithms/internal/insertionsort"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

// numArr array of unsorted ints
var numArr []int

// wg WaitGroup
var wg sync.WaitGroup

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter numbers splitted by spaces: ")
	numbers, err := reader.ReadString('\n')
	if err != nil {
		panic(fmt.Errorf("can't read string: %w", err))
	}

	numbers = strings.TrimSuffix(numbers, "\n")
	numbers = strings.TrimSpace(numbers)
	stringArr := strings.Split(numbers, " ")
	for _, value := range stringArr {
		number, err := strconv.Atoi(value)
		if err != nil {
			panic(fmt.Errorf("can't convert string to int: %w", err))
		}

		numArr = append(numArr, number)
	}

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
	str = str + fmt.Sprintf("result: %v\n", sortedArr)
	str = str + fmt.Sprintf("time: %f\n", finish)
	str = str + "==============\n"
	fmt.Println(str)

	wg.Done()
}
