package main

import "testing"

func TestQueue(t *testing.T) {
	q := Queue[string]{}

	q.enqueue("1")
	q.enqueue("2")
	q.enqueue("3")

	val, ok := q.dequeue()
	if !ok || val != "1" {
		t.Error()
	}

	val, ok = q.dequeue()
	if !ok || val != "2" {
		t.Error()
	}

	val, ok = q.dequeue()
	if !ok || val != "3" {
		t.Error()
	}

	_, ok = q.dequeue()
	if ok {
		t.Error()
	}
}
