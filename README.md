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
cpu: 12th Gen Intel(R) Core(TM) i5-1235U
BenchmarkMap/DeleteEmpty/builtin-12             1000000000               1.067 ns/op           0 B/op          0 allocs/op
BenchmarkMap/DeleteEmpty/sync-12                1000000000               1.089 ns/op           0 B/op          0 allocs/op
BenchmarkMap/DeleteEmpty/ordered-12             10550479               107.1 ns/op             0 B/op          0 allocs/op
BenchmarkMap/LoadEmpty/builtin-12               934932499                1.165 ns/op           0 B/op          0 allocs/op
BenchmarkMap/LoadEmpty/sync-12                  1000000000               1.144 ns/op           0 B/op          0 allocs/op
BenchmarkMap/LoadEmpty/ordered-12               20735286                56.56 ns/op            0 B/op          0 allocs/op
BenchmarkMap/Store1/builtin-12                  157484062                7.507 ns/op           0 B/op          0 allocs/op
BenchmarkMap/Store1/sync-12                      4521426               268.2 ns/op            48 B/op          3 allocs/op
BenchmarkMap/Store1/ordered-12                   9733495               122.5 ns/op            32 B/op          1 allocs/op
BenchmarkMap/Store2/builtin-12                  151412253                7.859 ns/op           0 B/op          0 allocs/op
BenchmarkMap/Store2/sync-12                      4417717               264.8 ns/op            48 B/op          3 allocs/op
BenchmarkMap/Store2/ordered-12                  10051618               121.7 ns/op            32 B/op          1 allocs/op
BenchmarkMap/Update1/builtin-12                 161776555                7.772 ns/op           0 B/op          0 allocs/op
BenchmarkMap/Update1/sync-12                     4492196               301.7 ns/op            48 B/op          3 allocs/op
BenchmarkMap/Update1/ordered-12                  7688208               147.1 ns/op            32 B/op          1 allocs/op
BenchmarkMap/Load1/builtin-12                   436247556                2.537 ns/op           0 B/op          0 allocs/op
BenchmarkMap/Load1/sync-12                      425829909                4.240 ns/op           0 B/op          0 allocs/op
BenchmarkMap/Load1/ordered-12                   17000335                64.42 ns/op            0 B/op          0 allocs/op
BenchmarkMap/Load3/builtin-12                   130940556                8.450 ns/op           0 B/op          0 allocs/op
BenchmarkMap/Load3/sync-12                      602772472                3.151 ns/op           0 B/op          0 allocs/op
BenchmarkMap/Load3/ordered-12                   15910669                71.54 ns/op            0 B/op          0 allocs/op
BenchmarkLoadOrStoreStore/builtin-12            67150718                15.16 ns/op            0 B/op          0 allocs/op
BenchmarkLoadOrStoreStore/sync-12                7958462               187.2 ns/op           360 B/op          8 allocs/op
BenchmarkLoadOrStoreStore/ordered-12             6706869               184.9 ns/op           416 B/op          5 allocs/op
BenchmarkLoadOrStoreLoad/builtin-12             470602132                2.529 ns/op           0 B/op          0 allocs/op
BenchmarkLoadOrStoreLoad/sync-12                75306667                21.58 ns/op           32 B/op          2 allocs/op
BenchmarkLoadOrStoreLoad/ordered-12              8332143               136.6 ns/op             0 B/op          0 allocs/op
BenchmarkStoreThenDelete/builtin-12             55998295                21.19 ns/op            0 B/op          0 allocs/op
BenchmarkStoreThenDelete/sync-12                 1961560               670.6 ns/op           346 B/op          8 allocs/op
BenchmarkStoreThenDelete/ordered-12              3997981               293.8 ns/op            78 B/op          1 allocs/op
BenchmarkRange/10/builtin-12                    12832806                85.99 ns/op            0 B/op          0 allocs/op
BenchmarkRange/10/sync-12                       73554632                26.04 ns/op            0 B/op          0 allocs/op
BenchmarkRange/10/ordered-12                     1466373               756.6 ns/op             0 B/op          0 allocs/op
BenchmarkRange/100/builtin-12                    1762689               667.0 ns/op             0 B/op          0 allocs/op
BenchmarkRange/100/sync-12                       7852339               264.1 ns/op             0 B/op          0 allocs/op
BenchmarkRange/100/ordered-12                     152163              6690 ns/op               0 B/op          0 allocs/op
BenchmarkRange/1000/builtin-12                    118522              8893 ns/op               0 B/op          0 allocs/op
BenchmarkRange/1000/sync-12                       682383              2591 ns/op               0 B/op          0 allocs/op
BenchmarkRange/1000/ordered-12                     16834             64454 ns/op               0 B/op          0 allocs/op
PASS
ok      go.yhsif.com/orderedmap 64.694s
```

Note that for the benchmark tests,
`sync` and `ordered` are parallel benchmarks while `builtin` are sequential.

## License

[BSD License](LICENSE).
