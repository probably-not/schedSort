package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"

	"github.com/probably-not/schedSort"
)

func main() {
	unsorted := make([]int, 10)

	for i := 0; i < len(unsorted); i++ {
		unsorted[i] = rand.Intn(100)
	}

	start := time.Now()
	sorted := schedSort.SortInts(unsorted)
	end := time.Now()

	fmt.Println("==============================================================")
	fmt.Println("Sorted:", sort.IntsAreSorted(sorted))
	fmt.Println("Completed:", sorted, "in", end.Sub(start).String())
	fmt.Println("==============================================================")
}
