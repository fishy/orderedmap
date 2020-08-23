[![PkgGoDev](https://pkg.go.dev/badge/github.com/fishy/orderedmap)](https://pkg.go.dev/github.com/fishy/orderedmap)
[![Go Report Card](https://goreportcard.com/badge/github.com/fishy/orderedmap)](https://goreportcard.com/report/github.com/fishy/orderedmap)

# Go OrderedMap

An implementation of ordered map in Go.
An ordered map preserves the original insertion order when iterating.

It's implemented by wrapping a [`sync.Map`](https://pkg.go.dev/sync?tab=doc#Map)
and a [doubly linked list](https://pkg.go.dev/container/list?tab=doc#List).
The interface is intentionally kept the same with `sync.Map` to be used
interchangeably.

## Benchmark

This package comes with benchmark test to test against
builtin map and [`sync.Map`](https://pkg.go.dev/sync?tab=doc#Map):

```
$ go test -test.benchmem -bench .
goos: linux
goarch: amd64
pkg: github.com/fishy/orderedmap
BenchmarkMap/DeleteEmpty/builtin-4              500000000                3.03 ns/op            0 B/op          0 allocs/op
BenchmarkMap/DeleteEmpty/sync-4                 200000000                6.61 ns/op            0 B/op          0 allocs/op
BenchmarkMap/DeleteEmpty/ordered-4              100000000               11.4 ns/op             0 B/op          0 allocs/op
BenchmarkMap/LoadEmpty/builtin-4                500000000                3.62 ns/op            0 B/op          0 allocs/op
BenchmarkMap/LoadEmpty/sync-4                   200000000                7.62 ns/op            0 B/op          0 allocs/op
BenchmarkMap/LoadEmpty/ordered-4                100000000               12.0 ns/op             0 B/op          0 allocs/op
BenchmarkMap/Store1/builtin-4                   20000000                91.4 ns/op            32 B/op          2 allocs/op
BenchmarkMap/Store1/sync-4                      10000000               135 ns/op              48 B/op          3 allocs/op
BenchmarkMap/Store1/ordered-4                   10000000               125 ns/op              64 B/op          3 allocs/op
BenchmarkMap/Store2/builtin-4                   20000000               102 ns/op              32 B/op          2 allocs/op
BenchmarkMap/Store2/sync-4                      10000000               139 ns/op              48 B/op          3 allocs/op
BenchmarkMap/Store2/ordered-4                   10000000               129 ns/op              64 B/op          3 allocs/op
BenchmarkMap/Update1/builtin-4                  20000000                94.5 ns/op            32 B/op          2 allocs/op
BenchmarkMap/Update1/sync-4                     10000000               141 ns/op              48 B/op          3 allocs/op
BenchmarkMap/Update1/ordered-4                  10000000               129 ns/op              64 B/op          3 allocs/op
BenchmarkMap/Load1/builtin-4                    50000000                25.9 ns/op             0 B/op          0 allocs/op
BenchmarkMap/Load1/sync-4                       50000000                30.9 ns/op             0 B/op          0 allocs/op
BenchmarkMap/Load1/ordered-4                    50000000                38.3 ns/op             0 B/op          0 allocs/op
BenchmarkMap/Load3/builtin-4                    100000000               17.4 ns/op             0 B/op          0 allocs/op
BenchmarkMap/Load3/sync-4                       100000000               22.1 ns/op             0 B/op          0 allocs/op
BenchmarkMap/Load3/ordered-4                    50000000                27.1 ns/op             0 B/op          0 allocs/op
BenchmarkLoadOrStoreStore/builtin-4             20000000               106 ns/op              32 B/op          2 allocs/op
BenchmarkLoadOrStoreStore/sync-4                 3000000               411 ns/op             376 B/op          8 allocs/op
BenchmarkLoadOrStoreStore/ordered-4              3000000               542 ns/op             520 B/op         11 allocs/op
BenchmarkLoadOrStoreLoad/builtin-4              20000000                81.1 ns/op            32 B/op          2 allocs/op
BenchmarkLoadOrStoreLoad/sync-4                 20000000                88.6 ns/op            32 B/op          2 allocs/op
BenchmarkLoadOrStoreLoad/ordered-4              20000000                96.1 ns/op            32 B/op          2 allocs/op
BenchmarkStoreThenDelete/builtin-4              10000000               132 ns/op              32 B/op          2 allocs/op
BenchmarkStoreThenDelete/sync-4                  5000000               254 ns/op              72 B/op          5 allocs/op
BenchmarkStoreThenDelete/ordered-4               5000000               321 ns/op             128 B/op          5 allocs/op
BenchmarkRange/10/builtin-4                     10000000               154 ns/op               0 B/op          0 allocs/op
BenchmarkRange/10/sync-4                        10000000               169 ns/op               0 B/op          0 allocs/op
BenchmarkRange/10/ordered-4                     50000000                34.9 ns/op             0 B/op          0 allocs/op
BenchmarkRange/100/builtin-4                     1000000              1463 ns/op               0 B/op          0 allocs/op
BenchmarkRange/100/sync-4                        1000000              1494 ns/op               0 B/op          0 allocs/op
BenchmarkRange/100/ordered-4                     5000000               342 ns/op               0 B/op          0 allocs/op
BenchmarkRange/1000/builtin-4                     100000             15698 ns/op               0 B/op          0 allocs/op
BenchmarkRange/1000/sync-4                        100000             17595 ns/op               0 B/op          0 allocs/op
BenchmarkRange/1000/ordered-4                     500000              3497 ns/op               0 B/op          0 allocs/op
PASS
ok      github.com/fishy/orderedmap     68.391s
```

As you can see, all operations except `Range` are on-par or only slightly slower
than `sync.Map`, while `Range` is a lot faster.

## License

[BSD License](https://github.com/fishy/orderedmap/blob/master/LICENSE).
