package orderedmap_test

import (
	"encoding/hex"
	"fmt"
	"math/rand"
	"sync"
	"testing"

	ordered "github.com/fishy/orderedmap"
)

var sizes = []int{10, 100, 1000}

func TestMap(t *testing.T) {
	var om ordered.Map

	key1 := "key1"
	value1 := "value1"
	value1b := "new value1"
	key2 := "key2"
	value2 := "value2"
	key3 := "key3"
	value3 := "value3"
	value3b := "value3b"

	t.Run(
		"LoadEmpty",
		func(t *testing.T) {
			v, ok := om.Load(key1)
			if v != nil {
				t.Errorf("Load value expected nil, got %v", v)
			}
			if ok {
				t.Errorf("Load ok expected false, got %v", ok)
			}
		},
	)

	om.Store(key1, value1)
	om.Store(key2, value2)
	om.Store(key1, value1b)

	t.Run(
		"Load1",
		func(t *testing.T) {
			v, ok := om.Load(key1)
			if v != value1b {
				t.Errorf("Load value expected %v, got %v", value1b, v)
			}
			if !ok {
				t.Errorf("Load ok expected true, got %v", ok)
			}
		},
	)

	t.Run(
		"LoadOrStoreStore",
		func(t *testing.T) {
			v, loaded := om.LoadOrStore(key3, value3)
			if v != value3 {
				t.Errorf("Load value expected %v, got %v", value3, v)
			}
			if loaded {
				t.Errorf("Load loaded expected false, got %v", loaded)
			}
		},
	)

	t.Run(
		"LoadOrStoreLoad",
		func(t *testing.T) {
			v, loaded := om.LoadOrStore(key3, value3b)
			if v != value3 {
				t.Errorf("Load value expected %v, got %v", value3, v)
			}
			if !loaded {
				t.Errorf("Load loaded expected true, got %v", loaded)
			}
		},
	)

	om.Delete(key3)

	t.Run(
		"Load3",
		func(t *testing.T) {
			v, ok := om.Load(key3)
			if v != nil {
				t.Errorf("Load value expected nil, got %v", v)
			}
			if ok {
				t.Errorf("Load ok expected false, got %v", ok)
			}
		},
	)
}

func TestRange(t *testing.T) {
	keySize := 32
	p := make([]byte, keySize)
	generateKey := func(t *testing.T) string {
		t.Helper()
		_, err := rand.Read(p)
		if err != nil {
			t.Fatalf("Failed to generate key: %v", err)
		}
		return hex.EncodeToString(p)
	}
	rangeIndex := 0
	rangeFunc := func(t *testing.T, keys []string) func(k, v interface{}) bool {
		t.Helper()
		return func(k, v interface{}) bool {
			i := v.(int)
			if i != rangeIndex {
				t.Errorf("Expected value %d, got %d", rangeIndex, i)
			}
			rangeIndex++
			if keys[i] != k {
				t.Errorf("Expected key %q at #%d, got %q", keys[i], i, k)
			}
			return true
		}
	}

	size := 1000
	var om ordered.Map

	keys := make([]string, size)
	for i := 0; i < size; i++ {
		keys[i] = generateKey(t)
		om.Store(keys[i], i)
	}

	om.Range(rangeFunc(t, keys))
}

