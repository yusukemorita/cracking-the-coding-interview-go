package main

import "fmt"

// re: https://go.dev/blog/slices-intro

type ArrayList[T comparable] struct {
	array []T
}

func NewArrayList[T comparable]() ArrayList[T] {
	return ArrayList[T]{
		array: make([]T, 0, 1),
	}
}

func (arrayList ArrayList[T]) get(index uint) T {
	fmt.Printf("%v, length: %v, cap: %v\n", arrayList.array, len(arrayList.array), cap(arrayList.array))
	return arrayList.array[index]
}

func (arrayList *ArrayList[T]) push(item T) {
	requiredCap := len(arrayList.array) + 1
	currentCap := cap(arrayList.array)
	fmt.Printf("%v, length: %v, cap: %v\n", arrayList.array, len(arrayList.array), currentCap)

	// if array does not have enough capacity to add a new item,
	// initialize a new array with twice the capacity
	if (currentCap < requiredCap) {
		newArray := make([]T, requiredCap, 2*requiredCap)

		// copy all items from old array to new array
		for index, item :=  range arrayList.array {
			newArray[index] = item
		}

		arrayList.array = newArray
	}

	arrayList.array = arrayList.array[:requiredCap]
	fmt.Printf("inserting item into index %d\n", len(arrayList.array))
	arrayList.array[requiredCap - 1] = item

	fmt.Printf("%v, length: %v, cap: %v\n", arrayList.array, len(arrayList.array), cap(arrayList.array))
}
