package main

import (
	"fmt"
	"goCache/lru"
)

type String string

func (s String) Len() int {
	return len(s)
}
func main() {
	cache := lru.NewCache(8, nil)
	cache.Insert("key", String("123"))
	if val, ok := cache.Get("key"); ok {
		fmt.Println("yes", val.(String))
	}
}
