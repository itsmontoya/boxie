package boxxy

import "testing"

var testVal interface{}

func TestBasic(t *testing.T) {
	b := New()
	for i := 0; i < (1024 * 1024); i++ {
		b.Append(i)
	}

	b.ForEach(func(i int, v interface{}) bool {
		if i != v.(int) {
			t.Fatalf("invalid value\nExpected: %v\nReturned: %v\n", i, v)
		}

		return false
	})

	b.Insert(51, "w00t")
	if str := b.Get(51).(string); str != "w00t" {
		t.Fatalf("invalid value\nExpected: %v\nReturned: %v\n", "w00t", str)
	}

	b.Prepend("beginning")
	if str := b.Get(0).(string); str != "beginning" {
		t.Fatalf("invalid value\nExpected: %v\nReturned: %v\n", "beginning", str)
	}
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
	bx := populatedBoxxy(100000)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		bx.ForEach(func(_ int, v interface{}) (end bool) {
			testVal = v
			return
		})
	}
	b.ReportAllocs()
}

func BenchmarkSliceForEach(b *testing.B) {
	s := populatedSlice(100000)
	b.ResetTimer()

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

func populatedSlice(n int) (s []interface{}) {
	s = make([]interface{}, n)

	for i := 0; i < n; i++ {
		s[i] = i
	}

	return
}
