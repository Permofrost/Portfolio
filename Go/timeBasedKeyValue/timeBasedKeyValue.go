package main

import (
	"fmt"
	"sort"
)

type TimeMap struct {
	m map[string][]data
}

type data struct {
	val       string
	timestamp int
}

func Constructor() TimeMap {
	return TimeMap{m: make(map[string][]data)}
}

func (tm *TimeMap) Set(key string, value string, timestamp int) {
	tm.m[key] = append(tm.m[key], data{val: value, timestamp: timestamp})
}

func (tm *TimeMap) Get(key string, timestamp int) string {
	if _, ok := tm.m[key]; !ok {
		return ""
	}
	i := sort.Search(len(tm.m[key]), func(i int) bool {
		return tm.m[key][i].timestamp > timestamp
	})
	if i > 0 {
		return tm.m[key][i-1].val
	}
	return ""
}

func main() {
	obj := Constructor()
	obj.Set("foo", "bar", 1)
	fmt.Println(obj.Get("foo", 1)) // Output: bar
	fmt.Println(obj.Get("foo", 3)) // Output: bar
	obj.Set("foo", "bar2", 4)
	fmt.Println(obj.Get("foo", 4)) // Output: bar2
	fmt.Println(obj.Get("foo", 5)) // Output: bar2
}
