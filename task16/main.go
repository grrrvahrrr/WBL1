package main

import (
	"fmt"
	"math/rand"
)

func main() {
	slice := []int{3, 123, 455, 345, 2, 5, 65, 567, 45, 345, 67678, 56433, 3, 45, 456}
	sorted := quicksort(slice)
	fmt.Println(sorted)
}

func quicksort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	pivot := rand.Int() % len(a)

	a[pivot], a[right] = a[right], a[pivot]

	for i := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	quicksort(a[:left])
	quicksort(a[left+1:])

	return a
}
