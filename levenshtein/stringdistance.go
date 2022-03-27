package levenshtein

import "math"

// //Interface interface for calculating string distance
// type Interface interface {
// 	Distance(first string, second string) int
// }

// Levenshtein distance is a string metric for measuring the difference between two sequences
// It returns the distance if distance doesn't exceed limit and limit otherwise.
// Informally, the Levenshtein distance between two words is the minimum number of single-character edits (insertions, deletions or substitutions) required to change one word into the other
// see https://en.wikipedia.org/wiki/Levenshtein_distance
type Levenshtein struct {
	Limit int //ignored if negative
}

//Distance calculates Levenshtein distance between two sequences
func (l *Levenshtein) Distance(first, second string) int {
	if l.Limit < 0 {
		return distanceR([]rune(first), []rune(second))
	}
	return distanceL([]rune(first), []rune(second), l.Limit)
}

//distanceL calculates Levenshtein distance between two sequences
// It returns limit+1 if distance > limit
func distanceL(s []rune, t []rune, limit int) int {
	m, n := len(s)+1, len(t)+1
	if n-m > limit || m-n > limit { //lower bound
		return limit + 1
	}
	prevd := make([]int, n, n)
	curd := make([]int, n, n)
	for j := 0; j < n; j++ {
		prevd[j] = j
	}

	for i := 1; i < m; i++ {
		curd[0] = i
		rowmin := math.MaxInt32
		for j := 1; j < n; j++ {
			cost := 0
			if s[i-1] != t[j-1] {
				cost = 1
			}
			curd[j] = min(
				prevd[j-1]+cost, // substitution
				prevd[j]+1,      // deletion
				curd[j-1]+1)     // insertion
			if curd[j] < rowmin {
				rowmin = curd[j]
			}
		}
		if rowmin > limit && n > 1 {
			return limit + 1
		}
		prevd, curd = curd, prevd
	}
	if prevd[n-1] > limit {
		return limit + 1
	}
	return prevd[n-1]
}

//1 2 3 4 5 6
//a b c d e f
//
//distanceR calculates Levenshtein distance between two sequences.
func distanceR(s []rune, t []rune) int {
	m, n := len(s)+1, len(t)+1
	prevd := make([]int, n, n)
	curd := make([]int, n, n)
	for j := 0; j < n; j++ {
		prevd[j] = j
	}

	for i := 1; i < m; i++ {
		curd[0] = i
		for j := 1; j < n; j++ {
			cost := 0
			if s[i-1] != t[j-1] {
				cost = 1
			}
			curd[j] = min(
				prevd[j-1]+cost, // substitution
				prevd[j]+1,      // deletion
				curd[j-1]+1)     // insertion
			//	log.Printf("(%d, %d) = %d\t", i, j, dm[i][j])
		}
		prevd, curd = curd, prevd
	}

	return prevd[n-1]
}

func min(a, b, c int) int {
	if a <= b && a <= c {
		return a
	}
	if b < c {
		return b
	}
	return c
}
