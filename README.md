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
BenchmarkMap/DeleteEmpty/builtin-4              94249322                12.1 ns/op             0 B/op          0 allocs/op
BenchmarkMap/DeleteEmpty/sync-4                 79167520                15.1 ns/op             0 B/op          0 allocs/op
BenchmarkMap/DeleteEmpty/ordered-4              66040401                17.9 ns/op             0 B/op          0 allocs/op
BenchmarkMap/LoadEmpty/builtin-4                89343954                13.3 ns/op             0 B/op          0 allocs/op
BenchmarkMap/LoadEmpty/sync-4                   71315845                16.7 ns/op             0 B/op          0 allocs/op
BenchmarkMap/LoadEmpty/ordered-4                65312764                18.2 ns/op             0 B/op          0 allocs/op
BenchmarkMap/Store1/builtin-4                   13019566                91.2 ns/op            32 B/op          2 allocs/op
BenchmarkMap/Store1/sync-4                       7633965               157 ns/op              48 B/op          3 allocs/op
BenchmarkMap/Store1/ordered-4                    9198397               131 ns/op              64 B/op          3 allocs/op
BenchmarkMap/Store2/builtin-4                   13053826                92.2 ns/op            32 B/op          2 allocs/op
BenchmarkMap/Store2/sync-4                       7638517               157 ns/op              48 B/op          3 allocs/op
BenchmarkMap/Store2/ordered-4                    9242421               130 ns/op              64 B/op          3 allocs/op
BenchmarkMap/Update1/builtin-4                  13433424                91.2 ns/op            32 B/op          2 allocs/op
BenchmarkMap/Update1/sync-4                      7610348               158 ns/op              48 B/op          3 allocs/op
BenchmarkMap/Update1/ordered-4                   9210226               131 ns/op              64 B/op          3 allocs/op
BenchmarkMap/Load1/builtin-4                    41253802                29.1 ns/op             0 B/op          0 allocs/op
BenchmarkMap/Load1/sync-4                       37282363                38.0 ns/op             0 B/op          0 allocs/op
BenchmarkMap/Load1/ordered-4                    28864215                36.2 ns/op             0 B/op          0 allocs/op
BenchmarkMap/Load3/builtin-4                    69284424                17.2 ns/op             0 B/op          0 allocs/op
BenchmarkMap/Load3/sync-4                       55919145                21.3 ns/op             0 B/op          0 allocs/op
BenchmarkMap/Load3/ordered-4                    51765396                22.8 ns/op             0 B/op          0 allocs/op
BenchmarkLoadOrStoreStore/builtin-4             11607195               103 ns/op              32 B/op          2 allocs/op
BenchmarkLoadOrStoreStore/sync-4                 2843066               422 ns/op             376 B/op          8 allocs/op
BenchmarkLoadOrStoreStore/ordered-4              2148310               559 ns/op             520 B/op         11 allocs/op
BenchmarkLoadOrStoreLoad/builtin-4              14578209                82.0 ns/op            32 B/op          2 allocs/op
BenchmarkLoadOrStoreLoad/sync-4                 13124654                90.3 ns/op            32 B/op          2 allocs/op
BenchmarkLoadOrStoreLoad/ordered-4              13257388                92.3 ns/op            32 B/op          2 allocs/op
BenchmarkStoreThenDelete/builtin-4               9425214               127 ns/op              32 B/op          2 allocs/op
BenchmarkStoreThenDelete/sync-4                  2097183               572 ns/op             360 B/op          9 allocs/op
BenchmarkStoreThenDelete/ordered-4               3656547               331 ns/op             128 B/op          5 allocs/op
BenchmarkRange/10/builtin-4                      8751439               136 ns/op               0 B/op          0 allocs/op
BenchmarkRange/10/sync-4                         7667242               156 ns/op               0 B/op          0 allocs/op
BenchmarkRange/10/ordered-4                     34921269                34.1 ns/op             0 B/op          0 allocs/op
BenchmarkRange/100/builtin-4                      933817              1272 ns/op               0 B/op          0 allocs/op
BenchmarkRange/100/sync-4                         822381              1451 ns/op               0 B/op          0 allocs/op
BenchmarkRange/100/ordered-4                     3572973               336 ns/op               0 B/op          0 allocs/op
BenchmarkRange/1000/builtin-4                      78133             15293 ns/op               0 B/op          0 allocs/op
BenchmarkRange/1000/sync-4                         72427             16418 ns/op               0 B/op          0 allocs/op
BenchmarkRange/1000/ordered-4                     351825              3389 ns/op               0 B/op          0 allocs/op
PASS
ok      go.yhsif.com/orderedmap 51.620s
```

As you can see, all operations except `Range` are on-par or only slightly slower
than `sync.Map`, while `Range` is a lot faster.

## License

[BSD License](LICENSE).
