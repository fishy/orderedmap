[![GoDoc](https://godoc.org/github.com/fishy/orderedmap?status.svg)](https://godoc.org/github.com/fishy/orderedmap)
[![Go Report Card](https://goreportcard.com/badge/github.com/fishy/orderedmap)](https://goreportcard.com/report/github.com/fishy/orderedmap)

# Go OrderedMap

An implementation of ordered map in Go.
An ordered map preserves the original insertion order when iterating.

It's implemented by wrapping a [`sync.Map`](https://godoc.org/sync#Map)
and a [doubly linked list](https://godoc.org/container/list#List).
The interface is intentionally kept the same with `sync.Map` to be used
interchangeably.

## Benchmark

This package comes with benchmark test to test against
[`sync.Map`](https://godoc.org/sync#Map):

```
$ go test -bench .
goos: linux
goarch: amd64
pkg: github.com/fishy/orderedmap
BenchmarkMap/DeleteEmpty/sync-4                 200000000                6.70 ns/op
BenchmarkMap/DeleteEmpty/ordered-4              100000000               11.5 ns/op
BenchmarkMap/LoadEmpty/sync-4                   200000000                7.62 ns/op
BenchmarkMap/LoadEmpty/ordered-4                100000000               11.9 ns/op
BenchmarkMap/Store1/sync-4                      10000000               138 ns/op
BenchmarkMap/Store1/ordered-4                   10000000               137 ns/op
BenchmarkMap/Store2/sync-4                      10000000               139 ns/op
BenchmarkMap/Store2/ordered-4                   10000000               136 ns/op
BenchmarkMap/Update1/sync-4                     10000000               138 ns/op
BenchmarkMap/Update1/ordered-4                  10000000               139 ns/op
BenchmarkMap/Load1/sync-4                       50000000                30.4 ns/op
BenchmarkMap/Load1/ordered-4                    50000000                38.1 ns/op
BenchmarkMap/Load3/sync-4                       100000000               21.8 ns/op
BenchmarkMap/Load3/ordered-4                    50000000                26.5 ns/op
BenchmarkLoadOrStoreStore/sync-4                 3000000               406 ns/op
BenchmarkLoadOrStoreStore/ordered-4              3000000               544 ns/op
BenchmarkLoadOrStoreLoad/sync-4                 20000000                90.3 ns/op
BenchmarkLoadOrStoreLoad/ordered-4              20000000                95.3 ns/op
BenchmarkStoreThenDelete/sync-4                  5000000               256 ns/op
BenchmarkStoreThenDelete/ordered-4               5000000               338 ns/op
BenchmarkRange/10/sync-4                        10000000               166 ns/op
BenchmarkRange/10/ordered-4                     50000000                32.7 ns/op
BenchmarkRange/100/sync-4                        1000000              1530 ns/op
BenchmarkRange/100/ordered-4                     5000000               311 ns/op
BenchmarkRange/1000/sync-4                        100000             17022 ns/op
BenchmarkRange/1000/ordered-4                     300000              3458 ns/op
PASS
ok      github.com/fishy/orderedmap     44.991s
```

As you can see, all operations except `Range` are on-par or only slightly slower
than `sync.Map`, while `Range` is a lot faster.

## License

[BSD License](https://github.com/fishy/orderedmap/blob/master/LICENSE).
