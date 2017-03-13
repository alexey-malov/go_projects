package anagrams

import "testing"

func TestIsAnagram(t *testing.T) {
	cases := []struct {
		first, second string
		isAnagram     bool
	}{
		{"", "", true},
		{"a", "a", true},
		{"a", "", false},
		{"", "a", false},
		{"a", "b", false},
		{"ab", "ba", true},
		{"aba", "baa", true},
		{"с новым годом", "говно с дымом", true},
		{"сия сцена", "сенсация", true},
	}
	for _, c := range cases {
		got := IsAnagram(c.first, c.second)
		if got != c.isAnagram {
			t.Errorf("isAnagram('%s', '%s')=%v, want: %v", c.first, c.second, got, c.isAnagram)
		}
	}
}
