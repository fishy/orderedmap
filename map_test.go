package orderedmap_test

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"sync"
	"sync/atomic"
	"testing"

	ordered "go.yhsif.com/orderedmap"
)

func TestMap(t *testing.T) {
	var om ordered.Map[string, string]

	key1 := "key1"
	value1 := "value1"
	value1b := "new value1"
	key2 := "key2"
	value2 := "value2"
	key3 := "key3"
	value3 := "value3"
	value3b := "value3b"

	t.Run("LoadEmpty", func(t *testing.T) {
		v, ok := om.Load(key1)
		if v != "" {
			t.Errorf("Load value expected empty string, got %q", v)
		}
		if ok {
			t.Errorf("Load ok expected false, got %v", ok)
		}
	})

	om.Store(key1, value1)
	om.Store(key2, value2)
	om.Store(key1, value1b)

	t.Run("Load1", func(t *testing.T) {
		v, ok := om.Load(key1)
		if v != value1b {
			t.Errorf("Load value expected %q, got %q", value1b, v)
		}
		if !ok {
			t.Errorf("Load ok expected true, got %v", ok)
		}
	})

	t.Run("LoadOrStoreStore", func(t *testing.T) {
		v, loaded := om.LoadOrStore(key3, value3)
		if v != value3 {
			t.Errorf("Load value expected %q, got %q", value3, v)
		}
		if loaded {
			t.Errorf("Load loaded expected false, got %v", loaded)
		}
	})

	t.Run("LoadOrStoreLoad", func(t *testing.T) {
		v, loaded := om.LoadOrStore(key3, value3b)
		if v != value3 {
			t.Errorf("Load value expected %q, got %q", value3, v)
		}
		if !loaded {
			t.Errorf("Load loaded expected true, got %v", loaded)
		}
	})

	om.Delete(key3)

	t.Run("Load3", func(t *testing.T) {
		v, ok := om.Load(key3)
		if v != "" {
			t.Errorf("Load value expected empty string, got %q", v)
		}
		if ok {
			t.Errorf("Load ok expected false, got %v", ok)
		}
	})

	t.Run("LoadAndDelete2", func(t *testing.T) {
		v, loaded := om.LoadAndDelete(key2)
		if v != value2 {
			t.Errorf("Loaded value expected %q, got %q", value2, v)
		}
		if !loaded {
			t.Errorf("Loaded expected true, got %v", loaded)
		}
	})

	t.Run("LoadAndDelete2Again", func(t *testing.T) {
		v, loaded := om.LoadAndDelete(key2)
		if v != "" {
			t.Errorf("Loaded value expected empty string, got %q", v)
		}
		if loaded {
			t.Errorf("Loaded expected false, got %v", loaded)
		}
	})
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
	rangeFunc := func(t *testing.T, keys []string) func(k string, v int) bool {
		t.Helper()
		return func(k string, v int) bool {
			if v != rangeIndex {
				t.Errorf("Expected value %d, got %d", rangeIndex, v)
			}
			rangeIndex++
			if keys[v] != k {
				t.Errorf("Expected key %q at #%d, got %q", keys[v], v, k)
			}
			return true
		}
	}

	size := 1000
	var om ordered.Map[string, int]

	keys := make([]string, size)
	for i := range size {
		keys[i] = generateKey(t)
		om.Store(keys[i], i)
	}

	om.Range(rangeFunc(t, keys))
}

func TestDeleteOverRange(t *testing.T) {
	var om ordered.Map[int, any]
	numKeys := 10
	for i := range numKeys {
		om.Store(i, nil)
	}

	keyFuncCounter := func(counter *int) func(k int, v any) bool {
		*counter = 0
		return func(k int, v any) bool {
			*counter++
			om.Delete(k)
			return true
		}
	}

	var counter int
	om.Range(keyFuncCounter(&counter))
	if counter != numKeys {
		t.Errorf("Expected iteration of %d keys, got %d", numKeys, counter)
	}

	om.Range(keyFuncCounter(&counter))
	if counter != 0 {
		t.Errorf("Expected iteration of 0 keys, got %d", counter)
	}
}

/*
 * Benchmarks
 */

var sizes = []int{10, 100, 1000}

