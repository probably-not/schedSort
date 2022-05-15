package schedSort

import (
	"math/rand"
	"testing"
)

var sortedInts []int

// BenchmarkSortInts-10    	       1	30393834375 ns/op	3431821632 B/op	142811224 allocs/op
func BenchmarkSortInts(b *testing.B) {
	b.StopTimer()
	unsorted := make([]int, 10)

	for i := 0; i < len(unsorted); i++ {
		unsorted[i] = rand.Intn(1000000)
	}

	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		sortedInts = SortInts(unsorted)
	}
}
