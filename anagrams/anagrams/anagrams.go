package anagrams

import (
	"sort"
	"strings"
	"unicode"
	"unicode/utf8"
)

type ByRune []rune

func (a ByRune) Len() int           { return len(a) }
func (a ByRune) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByRune) Less(i, j int) bool { return a[i] < a[j] }

func stripSpaces(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			// if the character is a space, drop it
			return -1
		}
		// else keep it in the string
		return r
	}, str)
}

func runesFromString(s string) []rune {
	n := utf8.RuneCountInString(s)
	runes := make([]rune, n)

	for i := 0; i < n; i++ {
		r, size := utf8.DecodeRuneInString(s)
		runes[i] = r
		s = s[size:]
	}
	return runes
}

func IsAnagram(first, second string) bool {
	runes1 := runesFromString(stripSpaces(first))
	runes2 := []rune(stripSpaces(second))

	if len(runes1) != len(runes2) {
		return false
	}

	sort.Sort(ByRune(runes1))
	sort.Sort(ByRune(runes2))

	for i, r := range runes1 {
		if r != runes2[i] {
			return false
		}
	}
	return true
}
