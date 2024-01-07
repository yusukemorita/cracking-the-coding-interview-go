package main

import (
	"hash/maphash"
	"testing"
)

func TestHashTable_Length(t *testing.T) {
	t.Run("when 1 linked list with one node", func(t *testing.T) {
		hashTable := HashTable{
			linkedLists: [1000]*LinkedListNode{{
				key:     "foo",
				value:   "valueForFoo",
				pointer: nil,
			}},
			seed: maphash.MakeSeed(),
		}

		length := hashTable.length()

		if length != 1 {
			t.Errorf("expected: 1, got: %d", length)
		}
	})

	t.Run("when 2 linked lists with 1 node", func(t *testing.T) {
		hashTable := HashTable{
			linkedLists: [1000]*LinkedListNode{
				{
					key:   "foo",
					value: "valueForFoo",
				},
				{
					key:   "bar",
					value: "valueForBar",
				},
			},
			seed: maphash.MakeSeed(),
		}

		length := hashTable.length()

		if length != 2 {
			t.Errorf("expected: 2, got: %d", length)
		}
	})

	t.Run("when 1 linked list with 2 nodes", func(t *testing.T) {
		hashTable := HashTable{
			linkedLists: [1000]*LinkedListNode{{
				key:   "foo",
				value: "valueForFoo",
				pointer: &LinkedListNode{
					key:   "bar",
					value: "valueForBar",
				},
			}},
			seed: maphash.MakeSeed(),
		}

		length := hashTable.length()

		if length != 2 {
			t.Errorf("expected: 2, got: %d", length)
		}
	})
}

func TestHashTable_Get(t *testing.T) {
	t.Run("when key is present", func(t *testing.T) {
		hashTable := NewHashTable()
		hashTable.set("foo", "valueForFoo")

		found, value := hashTable.get("foo")

		if found != true {
			t.Errorf("expected found: true, got: %v", found)
		}
		if value != "valueForFoo" {
			t.Errorf("expected %s, got: %s", "valueForFoo", value)
		}
	})

	t.Run("when key is not present", func(t *testing.T) {
		hashTable := NewHashTable()
		hashTable.set("foo", "valueForFoo")

		found, _ := hashTable.get("bar")

		if found != false {
			t.Errorf("expected found: false, got: %v", found)
		}
	})
}
