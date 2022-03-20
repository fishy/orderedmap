[![PkgGoDev](https://pkg.go.dev/badge/go.yhsif.com/orderedmap)](https://pkg.go.dev/go.yhsif.com/orderedmap)
[![Go Report Card](https://goreportcard.com/badge/go.yhsif.com/orderedmap)](https://goreportcard.com/report/go.yhsif.com/orderedmap)

# Go OrderedMap

An implementation of ordered map in Go.
An ordered map preserves the original insertion order when iterating.

It's implemented by wrapping a map and a
[doubly linked list](https://pkg.go.dev/container/list#List).
The interface is intentionally kept the same with `sync.Map` to be used
interchangeably (except with added generics support).

## Benchmark

This package comes with benchmark test to test against
builtin map and [`sync.Map`](https://pkg.go.dev/sync#Map):

```
$ go test -bench .
goos: linux
goarch: amd64
pkg: go.yhsif.com/orderedmap
cpu: Intel(R) Core(TM) i5-7260U CPU @ 2.20GHz
BenchmarkMap/DeleteEmpty/builtin-4              577845259                2.074 ns/op           0 B/op          0 allocs/op
BenchmarkMap/DeleteEmpty/sync-4                 202597698                5.933 ns/op           0 B/op          0 allocs/op
BenchmarkMap/DeleteEmpty/ordered-4              20556426                58.21 ns/op            0 B/op          0 allocs/op
BenchmarkMap/LoadEmpty/builtin-4                540843502                2.219 ns/op           0 B/op          0 allocs/op
BenchmarkMap/LoadEmpty/sync-4                   202360122                5.920 ns/op           0 B/op          0 allocs/op
BenchmarkMap/LoadEmpty/ordered-4                26017804                46.28 ns/op            0 B/op          0 allocs/op
BenchmarkMap/Store1/builtin-4                   100000000               10.48 ns/op            0 B/op          0 allocs/op
BenchmarkMap/Store1/sync-4                       6121250               196.4 ns/op            48 B/op          3 allocs/op
BenchmarkMap/Store1/ordered-4                   12191223                97.90 ns/op           32 B/op          1 allocs/op
BenchmarkMap/Store2/builtin-4                   100000000               11.28 ns/op            0 B/op          0 allocs/op
BenchmarkMap/Store2/sync-4                       6063440               197.1 ns/op            48 B/op          3 allocs/op
BenchmarkMap/Store2/ordered-4                   11598100               102.0 ns/op            32 B/op          1 allocs/op
BenchmarkMap/Update1/builtin-4                  100000000               10.46 ns/op            0 B/op          0 allocs/op
BenchmarkMap/Update1/sync-4                      6090291               196.7 ns/op            48 B/op          3 allocs/op
BenchmarkMap/Update1/ordered-4                  12118083                98.08 ns/op           32 B/op          1 allocs/op
BenchmarkMap/Load1/builtin-4                    289113676                4.144 ns/op           0 B/op          0 allocs/op
BenchmarkMap/Load1/sync-4                       100000000               11.25 ns/op            0 B/op          0 allocs/op
BenchmarkMap/Load1/ordered-4                    26038620                46.19 ns/op            0 B/op          0 allocs/op
BenchmarkMap/Load3/builtin-4                    90047526                12.65 ns/op            0 B/op          0 allocs/op
BenchmarkMap/Load3/sync-4                       149176874                8.037 ns/op           0 B/op          0 allocs/op
BenchmarkMap/Load3/ordered-4                    26861779                44.67 ns/op            0 B/op          0 allocs/op
BenchmarkLoadOrStoreStore/builtin-4             49824517                20.92 ns/op            0 B/op          0 allocs/op
BenchmarkLoadOrStoreStore/sync-4                 4937766               244.9 ns/op           376 B/op          8 allocs/op
BenchmarkLoadOrStoreStore/ordered-4              6684468               177.8 ns/op           416 B/op          5 allocs/op
BenchmarkLoadOrStoreLoad/builtin-4              279954810                4.291 ns/op           0 B/op          0 allocs/op
BenchmarkLoadOrStoreLoad/sync-4                 23575126                51.09 ns/op           32 B/op          2 allocs/op
BenchmarkLoadOrStoreLoad/ordered-4              19255971                62.83 ns/op            0 B/op          0 allocs/op
BenchmarkStoreThenDelete/builtin-4              40588748                28.97 ns/op            0 B/op          0 allocs/op
BenchmarkStoreThenDelete/sync-4                  1878607               633.8 ns/op           357 B/op          8 allocs/op
BenchmarkStoreThenDelete/ordered-4               4875866               244.9 ns/op            79 B/op          1 allocs/op
BenchmarkRange/10/builtin-4                      8666966               138.3 ns/op             0 B/op          0 allocs/op
BenchmarkRange/10/sync-4                        18859324                62.89 ns/op            0 B/op          0 allocs/op
BenchmarkRange/10/ordered-4                      2830831               423.8 ns/op             0 B/op          0 allocs/op
BenchmarkRange/100/builtin-4                      934752              1276 ns/op               0 B/op          0 allocs/op
BenchmarkRange/100/sync-4                        2148686               560.0 ns/op             0 B/op          0 allocs/op
BenchmarkRange/100/ordered-4                      328244              3674 ns/op               0 B/op          0 allocs/op
BenchmarkRange/1000/builtin-4                      83199             14432 ns/op               0 B/op          0 allocs/op
BenchmarkRange/1000/sync-4                        183529              6375 ns/op               0 B/op          0 allocs/op
BenchmarkRange/1000/ordered-4                      33486             35836 ns/op               0 B/op          0 allocs/op
PASS
ok      go.yhsif.com/orderedmap 53.852s
```

Note that for the benchmark tests,
`sync` and `ordered` are parallel benchmarks while `builtin` are sequential.

## License

[BSD License](LICENSE).
