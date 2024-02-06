package main

// time complexity: O(N^2)
// space complexity: O(1)
func MergeSort(slice []int) {
	mergeSort(slice, 0, len(slice))
}

// startIndex is inclusive, endIndex is exclusive
func mergeSort(slice []int, startIndex, endIndex int) {
	if startIndex+1 >= endIndex {
		// section has length of 1 or less, nothing to sort
		return
	}

	// length is larger than 1
	newSectionLength := (endIndex - startIndex) / 2
	middleIndex := startIndex + newSectionLength
	mergeSort(slice, startIndex, middleIndex)
	mergeSort(slice, middleIndex, endIndex)
	merge(slice, startIndex, middleIndex, endIndex)
}

func merge(slice []int, startIndex, middleIndex, endIndex int) {
	slice1Index := startIndex
	slice2Index := middleIndex
	result := make([]int, endIndex-startIndex)
	resultIndex := 0

	for {
		completedSlice1 := slice1Index >= middleIndex
		completedSlice2 := slice2Index >= endIndex

		if completedSlice1 && completedSlice2 {
			break
		}

		if completedSlice1 {
			result[resultIndex] = slice[slice2Index]
			resultIndex++
			slice2Index++
			continue
		}

		if completedSlice2 {
			result[resultIndex] = slice[slice1Index]
			resultIndex++
			slice1Index++
			continue
		}

		item1 := slice[slice1Index]
		item2 := slice[slice2Index]

		if item1 <= item2 {
			result[resultIndex] = slice[slice1Index]
			resultIndex++
			slice1Index++
		} else {
			result[resultIndex] = slice[slice2Index]
			resultIndex++
			slice2Index++
		}
	}

	for i, item := range result {
		slice[i+startIndex] = item
	}
}