// builtin is a simple wrapper around builtin map to satisify Interface.
type builtin struct {
	m map[string]string
}

func newMap() *builtin {
	return &builtin{
		m: make(map[string]string),
	}
}

func (m *builtin) Delete(key string) {
	delete(m.m, key)
}

func (m *builtin) Load(key string) (value string, ok bool) {
	value, ok = m.m[key]
	return
}

func (m *builtin) LoadAndDelete(key string) (value string, ok bool) {
	value, ok = m.m[key]
	delete(m.m, key)
	return
}

func (m *builtin) LoadOrStore(key, value string) (actual string, loaded bool) {
	if v, ok := m.m[key]; ok {
		return v, true
	}
	m.m[key] = value
	return value, false
}

func (m *builtin) Range(f func(key, value string) bool) {
	for key, value := range m.m {
		if !f(key, value) {
			break
		}
	}
}

func (m *builtin) Store(key, value string) {
	m.m[key] = value
}

func BenchmarkMap(b *testing.B) {
	var sm sync.Map
	var om ordered.Map[string, string]
	m := newMap()

	key1 := "key1"
	value1 := "value1"
	value1b := "new value1"
	key2 := "key2"
	value2 := "value2"
	key3 := "key3"

	b.Run("DeleteEmpty", func(b *testing.B) {
		b.Run("builtin", func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				m.Delete(key1)
			}
		})
		b.Run("sync", func(b *testing.B) {
			b.ReportAllocs()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					sm.Delete(key1)
				}
			})
		})
		b.Run("ordered", func(b *testing.B) {
			b.ReportAllocs()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					om.Delete(key1)
				}
			})
		})
	})

	b.Run("LoadEmpty", func(b *testing.B) {
		b.Run("builtin", func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				m.Load(key1)
			}
		})
		b.Run("sync", func(b *testing.B) {
			b.ReportAllocs()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					sm.Load(key1)
				}
			})
		})
		b.Run("ordered", func(b *testing.B) {
			b.ReportAllocs()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					om.Load(key1)
				}
			})
		})
	})

	b.Run("Store1", func(b *testing.B) {
		b.Run("builtin", func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				m.Store(key1, value1)
			}
		})
		b.Run("sync", func(b *testing.B) {
			b.ReportAllocs()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					sm.Store(key1, value1)
				}
			})
		})
		b.Run("ordered", func(b *testing.B) {
			b.ReportAllocs()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					om.Store(key1, value1)
				}
			})
		})
	})

	b.Run("Store2", func(b *testing.B) {
		b.Run("builtin", func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				m.Store(key2, value2)
			}
		})
		b.Run("sync", func(b *testing.B) {
			b.ReportAllocs()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					sm.Store(key2, value2)
				}
			})
		})
		b.Run("ordered", func(b *testing.B) {
			b.ReportAllocs()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					om.Store(key2, value2)
				}
			})
		})
	})

	b.Run("Update1", func(b *testing.B) {
		b.Run("builtin", func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				m.Store(key1, value1b)
			}
		})
		b.Run("sync", func(b *testing.B) {
			b.ReportAllocs()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					sm.Store(key1, value1b)
				}
			})
		})
		b.Run("ordered", func(b *testing.B) {
			b.ReportAllocs()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					om.Store(key1, value1b)
				}
			})
		})
	})

	b.Run("Load1", func(b *testing.B) {
		b.Run("builtin", func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				m.Load(key1)
			}
		})
		b.Run("sync", func(b *testing.B) {
			b.ReportAllocs()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					sm.Load(key1)
				}
			})
		})
		b.Run("ordered", func(b *testing.B) {
			b.ReportAllocs()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					om.Load(key1)
				}
			})
		})
	})

	b.Run("Load3", func(b *testing.B) {
		b.Run("builtin", func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				m.Load(key3)
			}
		})
		b.Run("sync", func(b *testing.B) {
			b.ReportAllocs()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					sm.Load(key3)
				}
			})
		})
		b.Run("ordered", func(b *testing.B) {
			b.ReportAllocs()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					om.Load(key3)
				}
			})
		})
	})
}

