package main

import "fmt"

// time complexity: O(N^2)
// space complexity: O(1)
func MergeSort(slice []int) []int {
	return mergeSort(slice)
}

func mergeSort(slice []int) []int {
	if len(slice) <= 1 {
		// nothing to sort
		return slice
	}

	// length is larger than 1
	half1 := slice[:len(slice)/2]
	half2 := slice[len(slice)/2:]
	fmt.Printf("half1: %v, half2: %v\n", half1, half2)

	return merge(mergeSort(half1), mergeSort(half2))
}

func merge(slice1, slice2 []int) []int {
	slice1Index := 0
	slice2Index := 0
	var result [] int

	for {
		completedSlice1 := slice1Index >= len(slice1)
		completedSlice2 := slice2Index >= len(slice2)

		if completedSlice1 && completedSlice2 {
			break
		}

		if completedSlice1 {
			result = append(result, slice2[slice2Index])
			slice2Index++
			continue
		}

		if completedSlice2 {
			result = append(result, slice1[slice1Index])
			slice1Index++
			continue
		}

		item1 := slice1[slice1Index]
		item2 := slice2[slice2Index]

		if item1 <= item2 {
			result = append(result, slice1[slice1Index])
			slice1Index++
		} else {
			result = append(result, slice2[slice2Index])
			slice2Index++
		}
	}

	return result
}
