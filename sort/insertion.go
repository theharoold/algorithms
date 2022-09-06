package sort

import (
	"golang.org/x/exp/constraints"
)

func InsertionSort[T constraints.Ordered](slice []T, isAsc bool) {
	n := len(slice)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if isAsc {
				if slice[j] < slice[j-1] {
					slice[j], slice[j-1] = slice[j-1], slice[j]
				}
			} else {
				if slice[j] > slice[j-1] {
					slice[j], slice[j-1] = slice[j-1], slice[j]
				}
			}
			j--
		}
	}
}