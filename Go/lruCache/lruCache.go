package main

import (
	"container/list"
	"fmt"
)

type LRUCache struct {
	capacity int
	cache    map[int]*list.Element
	pages    *list.List
}

type page struct {
	key   int
	value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*list.Element),
		pages:    list.New(),
	}
}

func (lru *LRUCache) Get(key int) int {
	if element, ok := lru.cache[key]; ok {
		lru.pages.MoveToFront(element)
		return element.Value.(*page).value
	}
	return -1
}

func (lru *LRUCache) Put(key int, value int) {
	if element, ok := lru.cache[key]; ok {
		lru.pages.MoveToFront(element)
		element.Value.(*page).value = value
	} else {
		if lru.pages.Len() >= lru.capacity {
			delete(lru.cache, lru.pages.Back().Value.(*page).key)
			lru.pages.Remove(lru.pages.Back())
		}
		element := lru.pages.PushFront(&page{key, value})
		lru.cache[key] = element
	}
}

func main() {
	obj := Constructor(2)
	obj.Put(1, 1)
	obj.Put(2, 2)
	fmt.Println(obj.Get(1)) // returns 1
	obj.Put(3, 3)           // evicts key 2
	fmt.Println(obj.Get(2)) // returns -1 (not found)
	obj.Put(4, 4)           // evicts key 1
	fmt.Println(obj.Get(1)) // returns -1 (not found)
	fmt.Println(obj.Get(3)) // returns 3
	fmt.Println(obj.Get(4)) // returns 4
}
