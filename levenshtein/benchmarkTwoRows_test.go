package levenshtein

import "testing"

var with2rows = Levenshtein{Limit: -1}

func Benchmark2kRows8(b *testing.B) {
	benchmark(b, 3, &with2rows)
}

func Benchmark2kRows16(b *testing.B) {
	benchmark(b, 4, &with2rows)
}
func Benchmark2kRows32(b *testing.B) {
	benchmark(b, 5, &with2rows)
}
func Benchmark2kRows64(b *testing.B) {
	benchmark(b, 6, &with2rows)
}
func Benchmark2kRows128(b *testing.B) {
	benchmark(b, 7, &with2rows)
}

func Benchmark2kRows256(b *testing.B) {
	benchmark(b, 8, &with2rows)
}

func Benchmark2kRows512(b *testing.B) {
	benchmark(b, 9, &with2rows)
}

func Benchmark2kRows1024(b *testing.B) {
	benchmark(b, 10, &with2rows)
}

func Benchmark2Rows4096(b *testing.B) {
	benchmark(b, 12, &with2rows)
}

func Benchmark2kRows8192(b *testing.B) {
	benchmark(b, 13, &with2rows)
}

func Benchmark2kRows16384(b *testing.B) {
	benchmark(b, 14, &with2rows)
}
