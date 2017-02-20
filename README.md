# Boxxy [![GoDoc](https://godoc.org/github.com/itsmontoya/boxxy?status.svg)](https://godoc.org/github.com/itsmontoya/boxxy) ![Status](https://img.shields.io/badge/status-alpha-red.svg)

Boxxy is a sharded slice solution which offers:
- Appending
- Prepending
- Inserting
- Trimming (not yet implemented)

## Benchmarks
```bash
BenchmarkBoxxyGet_10000-4       20000       80182 ns/op          0 B/op          0 allocs/op
BenchmarkBoxxyGet_100000-4       2000      805412 ns/op          0 B/op          0 allocs/op
BenchmarkBoxxyGet_1000000-4       200     8436888 ns/op          0 B/op          0 allocs/op
BenchmarkBoxxyForEach-4         30000      117813 ns/op          0 B/op          0 allocs/op
BenchmarkBoxxyAppend-4       30000000        52.7 ns/op         27 B/op          1 allocs/op
BenchmarkBoxxyPrepend-4       1000000       79217 ns/op         27 B/op          1 allocs/op

BenchmarkSliceGet_10000-4        5000      307127 ns/op      80000 B/op      10000 allocs/op
BenchmarkSliceGet_100000-4        500     3019658 ns/op     800005 B/op     100000 allocs/op
BenchmarkSliceGet_1000000-4        50    32099224 ns/op    8000009 B/op    1000000 allocs/op
BenchmarkSliceForEach-4        200000      243569 ns/op          0 B/op          0 allocs/op
BenchmarkSliceAppend-4        5000000         361 ns/op         92 B/op          1 allocs/op
BenchmarkSlicePrepend-4         30000      382389 ns/op     243917 B/op          3 allocs/op
# Still need to create benchmarks
```

## Related projects
[SegmentedSlice](https://github.com/OneOfOne/segmentedSlice) by @OneOfOne was the inspiration for Boxxy. SegmentedSlice is a fast, index-able, sort-able, grow-only Slice. If your project does not require prepend and insert and/or if your project requires sorting, please give SegmentedSlice a try!

