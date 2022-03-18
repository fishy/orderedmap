package orderedmap_test

import (
	"fmt"

	"go.yhsif.com/orderedmap"
)

func Example() {
	var m orderedmap.Map[string, int]

	fmt.Println(m.Load("key1"))
	fmt.Println(m.LoadOrStore("key1", 1))
	fmt.Println(m.LoadOrStore("key1", 2))
	m.Store("key2", 2)
	fmt.Println(m.Load("key2"))
	m.Store("key1", 11)
	fmt.Println(m.Load("key1"))

	fmt.Println("Range1:")
	m.Range(func(k string, v int) bool {
		fmt.Println(k, v)
		return true
	})

	m.Delete("key1")
	m.Store("key1", 1)

	fmt.Println("Range2:")
	m.Range(func(k string, v int) bool {
		fmt.Println(k, v)
		return true
	})
	// Output:
	// 0 false
	// 1 false
	// 1 true
	// 2 true
	// 11 true
	// Range1:
	// key1 11
	// key2 2
	// Range2:
	// key2 2
	// key1 1
}
