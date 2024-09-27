package day01

import (
	utils "aoc2015/utils"
	"testing"
)

func TestDay01_1(t *testing.T) {
	cases := []utils.TestCase{
		utils.NewTestCase("(())", 0),
		utils.NewTestCase("()()", 0),
		utils.NewTestCase("(((", 3),
		utils.NewTestCase("(()(()(", 3),
		utils.NewTestCase("))(((((", 3),
		utils.NewTestCase("())", -1),
		utils.NewTestCase("))(", -1),
		utils.NewTestCase(")))", -3),
		utils.NewTestCase(")())())", -3),
	}

	for _, c := range cases {
		got := d01_p1(c.Input)
		if got != c.Expected {
			t.Errorf("ERROR %s should result in %d, but resulted in %d instead", c.Input, c.Expected, got)
		}
	}
}

func TestDay01_2(t *testing.T) {
	cases := []utils.TestCase{
		utils.NewTestCase(")", 1),
		utils.NewTestCase("()())", 5),
	}

	for _, c := range cases {
		got, err := d01_p2(c.Input)
		if err != nil {
			t.Error(err)
		}
		if got != c.Expected {
			t.Errorf("ERROR %s should result in %d, but resulted in %d instead", c.Input, c.Expected, got)
		}
	}
}
