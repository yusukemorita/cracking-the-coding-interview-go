package main

import "time"

type Queue[T comparable] struct {
	first *LinkedListNode[T]
	last  *LinkedListNode[T]
}

func (q *Queue[T]) dequeue() (T, bool) {
	if q.first == nil {
		var zero T
		return zero, false
	}

	first := q.first
	q.first = q.first.next

	return first.value, true
}

func (q *Queue[T]) enqueue(item T) {
	node := LinkedListNode[T]{
		value: item,
	}

	if q.first == nil {
		q.first = &node
	}

	if q.last != nil {
		q.last.next = &node
	}

	q.last = &node
}

func (q *Queue[T]) peek() (T, bool) {
	if q.first == nil {
		var zero T
		return zero, false
	}

	return first.value, true
}

type Animal struct {
	arrivalTime time.Time
	species     string // "dog" || "cat"
}

type AnimalShelter struct {
	dogQueue Queue[Animal]
	catQueue Queue[Animal]
}

func (shelter *AnimalShelter) DequeueDog() (Animal, bool) {
	return shelter.dogQueue.dequeue()
}

func (shelter *AnimalShelter) DequeueCat() (Animal, bool) {
	return shelter.catQueue.dequeue()
}

func (shelter *AnimalShelter) DequeueAny() (Animal, bool) {
	lastCat, catFound := shelter.catQueue.peek()
	lastDog, dogFound := shelter.dogQueue.peek()

	if !catFound && !dogFound {
		return Animal{}, false
	}

	if !catFound && dogFound {
		return shelter.dogQueue.dequeue()
	}

	if catFound && !dogFound {
		return shelter.dogQueue.dequeue()
	}

	if lastCat.arrivalTime.Before(lastDog.arrivalTime) {
		// cat arrived first
		return shelter.catQueue.dequeue()
	} else {
		// dog arrived first
		return shelter.dogQueue.dequeue()
	}
}
