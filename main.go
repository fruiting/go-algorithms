package main

import (
	"algorythms/mergesort"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

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
	var numArr []int
	for _, value := range stringArr {
		number, err := strconv.Atoi(value)
		if err != nil {
			panic(fmt.Errorf("can't convert string to int: %w", err))
		}

		numArr = append(numArr, number)
	}

	fmt.Println("==============")
	start := time.Now()
	mergeSortArr := mergesort.Mergesort(numArr)
	fmt.Println("merge sort")
	fmt.Printf("result: %v\n", mergeSortArr)
	fmt.Println("time: %f", time.Since(start).Seconds())
	fmt.Println("==============")
}