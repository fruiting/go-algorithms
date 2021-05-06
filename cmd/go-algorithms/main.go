package main

import (
	"bufio"
	"fmt"
	"github.com/fruiting/go-algorithms/internal"
	"github.com/fruiting/go-algorithms/internal/mergesort"
	"os"
	"strconv"
	"strings"
	"time"
)

// numArr array of unsorted ints
var numArr []int

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
	}

	for _, algorithm := range algorithms {
		fmt.Println("==============")
		start := time.Now()
		sortedArr := algorithm.Sort(numArr)
		fmt.Printf("name: %s\n", algorithm.Name())
		fmt.Printf("complexity: %s\n", algorithm.Complexity())
		fmt.Printf("result: %v\n", sortedArr)
		fmt.Printf("time: %f\n", time.Since(start).Seconds())
		fmt.Println("==============")
	}
}
