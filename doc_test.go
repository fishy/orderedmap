package orderedmap_test

import (
	"fmt"

	"go.yhsif.com/orderedmap"
)

func Example() {
	var m orderedmap.Map

	fmt.Println(m.Load("key1"))
	fmt.Println(m.LoadOrStore("key1", "value1"))
	fmt.Println(m.LoadOrStore("key1", "value2"))
	m.Store("key2", "value2")
	fmt.Println(m.Load("key2"))
	m.Store("key1", "new value1")
	fmt.Println(m.Load("key1"))

	fmt.Println("Range1:")
	m.Range(func(k, v interface{}) bool {
		fmt.Println(k, v)
		return true
	})

	m.Delete("key1")
	m.Store("key1", "value1")

	fmt.Println("Range2:")
	m.Range(func(k, v interface{}) bool {
		fmt.Println(k, v)
		return true
	})
	// Output:
	// <nil> false
	// value1 false
	// value1 true
	// value2 true
	// new value1 true
	// Range1:
	// key1 new value1
	// key2 value2
	// Range2:
	// key2 value2
	// key1 value1
}
