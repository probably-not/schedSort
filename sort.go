package schedSort

import "sync"

func SortInts(unsorted []int) []int {
	sorted := make([]int, len(unsorted))

	out := make(chan int, 1)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()

	sorter:
		i := 0
		for v := range out {
			sorted[i] = v
			i++
			if i >= len(sorted) {
				break
			}
		}

		if isSorted(sorted) {
			return
		}

		wg.Add(len(sorted))
		for _, v := range sorted {
			go func(x int) {
				defer wg.Done()
				out <- x
			}(v)
		}
		goto sorter
	}()

	wg.Add(len(unsorted))

	for _, v := range unsorted {
		go func(x int) {
			defer wg.Done()
			out <- x
		}(v)
	}

	wg.Wait()
	close(out)

	return sorted
}

func isSorted(s []int) bool {
	n := len(s)
	for i := n - 1; i > 0; i-- {
		if s[i] < s[i-1] {
			return false
		}
	}
	return true
}
