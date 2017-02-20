package boxxy

import (
	"strconv"
	"testing"

	"fmt"
)

var testVal interface{}

func TestBasic(t *testing.T) {
	b := New()
	for i := 0; i < 54; i++ {
		b.Append(strconv.Itoa(i))
	}

	b.Insert(51, "w00t")
	b.Insert(2, "scoot")
	b.Prepend("beginning")
	fmt.Println("Get 53", b.Get(53))
	fmt.Println("Get 51", b.Get(51))
	fmt.Println("Get 2", b.Get(0))
	fmt.Println(b.bs[1])

	//	b.ForEach(func(i int, val interface{}) (end bool) {
	//fmt.Println(i, val)
	//		return
	//	})
}

func BenchmarkBoxxyGet_10000(b *testing.B) {
	benchmarkBoxxyGet(b, 10000)
}

func BenchmarkBoxxyGet_100000(b *testing.B) {
	benchmarkBoxxyGet(b, 100000)
}

func BenchmarkBoxxyGet_1000000(b *testing.B) {
	benchmarkBoxxyGet(b, 1000000)
}

func BenchmarkSliceGet_10000(b *testing.B) {
	benchmarkSliceGet(b, 10000)
}

func BenchmarkSliceGet_100000(b *testing.B) {
	benchmarkSliceGet(b, 100000)
}

func BenchmarkSliceGet_1000000(b *testing.B) {
	benchmarkSliceGet(b, 1000000)
}

func BenchmarkBoxxyForEach(b *testing.B) {
	b.StopTimer()
	bx := New()
	for i := 0; i < b.N; i++ {
		bx.Append(i)
	}
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		bx.ForEach(func(i int, v interface{}) (end bool) {
			testVal = v
			return
		})
	}
	b.ReportAllocs()
}

func BenchmarkSliceForEach(b *testing.B) {
	b.StopTimer()
	var s []interface{}
	for i := 0; i < b.N; i++ {
		s = append(s, i)
	}
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		for _, v := range s {
			testVal = v
		}
	}
	b.ReportAllocs()
}

func BenchmarkBoxxyAppend(b *testing.B) {
	bx := New()
	for i := 0; i < b.N; i++ {
		bx.Append(i)
	}

	b.ReportAllocs()
}

func BenchmarkSliceAppend(b *testing.B) {
	var s []interface{}
	for i := 0; i < b.N; i++ {
		s = append(s, i)
	}

	b.ReportAllocs()
}

func BenchmarkBoxxyPrepend(b *testing.B) {
	bx := New()
	for i := 0; i < b.N; i++ {
		bx.Prepend(i)
	}

	b.ReportAllocs()
}

func BenchmarkSlicePrepend(b *testing.B) {
	var s []interface{}
	for i := 0; i < b.N; i++ {
		s = append([]interface{}{i}, s...)
	}

	b.ReportAllocs()
}

func benchmarkBoxxyGet(b *testing.B, n int) {
	b.StopTimer()
	bx := populatedBoxxy(n)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		for j := 0; j < n; j++ {
			testVal = bx.Get(j)
		}
	}

	b.ReportAllocs()
}

func benchmarkSliceGet(b *testing.B, n int) {
	b.StopTimer()
	s := populatedSlice(n)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		for j := 0; j < n; j++ {
			testVal = s[j]
		}
	}

	b.ReportAllocs()
}

func populatedBoxxy(n int) (b *Boxxy) {
	b = New()

	for i := 0; i < n; i++ {
		b.Append(i)
	}

	return
}

func populatedSlice(n int) (s []int) {
	s = make([]int, n)

	for i := 0; i < n; i++ {
		s[i] = i
	}

	return
}
