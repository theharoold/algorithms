package sort

import (
	"golang.org/x/exp/constraints"
)

func BubbleSort[T constraints.Ordered](slice []T, isAsc bool) {
	n := len(slice)
	for i := 0; i < n - 1; i++ {
		for j := 0; j < n - i - 1; j++ {
			if isAsc {
				if slice[j] > slice[j+1] {
					slice[j], slice[j+1] = slice[j+1], slice[j]
				}
			} else {
				if slice[j] < slice[j+1] {
					slice[j], slice[j+1] = slice[j+1], slice[j]
				}
			}
		}
	}
}