package day04

import (
	"aoc2015/utils"
	"testing"
)

func TestDay04_1(t *testing.T) {
	cases := []utils.TestCase{
		utils.NewTestCase("abcdef", 609043),
		utils.NewTestCase("pqrstuv", 1048970),
	}

	for _, c := range cases {
		got, err := d04_p1(c.Input)
		if err != nil {
			t.Error(err)
		}
		if got != c.Expected {
			t.Errorf("ERROR %s should result in %d, but resulted in %d instead", c.Input, c.Expected, got)
		}
	}
}
