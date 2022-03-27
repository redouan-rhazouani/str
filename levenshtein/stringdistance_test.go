package levenshtein

import (
	"testing"
)

var tests = []struct {
	first  string
	second string
	want   int
}{
	{"", "", 0},
	{"a", "", 1},
	{"", "a", 1},
	{"abc", "abc", 0},
	{"василий", "василий", 0},
	//insertion only
	{"a", "ab", 1},
	{"b", "ba", 1},
	{"ac", "abc", 1},
	{"abcdef", "xabxcdxxefx", 5},
	//deletion only
	{"ab", "a", 1},
	{"ab", "b", 1},
	{"abc", "ac", 1},
	{"xabxcdxxefx", "abcdef", 5},
	//substitution only
	{"a", "b", 1},
	{"ab", "ac", 1},
	{"ac", "bc", 1},
	{"abc", "axc", 1},
	{"xabxcdxxefx", "1ab2cd34ef5", 5},
	//combinaitons
	{"levenshtein", "frankenstein", 6},
	{"Haus", "Maus", 1},
	{"Haus", "Mausi", 2},
	{"Haus", "Häuser", 3},
	{"Kartoffelsalat", "Runkelrüben", 12},
	{"kitten", "sitting", 3},
	{"flaw", "lawn", 2},
	{"Tier", "Tor", 2},
}

func TestDistanceTowRows(t *testing.T) {
	method := Levenshtein{Limit: -1}
	for _, c := range tests {
		got := method.Distance(c.first, c.second)
		if got != c.want {
			t.Errorf("Lenenshtein(%s, %s) = %d want %d", c.first, c.second, got, c.want)
		}
	}
}

func TestDistanceLimitALL(t *testing.T) {
	for _, c := range tests {
		method := Levenshtein{Limit: c.want}
		got := method.Distance(c.first, c.second)
		if got != c.want {
			t.Errorf("Lenenshtein(%s, %s) = %d want %d", c.first, c.second, got, c.want)
		}
	}
}

func TestDistanceLimitNone(t *testing.T) {
	method := Levenshtein{Limit: 0}
	for _, c := range tests {
		got := method.Distance(c.first, c.second)
		if got != 1 && c.want > 0 {
			t.Errorf("Lenenshtein(%s, %s) = %d want %d", c.first, c.second, got, c.want)
		}
	}
}

func TestDistanceLimitMiddle(t *testing.T) {
	for _, c := range tests {
		middle := c.want / 2
		if middle < 1 {
			continue
		}
		method := Levenshtein{Limit: middle}

		got := method.Distance(c.first, c.second)
		if got != middle+1 && c.want > 0 {
			t.Errorf("Lenenshtein(%s, %s) = %d want %d", c.first, c.second, got, c.want)
		}
	}
}
