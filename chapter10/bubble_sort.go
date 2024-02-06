package main

// space complexity: O(1)
// time complexity: O(N^2)
// In the worst case scenario, the slice is in reverse order.
//   e.g. initial: 5, 4, 3, 2, 1
// In the first loop through the slice, four swaps will happen,
// and the 5 will be moved to the end of the slice.
//   after loop1: 4, 3, 2, 1, 5
// Then, the 4 will move to the second to last place, with three swaps.
//   after loop2: 3, 2, 1, 4, 5
// This will continue until the slice is sorted, so the total number of
// swaps will be
//   4 + 3 + 2 + 1
// If the number of items is N, this can be written as
//   (N-1) + (N-2) + ... + 1 = N(N-1)/2
// Therefore, the time complexity is O(N^2)
func BubbleSort(slice []int) {
	if len(slice) <= 1 {
		// nothing to sort
		return
	}

	for {
		changed := false

		for i := 0; i <= len(slice)-2; i++ {
			if slice[i] > slice[i+1] {
				// swap the two elements so that the smaller one
				// comes first
				smaller := slice[i+1]
				slice[i+1] = slice[i]
				slice[i] = smaller

				changed = true
			}
		}

		if changed == false {
			break
		}
	}
}
