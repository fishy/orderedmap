[![PkgGoDev](https://pkg.go.dev/badge/go.yhsif.com/orderedmap)](https://pkg.go.dev/go.yhsif.com/orderedmap)
[![Go Report Card](https://goreportcard.com/badge/go.yhsif.com/orderedmap)](https://goreportcard.com/report/go.yhsif.com/orderedmap)

# Go OrderedMap

An implementation of ordered map in Go.
An ordered map preserves the original insertion order when iterating.

It's implemented by wrapping a [`sync.Map`](https://pkg.go.dev/sync#Map)
and a [doubly linked list](https://pkg.go.dev/container/list#List).
The interface is intentionally kept the same with `sync.Map` to be used
interchangeably.

## Benchmark

This package comes with benchmark test to test against
builtin map and [`sync.Map`](https://pkg.go.dev/sync#Map):

```
$ go test -bench . -benchmem
goos: linux
goarch: amd64
pkg: go.yhsif.com/orderedmap
BenchmarkMap/DeleteEmpty/builtin-4              96833191                12.1 ns/op             0 B/op          0 allocs/op
BenchmarkMap/DeleteEmpty/sync-4                 79389284                15.0 ns/op             0 B/op          0 allocs/op
BenchmarkMap/DeleteEmpty/ordered-4              66165724                18.0 ns/op             0 B/op          0 allocs/op
BenchmarkMap/LoadEmpty/builtin-4                98423847                12.1 ns/op             0 B/op          0 allocs/op
BenchmarkMap/LoadEmpty/sync-4                   77803774                15.3 ns/op             0 B/op          0 allocs/op
BenchmarkMap/LoadEmpty/ordered-4                66868016                17.9 ns/op             0 B/op          0 allocs/op
BenchmarkMap/Store1/builtin-4                   13271734                88.9 ns/op            32 B/op          2 allocs/op
BenchmarkMap/Store1/sync-4                       7618335               157 ns/op              48 B/op          3 allocs/op
BenchmarkMap/Store1/ordered-4                    9210949               129 ns/op              64 B/op          3 allocs/op
BenchmarkMap/Store2/builtin-4                   12802537                94.4 ns/op            32 B/op          2 allocs/op
BenchmarkMap/Store2/sync-4                       7385720               162 ns/op              48 B/op          3 allocs/op
BenchmarkMap/Store2/ordered-4                    9084321               131 ns/op              64 B/op          3 allocs/op
BenchmarkMap/Update1/builtin-4                  13431109                89.5 ns/op            32 B/op          2 allocs/op
BenchmarkMap/Update1/sync-4                      7723395               156 ns/op              48 B/op          3 allocs/op
BenchmarkMap/Update1/ordered-4                   9167068               129 ns/op              64 B/op          3 allocs/op
BenchmarkMap/Load1/builtin-4                    44437922                26.9 ns/op             0 B/op          0 allocs/op
BenchmarkMap/Load1/sync-4                       38100231                31.4 ns/op             0 B/op          0 allocs/op
BenchmarkMap/Load1/ordered-4                    29717439                40.3 ns/op             0 B/op          0 allocs/op
BenchmarkMap/Load3/builtin-4                    76247488                15.6 ns/op             0 B/op          0 allocs/op
BenchmarkMap/Load3/sync-4                       59924874                20.0 ns/op             0 B/op          0 allocs/op
BenchmarkMap/Load3/ordered-4                    34028864                35.2 ns/op             0 B/op          0 allocs/op
BenchmarkLoadOrStoreStore/builtin-4             11663713               103 ns/op              32 B/op          2 allocs/op
BenchmarkLoadOrStoreStore/sync-4                 2815587               426 ns/op             376 B/op          8 allocs/op
BenchmarkLoadOrStoreStore/ordered-4              2138095               560 ns/op             520 B/op         11 allocs/op
BenchmarkLoadOrStoreLoad/builtin-4              14577663                81.9 ns/op            32 B/op          2 allocs/op
BenchmarkLoadOrStoreLoad/sync-4                 12393217                95.9 ns/op            32 B/op          2 allocs/op
BenchmarkLoadOrStoreLoad/ordered-4              12630664                94.0 ns/op            32 B/op          2 allocs/op
BenchmarkStoreThenDelete/builtin-4               9594682               126 ns/op              32 B/op          2 allocs/op
BenchmarkStoreThenDelete/sync-4                  2062261               580 ns/op             360 B/op          9 allocs/op
BenchmarkStoreThenDelete/ordered-4               1765857               680 ns/op             440 B/op         11 allocs/op
BenchmarkRange/10/builtin-4                      8392304               142 ns/op               0 B/op          0 allocs/op
BenchmarkRange/10/sync-4                         7705442               155 ns/op               0 B/op          0 allocs/op
BenchmarkRange/10/ordered-4                     34963722                34.2 ns/op             0 B/op          0 allocs/op
BenchmarkRange/100/builtin-4                      859714              1377 ns/op               0 B/op          0 allocs/op
BenchmarkRange/100/sync-4                         827470              1439 ns/op               0 B/op          0 allocs/op
BenchmarkRange/100/ordered-4                     3568192               337 ns/op               0 B/op          0 allocs/op
BenchmarkRange/1000/builtin-4                      77983             15334 ns/op               0 B/op          0 allocs/op
BenchmarkRange/1000/sync-4                         72570             16524 ns/op               0 B/op          0 allocs/op
BenchmarkRange/1000/ordered-4                     354681              3372 ns/op               0 B/op          0 allocs/op
PASS
ok      go.yhsif.com/orderedmap 51.884s
```

As you can see, all operations except `Range` are on-par or only slightly slower
than `sync.Map`, while `Range` is a lot faster.

## License

[BSD License](LICENSE).
