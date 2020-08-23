// Package orderedmap provides an implementation of ordered map.
//
// An ordered map guarantees that the iteration order preserves the original
// insertion order.
// If an existing key is updated later,
// that doesn't change its iteration order.
// But if a key was deleted then later inserted again,
// the iteration order reflects its last insertion order.
package orderedmap
