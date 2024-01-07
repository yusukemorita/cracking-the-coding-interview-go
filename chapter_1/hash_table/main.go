package main

import "hash/maphash"

func main() {
	print("hello")
}

const arrayLength = 1000

// key: string, value: string
type HashTable struct {
	linkedLists [arrayLength]*LinkedListNode // items in slice may be nil
	seed        maphash.Seed
}

func NewHashTable() HashTable {
	return HashTable{
		linkedLists: [arrayLength]*LinkedListNode{},
		seed:        maphash.MakeSeed(),
	}
}

func (hashTable *HashTable) length() int {
	length := 0
	for _, linkedList := range hashTable.linkedLists {
		if linkedList == nil {
			continue
		}

		length = length + linkedList.length()
	}

	return length
}

func (hashTable *HashTable) get(key string) (found bool, value string) {
	hash := maphash.String(hashTable.seed, key)
	index := hash % arrayLength
	linkedList := hashTable.linkedLists[index]

	if linkedList == nil {
		return false, ""
	}

	return linkedList.get(key)
}

func (hashTable *HashTable) set(key string, value string) {
	hash := maphash.String(hashTable.seed, key)
	index := hash % arrayLength
	linkedList := hashTable.linkedLists[index]

	if linkedList == nil {
		hashTable.linkedLists[index] = &LinkedListNode{key: key, value: value, pointer: nil}
	} else {
		// TODO: what if the key is overwritten?
		// TODO: what if there is already more than one item in the linkedList?
		linkedList.pointer = &LinkedListNode{key: key, value: value, pointer: nil}
	}
}

type LinkedListNode struct {
	key     string
	value   string
	pointer *LinkedListNode
}

func (node LinkedListNode) length() int {
	if node.pointer == nil {
		return 1
	}

	return node.pointer.length() + 1
}

func (node LinkedListNode) get(key string) (found bool, value string) {
	if node.key == key {
		return true, node.value
	}

	if node.pointer == nil {
		return false, ""
	}

	return node.pointer.get(key)
}
