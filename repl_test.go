package pokedex

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input string
		expected []string
	}{
		{
			input: " hello world ",
			expected: []string{"hello", "world"},
		}, 
		{
			input: " To the moOn alice ",
			expected: []string{"to", "the", "moon", "alice"},
		},
		{
			input: "   ",
			expected: []string{},
		},
		{
			input: "",
			expected: []string{},
		},
		{
			input: "a b\tc\n",
			expected: []string{"a", "b", "c"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("cleanInput(%q)=%v; expected %v (len %d vs %d)",
			c.input, 
			actual, 
			c.expected, 
			len(actual), 
			len(c.expected))
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("cleanInput(%q)[%d]=%q; expected %q (full=%v vs %v)", 
				c.input, 
				i, 
				word, 
				expectedWord, 
				actual, 
				c.expected)
			}
		}
	}
}