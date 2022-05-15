package main

import (
	"fmt"
	"sort"
	"time"

	"github.com/probably-not/schedSort"
)

func main() {
	start := time.Now()
	unsorted := []int{12, 123412, 1323, 12341241234, 123421341232, 124134, 12341234, 12341234}
	sorted := schedSort.SortInts(unsorted)
	end := time.Now()

	fmt.Println("==============================================================")
	fmt.Println("Sorted:", sort.IntsAreSorted(sorted))
	fmt.Println("Completed:", sorted, "in", end.Sub(start).String())
	fmt.Println("==============================================================")
}
