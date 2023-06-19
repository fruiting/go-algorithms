package main

import (
	"fmt"

	"github.com/fruiting/go-algorithms/internal/search"
)

func main() {
	b := search.NewBinarySearch()
	res := b.Run()
	fmt.Println(res)
}
