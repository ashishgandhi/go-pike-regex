package pike

import "testing"

var testCases = []struct {
	regexStr, textStr string
	result            bool
	err               error
}{
	{regexStr: ".*", textStr: "", result: true},
	{regexStr: ".*", textStr: "foo", result: true},
	{regexStr: "^", textStr: "", result: true},
	{regexStr: "^", textStr: "foo", result: true},
	{regexStr: "", textStr: "", result: true},
	{regexStr: "", textStr: "foo", result: true},
	{regexStr: ".**", textStr: "", err: invalidRegexErr},
	{regexStr: "*oo", textStr: "foo", err: invalidRegexErr},
	{regexStr: "foo$bar", textStr: "foo", err: invalidRegexErr},
	{regexStr: "foo*$bar", textStr: "foo", err: invalidRegexErr},
	{regexStr: "foo", textStr: "bar"},
	{regexStr: "foo", textStr: "foo", result: true},
	{regexStr: "foo", textStr: "foobar", result: true},
	{regexStr: "foo", textStr: "barfoobaz", result: true},
	{regexStr: "foo", textStr: "fofafoof", result: true},
	{regexStr: "f.o", textStr: "foo", result: true},
	{regexStr: "f.o", textStr: "fxo", result: true},
	{regexStr: "f.o", textStr: "fxyo"},
	{regexStr: "f*oo", textStr: "bar"},
	{regexStr: "f**oo", textStr: "bar", err: invalidRegexErr},
	{regexStr: "f*oo", textStr: "foo", result: true},
	{regexStr: "f*oo", textStr: "ffffoo", result: true},
	{regexStr: "f*oo", textStr: "oo", result: true},
	{regexStr: "^f.*o", textStr: "bar"},
	{regexStr: "^f.*o", textStr: "fo", result: true},
	{regexStr: "^f.*o", textStr: "fπo", result: true},
	{regexStr: "^f.*o", textStr: "fππππobarbaz", result: true},
	{regexStr: "^f.*o", textStr: "ππππobarbaz"},
	{regexStr: "^f.*o", textStr: "fππππo", result: true},
	{regexStr: "f.*o$", textStr: "fππππo", result: true},
	{regexStr: "f.*o$", textStr: "ππππobarbaz"},
	{regexStr: "f.*o$", textStr: "barbazfππππobarbaz"},
	{regexStr: "aa*ab", textStr: "foobarab"},
	{regexStr: "aa*ab", textStr: "foobaraaaabbaz", result: true},
	{regexStr: "a*b", textStr: "b", result: true},
	{regexStr: "aa*abb*ba", textStr: "aabba", result: true},
	{regexStr: "aa*abb*ba", textStr: "faaaabbbbaoo", result: true},
	{regexStr: "aa*abb*ba.*", textStr: "aabba", result: true},
}

func TestMatch(t *testing.T) {
	for _, tt := range testCases {
		result, err := Match(tt.regexStr, tt.textStr)
		if err != tt.err {
			t.Errorf("Match error incorrect for %s on %s", tt.regexStr, tt.textStr)
		}
		if result != tt.result {
			t.Errorf("Match result incorrect for %s on %s", tt.regexStr, tt.textStr)
		}
	}
}
