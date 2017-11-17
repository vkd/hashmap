package hashmap

import (
	"math/rand"
	"strconv"
	"testing"
)

func BenchmarkHashmap_Set_16(b *testing.B) {
	hm := NewHashMap(16, nil)
	benchmarkHashmapSet(b, hm)
}
func BenchmarkHashmap_Set_64(b *testing.B) {
	hm := NewHashMap(64, nil)
	benchmarkHashmapSet(b, hm)
}
func BenchmarkHashmap_Set_128(b *testing.B) {
	hm := NewHashMap(128, nil)
	benchmarkHashmapSet(b, hm)
}
func BenchmarkHashmap_Set_1024(b *testing.B) {
	hm := NewHashMap(1024, nil)
	benchmarkHashmapSet(b, hm)
}
func BenchmarkHashmap_Set_Native(b *testing.B) {
	hm := NewNativeMap()
	benchmarkHashmapSet(b, hm)
}
func benchmarkHashmapSet(b *testing.B, hm HashMaper) {
	var size = 2000

	var v int
	var err error

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v = rand.Intn(size)
		err = hm.Set(strconv.Itoa(v), v)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkHashmap_Get_16(b *testing.B) {
	hm := NewHashMap(16, nil)
	benchmarkHashmapGet(b, hm)
}
func BenchmarkHashmap_Get_64(b *testing.B) {
	hm := NewHashMap(64, nil)
	benchmarkHashmapGet(b, hm)
}
func BenchmarkHashmap_Get_128(b *testing.B) {
	hm := NewHashMap(128, nil)
	benchmarkHashmapGet(b, hm)
}
func BenchmarkHashmap_Get_1024(b *testing.B) {
	hm := NewHashMap(1024, nil)
	benchmarkHashmapGet(b, hm)
}
func BenchmarkHashmap_Get_Native(b *testing.B) {
	hm := NewNativeMap()
	benchmarkHashmapGet(b, hm)
}
func benchmarkHashmapGet(b *testing.B, hm HashMaper) {
	var size = 2000
	hashmapInitialize(hm, size/2)

	var v int
	var err error

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v = rand.Intn(size)
		_, err = hm.Get(strconv.Itoa(v))
		if err != nil && err != ErrKeyNotFound {
			panic(err)
		}
	}
}

func BenchmarkHashmap_Unset_16(b *testing.B) {
	hm := NewHashMap(16, nil)
	benchmarkHashmapUnset(b, hm)
}
func BenchmarkHashmap_Unset_64(b *testing.B) {
	hm := NewHashMap(64, nil)
	benchmarkHashmapUnset(b, hm)
}
func BenchmarkHashmap_Unset_128(b *testing.B) {
	hm := NewHashMap(128, nil)
	benchmarkHashmapUnset(b, hm)
}
func BenchmarkHashmap_Unset_1024(b *testing.B) {
	hm := NewHashMap(1024, nil)
	benchmarkHashmapUnset(b, hm)
}
func BenchmarkHashmap_Unset_Native(b *testing.B) {
	hm := NewNativeMap()
	benchmarkHashmapUnset(b, hm)
}
func benchmarkHashmapUnset(b *testing.B, hm HashMaper) {
	size := 2000
	hashmapInitialize(hm, size/2)

	var v int
	var err error

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v = rand.Intn(size)
		err = hm.Unset(strconv.Itoa(v))
		if err != nil && err != ErrKeyNotFound {
			panic(err)
		}
		if err == nil {
			b.StopTimer()
			hm.Set(strconv.Itoa(v), v)
			b.StartTimer()
		}
	}
}

func hashmapInitialize(hm HashMaper, size int) {
	var err error
	for i := 0; i < size; i++ {
		err = hm.Set(strconv.Itoa(i), i)
		if err != nil {
			panic(err)
		}
	}
}
