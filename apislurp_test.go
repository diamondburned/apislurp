package apislurp

import "testing"

func TestEnglishNumber(t *testing.T) {
	type numberTest struct {
		words  string
		number int
		word   string
	}

	var tests = []numberTest{
		{"one sheep", 1, "sheep"},
		{"two sheeps", 2, "sheeps"},
		{"three cows", 3, "cows"},
		{"word", 0, "word"},
		{"not a number", 0, "not a number"},
	}

	for _, test := range tests {
		n, w := englishNumber(test.words)
		if n != test.number {
			t.Errorf("wrong number, input %q, expected %d, got %d", test.words, test.number, n)
		}
		if w != test.word {
			t.Errorf("wrong word, input %q, expected %q, got %q", test.words, test.word, w)
		}
	}
}

func TestTypeNameContains(t *testing.T) {
	type containsTest struct {
		names    []TypeName
		contains TypeName
		found    bool
	}

	var tests = []containsTest{
		{[]TypeName{TypeString, TypeInteger}, TypeInteger, true},
		{[]TypeName{TypeString, TypeInteger}, TypeUser, false},
	}

	for i, test := range tests {
		c := typeNameContains(test.names, test.contains)
		if c != test.found {
			t.Errorf("unexpected results %d: expected %t, got %t", i, test.found, c)
		}
	}
}
