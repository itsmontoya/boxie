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
	fmt.Println(b.Get(53))

	//	b.ForEach(func(i int, val interface{}) (end bool) {
	//fmt.Println(i, val)
	//		return
	//	})
}

func BenchmarkBoxieGet(b *testing.B) {
	b.StopTimer()
	bx := New()
	for i := 0; i < b.N; i++ {
		bx.Append(i)
	}
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		testVal = bx.Get(i)
	}
	b.ReportAllocs()
}

func BenchmarkSliceGet(b *testing.B) {
	b.StopTimer()
	var s []interface{}
	for i := 0; i < b.N; i++ {
		s = append(s, i)
	}
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		testVal = s[i]
	}
	b.ReportAllocs()
}

func BenchmarkBoxieForEach(b *testing.B) {
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

func BenchmarkBoxieAppend(b *testing.B) {
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

func BenchmarkBoxiePrepend(b *testing.B) {
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
