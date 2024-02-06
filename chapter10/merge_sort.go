package main

// time complexity: O(NlogN)
//   - each merge takes x operations, where x is the
//     sum of the length of the two merged slices
//   - since `mergeSort` is a recursive function, and in each recursion
//     the slices are halved, we can think of the structure of the recursion
//     like a binary tree
//   - at each level of the binary tree, the total number of items merged together
//     is always N. At the top level merge, two slices of length N/2 are merged,
//     and at the bottom level, N separate slices of length 1 are merged into N/2 slices.
//   - therefore, the time complexity is
//     (depth of binary tree) * (length of items in the original slice)
//   - since the depth of the binary tree will correlate to logN, this means the time complexity is
//     O(logN * N)
//   - ref: https://www.youtube.com/watch?v=MBxqi6GrfN8
// space complexity: O(N)
//   - for the helper array, which has a length of N
func MergeSort(slice []int) {
	helper := make([]int, len(slice))
	mergeSort(slice, helper, 0, len(slice))
}

// startIndex is inclusive, endIndex is exclusive
func mergeSort(slice, helper []int, startIndex, endIndex int) {
	if startIndex+1 >= endIndex {
		// section has length of 1 or less, nothing to sort
		return
	}

	// length is larger than 1
	newSectionLength := (endIndex - startIndex) / 2
	middleIndex := startIndex + newSectionLength
	mergeSort(slice, helper, startIndex, middleIndex)
	mergeSort(slice, helper, middleIndex, endIndex)
	merge(slice, helper, startIndex, middleIndex, endIndex)
}

func merge(slice, helper []int, startIndex, middleIndex, endIndex int) {
	slice1Index := startIndex
	slice2Index := middleIndex

	// copy items in both sections into the helper
	for i := startIndex; i < endIndex; i++ {
		helper[i] = slice[i]
	}
	// to track where to insert the next item
	sliceIndex := startIndex

	for {
		completedSlice1 := slice1Index >= middleIndex
		completedSlice2 := slice2Index >= endIndex

		if completedSlice1 && completedSlice2 {
			break
		}

		if completedSlice1 {
			slice[sliceIndex] = helper[slice2Index]
			sliceIndex++
			slice2Index++
			continue
		}

		if completedSlice2 {
			slice[sliceIndex] = helper[slice1Index]
			sliceIndex++
			slice1Index++
			continue
		}

		item1 := helper[slice1Index]
		item2 := helper[slice2Index]

		if item1 <= item2 {
			slice[sliceIndex] = item1
			slice1Index++
		} else {
			slice[sliceIndex] = item2
			slice2Index++
		}
		sliceIndex++
	}
}
