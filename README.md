# schedSort

A brilliant sorting algorithm that uses Go's high-performance concurrency primitives to sort in parallel.

As all of you know, Go is a high-performance language with very powerful concurrency primitives.
But, some of you may not know, that Go's standard library barely uses these primitives, such as in it's sort package.

My sorting algorithm uses Go's scheduler and goroutines to do all of the sorting with goroutines and waitgroups.

## Notes

This is a joke sorting algorithm, and the whole README is a joke.
It's based on a stupid idea that I had to try and see how random Go's scheduler is.
Please do not use this in production or take any of what is written here as the truth.