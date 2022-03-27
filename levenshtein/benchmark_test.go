package levenshtein

import (
	"math/rand"
	"testing"
	"time"
)

func init() {
	generate(16)
}

type pair struct {
	first  string
	second string
}

var data []pair

func randomString(rng *rand.Rand, n int) string {
	runes := make([]rune, n)
	for i := 0; i < n; i++ {
		runes[i] = rune(rng.Intn(0x1000)) // random rune up to '\u0999'
	}
	return string(runes)
}

func generate(n uint) {
	data = make([]pair, 0, n)
	seed := time.Now().UTC().UnixNano()
	rng := rand.New(rand.NewSource(seed))
	for i := uint(0); i < n; i++ {
		m := 1 << i
		data = append(data, pair{randomString(rng, m), randomString(rng, m)})
	}
}

func benchmark(b *testing.B, idx uint, strategy *Levenshtein) {
	s, t := data[idx].first, data[idx].second
	for i := 0; i < b.N; i++ {
		strategy.Distance(s, t)
	}
}
