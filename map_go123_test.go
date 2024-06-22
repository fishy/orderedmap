//go:build go1.23

package orderedmap_test

import (
	"encoding/hex"
	"math/rand"
	"sync/atomic"
	"testing"

	ordered "go.yhsif.com/orderedmap"
)

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
	for i := 0; i < numKeys; i++ {
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
		t.Errorf("Expected iteration of 0 keys, got %d", counter)
	}
}
