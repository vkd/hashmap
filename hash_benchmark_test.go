package hashmap

import (
	"testing"
)

func BenchmarkHash_16(b *testing.B) {
	benchmarkHash(b, 16)
}
func BenchmarkHash_64(b *testing.B) {
	benchmarkHash(b, 64)
}
func BenchmarkHash_128(b *testing.B) {
	benchmarkHash(b, 128)
}
func BenchmarkHash_1024(b *testing.B) {
	benchmarkHash(b, 1024)
}

func benchmarkHash(b *testing.B, size uint) {
	uniq := make(map[uint]bool, size)

	var s struct{ I int }
	var iface interface{} = &s

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.I = i

		v, err := hash(size, iface)
		v = v % size
		if err != nil {
			panic(err)
		}

		uniq[v] = true
	}

	b.Logf("Len uniqs values: %d (%d)", len(uniq), b.N)
}
