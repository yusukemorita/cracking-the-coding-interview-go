package main

// time complexity:
//   best case: O(NlogN)
//     reasoning is similar to merge sort.
//     If the pivot is chosen every time so that it is the median,
//     then every recursive call will cause the section to be split evenly
//     into two parts, so the structure is analogous to a binary tree.
//     The calls to quicksort in one level of the binary tree results in
//     a number of operations that correlate to N.
//     The number of levels in the tree correlates to logN, so the time complexity
//     is O(NlogN).
//   worst: O(n^2)
// space complexity: O(logN)
//   space required on the stack for recursive calls
func QuickSort(slice []int) {
	quickSort(slice, 0, len(slice))
}

// startIndex is inclusive, endIndex is exclusive
func quickSort(slice []int, startIndex, endIndex int) {
	if endIndex-startIndex <= 1 {
		// nothing to sort
		return
	}

	// pick a pivot halfway between startIndex and endIndex
	pivotIndex := (endIndex + startIndex) / 2
	pivot := slice[pivotIndex]

	// swap pivot with last item in slice
	slice[pivotIndex] = slice[endIndex-1]
	slice[endIndex-1] = pivot

	var leftIndex int

	for {
		leftIndex = startIndex
		// find first item from the left of the slice that is larger than pivot
		for {
			if slice[leftIndex] > pivot || leftIndex >= endIndex-1 {
				break
			}

			leftIndex++
		}

		// find first item from the right of the slice that is smaller than or equal to pivot
		// start with last item in slice, excluding the pivot
		rightIndex := endIndex - 2
		for {
			if slice[rightIndex] <= pivot || rightIndex <= startIndex {
				break
			}

			rightIndex--
		}

		if leftIndex >= rightIndex {
			break
		}

		// swap values at leftIndex and rightIndex
		tmp := slice[leftIndex]
		slice[leftIndex] = slice[rightIndex]
		slice[rightIndex] = tmp
	}

	// swap pivot with item at left index position, so that pivot is in "correct" position
	// where all items left of pivot are smaller than or equal to pivot, and
	// all items right of pivot are larger than pivot
	slice[endIndex-1] = slice[leftIndex]
	slice[leftIndex] = pivot

	// sort all items left of pivot
	quickSort(slice, startIndex, leftIndex)
	// sort all items right of pivot
	quickSort(slice, leftIndex+1, endIndex)
}
