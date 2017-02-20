# Boxxy [![GoDoc](https://godoc.org/github.com/itsmontoya/boxxy?status.svg)](https://godoc.org/github.com/itsmontoya/boxxy) ![Status](https://img.shields.io/badge/status-alpha-red.svg)

Boxxy is a sharded slice solution which offers:
- Appending
- Prepending
- Inserting
- Trimming (not yet implemented)

## Benchmarks
```bash
# Still need to create benchmarks
```

## Related projects
[SegmentedSlice](https://github.com/OneOfOne/segmentedSlice) by @OneOfOne was the inspiration for Boxxy. SegmentedSlice is a fast, index-able, sort-able, grow-only Slice. If your project does not require prepend and insert and/or if your project requires sorting, please give SegmentedSlice a try!

