package sort

import (
	"golang.org/x/exp/constraints"
)

func SelectionSort[T constraints.Ordered](slice []T, isAsc bool) {
	n := len(slice)
	for i := 0; i < n-1; i++ {
		swapIndex := i
		for j := i + 1; j < n; j++ {
			if isAsc {
				if slice[j] < slice[swapIndex] {
					swapIndex = j
				}
			} else {
				if slice[j] > slice[swapIndex] {
					swapIndex = j
				}
			}
		}

		slice[i], slice[swapIndex] = slice[swapIndex], slice[i]
	}
}
