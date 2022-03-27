package levenshtein

import (
	"math"
	"testing"
)

var withlimit = Levenshtein{Limit: math.MaxInt32}

func BenchmarkLimit8(b *testing.B) {
	benchmark(b, 3, &withlimit)
}

func BenchmarkLimit16(b *testing.B) {
	benchmark(b, 4, &withlimit)
}
func BenchmarkLimit32(b *testing.B) {
	benchmark(b, 5, &withlimit)
}
func BenchmarkLimit64(b *testing.B) {
	benchmark(b, 6, &withlimit)
}
func BenchmarkLimit128(b *testing.B) {
	benchmark(b, 7, &withlimit)
}

func BenchmarkLimit256(b *testing.B) {
	benchmark(b, 8, &withlimit)
}

func BenchmarkLimit512(b *testing.B) {
	benchmark(b, 9, &withlimit)
}

func BenchmarkLimit1024(b *testing.B) {
	benchmark(b, 10, &withlimit)
}

func BenchmarkLimit4096(b *testing.B) {
	benchmark(b, 12, &withlimit)
}

func BenchmarkLimit8192(b *testing.B) {
	benchmark(b, 13, &withlimit)
}

func BenchmarkLimit16384(b *testing.B) {
	benchmark(b, 14, &withlimit)
}
