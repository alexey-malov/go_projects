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
		{"аз есмь строка живу я мерой остр", "за семь морей ростка я вижу рост", true},
		{"лесопромышленность", "солепромышленность", true},
		{"старорежимность", "нерасторжимость", true},
		{"австралопитек", "ватерполистка", true},
		{"просветитель", "терпеливость", true},
		{"покраснение", "пенсионерка", true},
		{"равновесие", "своенравие", true},
		{"полковник", "клоповник", true},
		{"стационар", "соратница", true},
		{"вертикаль", "кильватер", true},
		{"апельсин", "спаниель", true},
		{"внимание", "вениамин", true},
	}
	for _, c := range cases {
		got := IsAnagram(c.first, c.second)
		if got != c.isAnagram {
			t.Errorf("IsAnagram('%s', '%s')=%v, want: %v", c.first, c.second, got, c.isAnagram)
		}
		got = IsAnagram2(c.first, c.second)
		if got != c.isAnagram {
			t.Errorf("IsAnagram2('%s', '%s')=%v, want: %v", c.first, c.second, got, c.isAnagram)
		}
	}
}

type AnagramArgs struct {
	first, second string
}

var benchmarkData = []AnagramArgs{
	{"", ""},
	{"с новым годом", "говно с дымом"},
	{"снегурочка", "негросучка"},
	{"сия сцена", "сенсация"},
	{"собака", "кошка"},
	{"кошка", "кошак"},
	{"аз есмь строка живу я мерой остр", "за семь морей ростка я вижу рост"},
	{"Я в мире — сирота", "Я в риме — ариост"},
	{"лесопромышленность", "солепромышленность"},
	{"старорежимность", "нерасторжимость"},
	{"австралопитек", "ватерполистка"},
	{"просветитель", "терпеливость"},
	{"покраснение", "пенсионерка"},
	{"равновесие", "своенравие"},
	{"полковник", "клоповник"},
	{"стационар", "соратница"},
	{"вертикаль", "кильватер"},
	{"апельсин", "спаниель"},
	{"внимание", "вениамин"},
}

func BenchmarkIsAnagram(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, sample := range benchmarkData {
			IsAnagram(sample.first, sample.second)
		}
	}
}

func BenchmarkIsAnagram2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, sample := range benchmarkData {
			IsAnagram2(sample.first, sample.second)
		}
	}
}
