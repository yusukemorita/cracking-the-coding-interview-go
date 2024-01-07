package main

import (
	"fmt"
	"testing"
)

func TestArrayList_Append(t *testing.T) {
	t.Run("appends 1 item", func(t *testing.T) {
		arrayList := NewArrayList[string]()
		arrayList.push("hi")

		fmt.Println("pushed")

		first := arrayList.get(0)
		fmt.Println("got")
		if first != "hi" {
			t.Errorf("expected first item to be 'hi', got %s", first)
		}
	})

	t.Run("appends 10 items", func(t *testing.T) {
		arrayList := NewArrayList[string]()
		arrayList.push("1")
		arrayList.push("2")
		arrayList.push("3")
		arrayList.push("4")
		arrayList.push("5")
		arrayList.push("6")
		arrayList.push("7")
		arrayList.push("8")
		arrayList.push("9")
		arrayList.push("10")
		arrayList.push("11")

		tenth := arrayList.get(10)
		if tenth != "11" {
			t.Errorf("expected 10th item to be '11', got %s", tenth)
		}
	})
}
