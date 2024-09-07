package main

import "testing"

func TestDay01_1(t *testing.T) {
	cases := []TestCase{
		NewTestCase("(())", 0),
		NewTestCase("()()", 0),
		NewTestCase("(((", 3),
		NewTestCase("(()(()(", 3),
		NewTestCase("))(((((", 3),
		NewTestCase("())", -1),
		NewTestCase("))(", -1),
		NewTestCase(")))", -3),
		NewTestCase(")())())", -3),
	}

	for _, c := range cases {
		got := D01_p1(c.input)
		if got != c.expected {
			t.Logf("ERROR %s should result in %d, but resulted in %d instead", c.input, c.expected, got)
			t.Fail()
		}
	}
}

func TestDay01_2(t *testing.T) {
	cases := []TestCase{
		NewTestCase(")", 1),
		NewTestCase("()())", 5),
	}

	for _, c := range cases {
		got := D01_p2(c.input)
		if got != c.expected {
			t.Logf("ERROR %s should result in %d, but resulted in %d instead", c.input, c.expected, got)
			t.Fail()
		}
	}
}
