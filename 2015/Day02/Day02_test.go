package day02

import (
	utils "aoc2015/utils"
	"testing"
)

func TestDay02_1(t *testing.T) {
	cases := []utils.TestCase{
		utils.NewTestCase("2x3x4", 58),
		utils.NewTestCase("1x1x10", 43),
	}

	for _, c := range cases {
		got := d02_p1([]string{c.Input})
		if got != c.Expected {
			t.Errorf("ERROR %s should result in %d, but resulted in %d instead", c.Input, c.Expected, got)
		}
	}
}

func TestDay02_2(t *testing.T) {
	cases := []utils.TestCase{
		utils.NewTestCase("2x3x4", 34),
		utils.NewTestCase("1x1x10", 14),
	}

	for _, c := range cases {
		got := d02_p2([]string{c.Input})
		if got != c.Expected {
			t.Errorf("ERROR %s should result in %d, but resulted in %d instead", c.Input, c.Expected, got)
		}
	}
}