func BenchmarkLoadOrStoreStore(b *testing.B) {
	key := "key"
	value := "value"

	b.Run("builtin", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			m := newMap()
			m.LoadOrStore(key, value)
		}
	})

	b.Run("sync", func(b *testing.B) {
		b.ReportAllocs()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				var m sync.Map
				m.LoadOrStore(key, value)
			}
		})
	})

	b.Run("ordered", func(b *testing.B) {
		b.ReportAllocs()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				var m ordered.Map[string, string]
				m.LoadOrStore(key, value)
			}
		})
	})
}

func BenchmarkLoadOrStoreLoad(b *testing.B) {
	key := "key"
	value1 := "value1"
	value2 := "value2"

	b.Run("builtin", func(b *testing.B) {
		b.ReportAllocs()
		m := newMap()
		m.LoadOrStore(key, value1)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			m.LoadOrStore(key, value2)
		}
	})

	b.Run("sync", func(b *testing.B) {
		b.ReportAllocs()
		var m sync.Map
		m.LoadOrStore(key, value1)
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				m.LoadOrStore(key, value2)
			}
		})
	})

	b.Run("ordered", func(b *testing.B) {
		b.ReportAllocs()
		var m ordered.Map[string, string]
		m.LoadOrStore(key, value1)
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				m.LoadOrStore(key, value2)
			}
		})
	})
}

func BenchmarkStoreThenDelete(b *testing.B) {
	key := "key"
	value := "value"

	b.Run("builtin", func(b *testing.B) {
		b.ReportAllocs()
		m := newMap()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			m.Store(key, value)
			m.Delete(key)
		}
	})

	b.Run("sync", func(b *testing.B) {
		b.ReportAllocs()
		var m sync.Map
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				m.Store(key, value)
				m.Delete(key)
			}
		})
	})

	b.Run("ordered", func(b *testing.B) {
		b.ReportAllocs()
		var m ordered.Map[string, string]
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				m.Store(key, value)
				m.Delete(key)
			}
		})
	})
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
	rangeFunc := func(k, v any) bool {
		return true
	}
	rangeFuncTyped := func(k, v string) bool {
		return true
	}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("%d", size), func(b *testing.B) {
			var sm sync.Map
			var om ordered.Map[string, string]
			m := newMap()

			keys := make([]string, size)
			for i := range size {
				keys[i] = generateKey(b)
				m.Store(keys[i], keys[i])
				sm.Store(keys[i], keys[i])
				om.Store(keys[i], keys[i])
			}

			b.Run("builtin", func(b *testing.B) {
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					m.Range(rangeFuncTyped)
				}
			})

			b.Run("sync", func(b *testing.B) {
				b.ReportAllocs()
				b.RunParallel(func(pb *testing.PB) {
					for pb.Next() {
						sm.Range(rangeFunc)
					}
				})
			})

			b.Run("ordered", func(b *testing.B) {
				b.ReportAllocs()
				b.RunParallel(func(pb *testing.PB) {
					for pb.Next() {
						om.Range(rangeFuncTyped)
					}
				})
			})
		})
	}
}

func TestAll(t *testing.T) {
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

	size := 1000
	var om ordered.Map[string, int]

	keys := make([]string, size)
	for i := range size {
		keys[i] = generateKey(t)
		om.Store(keys[i], i)
	}

	for k, v := range om.All() {
		if v != rangeIndex {
			t.Errorf("Expected value %d, got %d", rangeIndex, v)
		}
		rangeIndex++
		if keys[v] != k {
			t.Errorf("Expected key %q at #%d, got %q", keys[v], v, k)
		}
	}
}

func TestDeleteOverAll(t *testing.T) {
	var om ordered.Map[int, any]
	numKeys := 10
	for i := range numKeys {
		om.Store(i, nil)
	}

	var counter atomic.Int64
	for k := range om.All() {
		counter.Add(1)
		om.Delete(k)
	}
	if got := counter.Load(); got != int64(numKeys) {
		t.Errorf("Expected iteration of %d keys, got %d", numKeys, got)
	}

	counter.Store(0)
	for k := range om.All() {
		counter.Add(1)
		om.Delete(k)
	}
	if got := counter.Load(); got != 0 {
		t.Errorf("Expected iteration of 0 keys, got %d", got)
	}
}
