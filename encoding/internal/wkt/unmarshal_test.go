package wkt

import "testing"

func TestTrimSpaceBrackets(t *testing.T) {
	cases := []struct {
		s        string
		expected string
	}{
		{
			s:        "(1 2)",
			expected: "1 2",
		},
		{
			s:        "((1 2),(0.5 1.5))",
			expected: "(1 2),(0.5 1.5)",
		},
		{
			s:        "(1 2,0.5 1.5)",
			expected: "1 2,0.5 1.5",
		},
		{
			s:        "((1 2,3 4),(5 6,7 8))",
			expected: "(1 2,3 4),(5 6,7 8)",
		},
		{
			s:        "(((1 2,3 4)),((5 6,7 8),(1 2,5 4)))",
			expected: "((1 2,3 4)),((5 6,7 8),(1 2,5 4))",
		},
	}

	for _, tc := range cases {
		if trimSpaceBrackets(tc.s) != tc.expected {
			t.Log(trimSpaceBrackets(tc.s))
			t.Log(tc.expected)
			t.Errorf("trim space and brackets error")
		}
	}
}
