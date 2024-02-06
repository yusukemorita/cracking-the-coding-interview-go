package main

// time complexity: O(N^2)
// space complexity: O(1)
//
// time complexity explanation
// when i=0, j takes N-1 steps
// when i=1, j takes N-2 steps
// when i=2, j takes N-3 steps
//   ...
// when i=N-2, j takes 1 step
// when i=N-1, j takes 0 steps
//
// so the total number of steps is
// (N-1) + (N-2) + ... + 1 = N(N-1)/2
func SelectionSort(slice []int) {
	for i := 0; i < len(slice); i++ {
		for j := i + 1; j < len(slice); j++ {
			if slice[i] > slice[j] {
				smaller := slice[j]
				slice[j] = slice[i]
				slice[i] = smaller
			}
		}
	}
}
