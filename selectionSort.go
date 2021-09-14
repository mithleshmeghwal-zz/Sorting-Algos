package main

import "fmt"

// Selection Sort
// Time Complexity: O(n2) as there are two nested loops.
// Auxiliary Space: O(1)
// The good thing about selection sort is
// it never makes more than O(n) swaps and can be useful
// when memory write is a costly operation.

//Example
// 64 25 12 22 11
// 11 25 12 22 64
// 11 12 25 22 64
// 11 12 22 25 64
// 11 12 22 25 64

func main() {
	A := [][]int{
		{64, 25, 12, 22, 11},
	}
	for i := range A {
		fmt.Println(Solution(A[i]))
	}
}

func Solution(A []int) []int {
	n := len(A)
	// len is n
	// so n-1 pass is required to sort it
	var min_idx int
	for i := 0; i < n-1; i++ {
		min_idx = i // consider as intial minumum  and then run liner search to find the min in unsorted list
		for j := i + 1; j < n; j++ {
			if A[j] < A[min_idx] {
				min_idx = j
			}
		}
		swap(&A[min_idx], &A[i])
	}
	return A
}

func swap(xp *int, yp *int) {
	temp := *xp
	*xp = *yp
	*yp = temp
}
