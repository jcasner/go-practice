package kthLargest

import (
	"fmt"
	"sort"
)

// FindElement writes the kth largest element in an array of integers
// to the screen.
func FindElement(list []int, k int) {
	sort.Ints(list)
	fmt.Printf("The %dth largest element in %v is %d\n", k, list, list[len(list)-k])
}
