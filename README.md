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
$ go test -bench .
goos: linux
goarch: amd64
pkg: go.yhsif.com/orderedmap
cpu: Intel(R) Core(TM) i5-7260U CPU @ 2.20GHz
BenchmarkMap/DeleteEmpty/builtin-4              577244734                2.078 ns/op           0 B/op          0 allocs/op
BenchmarkMap/DeleteEmpty/sync-4                 100000000               11.73 ns/op            0 B/op          0 allocs/op
BenchmarkMap/DeleteEmpty/ordered-4              83530353                13.54 ns/op            0 B/op          0 allocs/op
BenchmarkMap/LoadEmpty/builtin-4                540074115                2.220 ns/op           0 B/op          0 allocs/op
BenchmarkMap/LoadEmpty/sync-4                   100000000               12.04 ns/op            0 B/op          0 allocs/op
BenchmarkMap/LoadEmpty/ordered-4                79147016                14.94 ns/op            0 B/op          0 allocs/op
BenchmarkMap/Store1/builtin-4                   100000000               10.48 ns/op            0 B/op          0 allocs/op
BenchmarkMap/Store1/sync-4                       8478454               143.3 ns/op            48 B/op          3 allocs/op
BenchmarkMap/Store1/ordered-4                   16938922                69.60 ns/op           32 B/op          1 allocs/op
BenchmarkMap/Store2/builtin-4                   100000000               11.29 ns/op            0 B/op          0 allocs/op
BenchmarkMap/Store2/sync-4                       8290592               142.4 ns/op            48 B/op          3 allocs/op
BenchmarkMap/Store2/ordered-4                   16879738                69.88 ns/op           32 B/op          1 allocs/op
BenchmarkMap/Update1/builtin-4                  100000000               10.48 ns/op            0 B/op          0 allocs/op
BenchmarkMap/Update1/sync-4                      8209876               144.8 ns/op            48 B/op          3 allocs/op
BenchmarkMap/Update1/ordered-4                  17162760                68.94 ns/op           32 B/op          1 allocs/op
BenchmarkMap/Load1/builtin-4                    276140088                4.145 ns/op           0 B/op          0 allocs/op
BenchmarkMap/Load1/sync-4                       49011340                23.15 ns/op            0 B/op          0 allocs/op
BenchmarkMap/Load1/ordered-4                    38655788                31.23 ns/op            0 B/op          0 allocs/op
BenchmarkMap/Load3/builtin-4                    98014983                12.84 ns/op            0 B/op          0 allocs/op
BenchmarkMap/Load3/sync-4                       70591510                16.10 ns/op            0 B/op          0 allocs/op
BenchmarkMap/Load3/ordered-4                    54961591                21.47 ns/op            0 B/op          0 allocs/op
BenchmarkLoadOrStoreStore/builtin-4             57805004                21.63 ns/op            0 B/op          0 allocs/op
BenchmarkLoadOrStoreStore/sync-4                 2946840               407.2 ns/op           376 B/op          8 allocs/op
BenchmarkLoadOrStoreStore/ordered-4              2284989               523.0 ns/op           504 B/op         10 allocs/op
BenchmarkLoadOrStoreLoad/builtin-4              279418766                4.450 ns/op           0 B/op          0 allocs/op
BenchmarkLoadOrStoreLoad/sync-4                 14499127                81.00 ns/op           32 B/op          2 allocs/op
BenchmarkLoadOrStoreLoad/ordered-4              39903387                30.23 ns/op            0 B/op          0 allocs/op
BenchmarkStoreThenDelete/builtin-4              39218394                28.86 ns/op            0 B/op          0 allocs/op
BenchmarkStoreThenDelete/sync-4                  2178794               552.1 ns/op           360 B/op          9 allocs/op
BenchmarkStoreThenDelete/ordered-4               1907264               631.5 ns/op           424 B/op         10 allocs/op
BenchmarkRange/10/builtin-4                      9251668               136.9 ns/op             0 B/op          0 allocs/op
BenchmarkRange/10/sync-4                         7960039               158.0 ns/op             0 B/op          0 allocs/op
BenchmarkRange/10/ordered-4                     40579003                29.55 ns/op            0 B/op          0 allocs/op
BenchmarkRange/100/builtin-4                      986728              1263 ns/op               0 B/op          0 allocs/op
BenchmarkRange/100/sync-4                         898848              1332 ns/op               0 B/op          0 allocs/op
BenchmarkRange/100/ordered-4                     4099268               309.9 ns/op             0 B/op          0 allocs/op
BenchmarkRange/1000/builtin-4                      85122             14102 ns/op               0 B/op          0 allocs/op
BenchmarkRange/1000/sync-4                         81168             14773 ns/op               0 B/op          0 allocs/op
BenchmarkRange/1000/ordered-4                     377306              3009 ns/op               0 B/op          0 allocs/op
PASS
ok      go.yhsif.com/orderedmap 51.768s
```

As you can see, all operations except `Range` are on-par or only slightly slower
than `sync.Map`, while `Range` is a lot faster.

## License

[BSD License](LICENSE).