func BenchmarkMap(b *testing.B) {
	var sm sync.Map
	var om ordered.Map

	key1 := "key1"
	value1 := "value1"
	value1b := "new value1"
	key2 := "key2"
	value2 := "value2"
	key3 := "key3"

	b.Run(
		"DeleteEmpty",
		func(b *testing.B) {
			b.Run(
				"sync",
				func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						sm.Delete(key1)
					}
				},
			)
			b.Run(
				"ordered",
				func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						om.Delete(key1)
					}
				},
			)
		},
	)

	b.Run(
		"LoadEmpty",
		func(b *testing.B) {
			b.Run(
				"sync",
				func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						sm.Load(key1)
					}
				},
			)
			b.Run(
				"ordered",
				func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						om.Load(key1)
					}
				},
			)
		},
	)

	b.Run(
		"Store1",
		func(b *testing.B) {
			b.Run(
				"sync",
				func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						sm.Store(key1, value1)
					}
				},
			)
			b.Run(
				"ordered",
				func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						om.Store(key1, value1)
					}
				},
			)
		},
	)

	b.Run(
		"Store2",
		func(b *testing.B) {
			b.Run(
				"sync",
				func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						sm.Store(key2, value2)
					}
				},
			)
			b.Run(
				"ordered",
				func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						om.Store(key2, value2)
					}
				},
			)
		},
	)

	b.Run(
		"Update1",
		func(b *testing.B) {
			b.Run(
				"sync",
				func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						sm.Store(key1, value1b)
					}
				},
			)
			b.Run(
				"ordered",
				func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						om.Store(key1, value1b)
					}
				},
			)
		},
	)

	b.Run(
		"Load1",
		func(b *testing.B) {
			b.Run(
				"sync",
				func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						sm.Load(key1)
					}
				},
			)
			b.Run(
				"ordered",
				func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						om.Load(key1)
					}
				},
			)
		},
	)

	b.Run(
		"Load3",
		func(b *testing.B) {
			b.Run(
				"sync",
				func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						sm.Load(key3)
					}
				},
			)
			b.Run(
				"ordered",
				func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						om.Load(key3)
					}
				},
			)
		},
	)
}

func BenchmarkLoadOrStoreStore(b *testing.B) {
	key := "key"
	value := "value"

	b.Run(
		"sync",
		func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				var m sync.Map
				m.LoadOrStore(key, value)
			}
		},
	)

	b.Run(
		"ordered",
		func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				var m ordered.Map
				m.LoadOrStore(key, value)
			}
		},
	)
}

func BenchmarkLoadOrStoreLoad(b *testing.B) {
	key := "key"
	value1 := "value1"
	value2 := "value2"

	b.Run(
		"sync",
		func(b *testing.B) {
			var m sync.Map
			m.LoadOrStore(key, value1)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				m.LoadOrStore(key, value2)
			}
		},
	)

	b.Run(
		"ordered",
		func(b *testing.B) {
			var m ordered.Map
			m.LoadOrStore(key, value1)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				m.LoadOrStore(key, value2)
			}
		},
	)
}

func BenchmarkStoreThenDelete(b *testing.B) {
	key := "key"
	value := "value"

	b.Run(
		"sync",
		func(b *testing.B) {
			var m sync.Map
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				m.Store(key, value)
				m.Delete(key)
			}
		},
	)

	b.Run(
		"ordered",
		func(b *testing.B) {
			var m ordered.Map
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				m.Store(key, value)
				m.Delete(key)
			}
		},
	)
}

func BenchmarkRange(b *testing.B) {
	keySize := 32
	p := make([]byte, keySize)
	generateKey := func(b *testing.B) string {
		b.Helper()
		_, err := rand.Read(p)
		if err != nil {
			b.Fatalf("Failed to generate key: %v", err)
		}
		return hex.EncodeToString(p)
	}
	rangeFunc := func(k, v interface{}) bool {
		return true
	}

	for _, size := range sizes {
		b.Run(
			fmt.Sprintf("%d", size),
			func(b *testing.B) {
				var sm sync.Map
				var om ordered.Map

				keys := make([]string, size)
				for i := 0; i < size; i++ {
					keys[i] = generateKey(b)
					sm.Store(keys[i], i)
					om.Store(keys[i], i)
				}

				b.Run(
					"sync",
					func(b *testing.B) {
						for i := 0; i < b.N; i++ {
							sm.Range(rangeFunc)
						}
					},
				)

				b.Run(
					"ordered",
					func(b *testing.B) {
						for i := 0; i < b.N; i++ {
							om.Range(rangeFunc)
						}
					},
				)
			},
		)
	}
}
