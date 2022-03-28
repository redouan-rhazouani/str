package quiz

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

// reverseSentence creates a new string by reversing the words order in the input string
func reverseSentence(s string) string {
	xs := strings.Fields(s)
	if len(xs) < 2 {
		return s
	}
	//b := strings.Builder
	//ret :=sort.Reverse((sort.StringSlice(xs)))
	//fmt.Println(ret)
	//return strings.Join(xs, " ")
	//s = xs[len(xs) - 1]
	b := strings.Builder{}
	b.Grow(len(s))
	for j := len(xs) - 1; j >= 0; j-- {
		b.WriteString(xs[j])
		if j > 0 {
			b.WriteString(" ")
		}
	}
	return b.String()
}


type span struct {
	start int
	end   int
}

func WordPostions(s string, ispace func(rune) bool) []span {
	spans := make([]span, 0, len(s)/8)
	start := -1
	for end, x := range s {
		if ispace(x) {
			if start < 0 {
				continue
			}
			spans = append(spans, span{start, end})
			start = -1
		} else {
			if start < 0 {
				start = end
			}
		}
	}

	if start >= 0 {
		spans = append(spans, span{start, len(s)})
	}
	return spans
}

//Mirror string s
func Mirror(s string) string {
	spans := WordPostions(s, unicode.IsSpace)
	b := strings.Builder{}
	b.Grow(len(s))
	if len(spans) < 1 {
		r := s
		for len(r) > 0 {
			rune, size := utf8.DecodeLastRuneInString(r)
			b.WriteRune(rune)
			r = r[:len(r)-size]
		}
		return b.String()
	}

	if last := len(spans) - 1; spans[last].end < len(s) {
		r := s[spans[last].end:]
		for len(r) > 0 {
			rune, size := utf8.DecodeLastRuneInString(r)
			b.WriteRune(rune)
			r = r[:len(r)-size]
		}
		//b.WriteString(s[spans[last].end:])
	}
	for n := len(spans) - 1; n >= 0; n-- {
		b.WriteString(s[spans[n].start:spans[n].end])
		left := 0
		if n-1 >= 0 {
			left = spans[n-1].end
		}
		//b.WriteString(s[left:spans[n].start])
		r := s[left:spans[n].start]
		for len(r) > 0 {
			rune, size := utf8.DecodeLastRuneInString(r)
			b.WriteRune(rune)
			r = r[:len(r)-size]
		}

	}
	return b.String()
}

func reverseUTF8(s string) string {
	runes := []rune(s)
	for i, n := 0, len(runes); i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}
	//fmt.Println(cap(runes))
	return string(runes)
}

func reverseUTF8c(input string) string {
	// Get Unicode code points.
	n := 0
	rune := make([]rune, len(input))
	for _, r := range input {
		rune[n] = r
		n++
	}
	rune = rune[0:n]

	// Reverse
	for i := 0; i < n/2; i++ {
		rune[i], rune[n-1-i] = rune[n-1-i], rune[i]
	}
	// Convert back to UTF-8.
	return string(rune)
}

func reverseUTF8b(s string) string {
	b := strings.Builder{}
	n := len(s)
	b.Grow(n)
	for n > 0 {
		rune, size := utf8.DecodeLastRuneInString(s)
		b.WriteRune(rune)
		n -= size
		s = s[:n]
	}
	return b.String()
}

//https://gist.github.com/acsellers/26f8f9cfc0cf5ed8353cc1eab4c07219
var combining = &unicode.RangeTable{
	R16: []unicode.Range16{
		{0x0300, 0x036f, 1}, // combining diacritical marks
		{0x1ab0, 0x1aff, 1}, // combining diacritical marks extended
		{0x1dc0, 0x1dff, 1}, // combining diacritical marks supplement
		{0x20d0, 0x20ff, 1}, // combining diacritical marks for symbols
		{0xfe20, 0xfe2f, 1}, // combining half marks
	},
}

func reverseUTF8b2(s string) string {
	b := strings.Builder{}
	n := len(s)
	b.Grow(n)
	cc := make([]rune, 0, 4)
	for n > 0 {
		rune, size := utf8.DecodeLastRuneInString(s)
		if unicode.In(rune, combining) {
			cc = append(cc, rune)
		} else {
			b.WriteRune(rune)
			for j := len(cc) - 1; j >= 0; j-- {
				b.WriteRune(cc[j])
			}
			cc = cc[:0]

		}
		n -= size
		s = s[:n]
	}
	return b.String()
}

func isCombiningChar(rune rune) bool {
	return rune >= 0x0300 && rune <= 0x036f ||
		rune >= 0x1ab0 && rune <= 0x1aff ||
		rune >= 0x1dc0 && rune <= 0x1dff ||
		rune >= 0x20d0 && rune <= 0x20ff ||
		rune >= 0xfe20 && rune <= 0xfe2f ||
		rune == 0x3099 || rune == 0x309A
}

func reverseUTF8b3(s string) string {
	b := strings.Builder{}
	n := len(s)
	b.Grow(n)
	cc := make([]rune, 0, 4)
	for n > 0 {
		rune, size := utf8.DecodeLastRuneInString(s)
		if isCombiningChar(rune) {
			cc = append(cc, rune)
		} else {
			b.WriteRune(rune)
			for j := len(cc) - 1; j >= 0; j-- {
				b.WriteRune(cc[j])
			}
			cc = cc[:0]

		}
		n -= size
		s = s[:n]
	}
	return b.String()
}

func reverseUTF8Com(s string) string {
	sv := []rune(s)
	rv := make([]rune, 0, len(sv))
	cv := make([]rune, 0)
	for ix := len(sv) - 1; ix >= 0; ix-- {
		r := sv[ix]
		if unicode.In(r, combining) {
			cv = append(cv, r)
		} else {
			rv = append(rv, r)
			for j := len(cv) - 1; j >= 0; j-- {
				rv = append(rv, cv[j])
			}
			cv = cv[:0]
		}
	}
	return string(rv)
}
