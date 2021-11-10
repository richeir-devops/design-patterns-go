package strategy

import "testing"

func Test01(t *testing.T) {
	dataset := []int{1, 5, 4, 3, 2, 8}

	sorter := &Sorter{sortStrategy: &BubbleSortStrategy{}}
	sorter.doSort(dataset)

	sorter = &Sorter{sortStrategy: &QuickSortStrategy{}}
	sorter.doSort(dataset)
}
