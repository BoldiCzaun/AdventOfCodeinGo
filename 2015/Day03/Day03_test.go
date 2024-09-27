package day03

import (
	utils "aoc2015/utils"
	"testing"
)

func TestDay03_1(t *testing.T) {
	cases := []utils.TestCase{
		utils.NewTestCase(">", 2),
		utils.NewTestCase("^>v<", 4),
		utils.NewTestCase("^v^v^v^v^v", 2),
	}

	for _, c := range cases {
		got := d03_p1(c.Input)
		if got != c.Expected {
			t.Errorf("ERROR %s should result in %d, but resulted in %d instead", c.Input, c.Expected, got)
		}
	}
}

func TestDay03_2(t *testing.T) {
	cases := []utils.TestCase{
		utils.NewTestCase("^v", 3),
		utils.NewTestCase("^>v<", 3),
		utils.NewTestCase("^v^v^v^v^v", 11),
	}

	for _, c := range cases {
		got := d03_p2(c.Input)
		if got != c.Expected {
			t.Errorf("ERROR %s should result in %d, but resulted in %d instead", c.Input, c.Expected, got)
		}
	}
}
