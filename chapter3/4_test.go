package main

import "testing"

func TestMyQueue(t *testing.T) {
	t.Run("pops first item that was pushed", func(t *testing.T) {
		myQueue := NewMyQueue()

		myQueue.push("1")
		myQueue.push("2")
		myQueue.push("3")

		value, ok := myQueue.pop()
		if !ok || value != "1" {
			t.Errorf("expected ok: true, value: 1, got ok: %v, value: %s", ok, value)
		}

		value, ok = myQueue.pop()
		if !ok || value != "2" {
			t.Errorf("expected ok: true, value: 2, got ok: %v, value: %s", ok, value)
		}

		value, ok = myQueue.pop()
		if !ok || value != "3" {
			t.Errorf("expected ok: true, value: 3, got ok: %v, value: %s", ok, value)
		}

		_, ok = myQueue.pop()
		if ok {
			t.Errorf("expected ok: false, got ok: %v", ok)
		}
	})

}
