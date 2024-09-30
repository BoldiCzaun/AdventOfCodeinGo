package day05

import (
	"aoc2015/utils"
	"strconv"
	"testing"
)

func TestDay05_1(t *testing.T) {
	cases := []utils.TestCase{
		utils.NewTestCase("ugknbfddgicrmopn", true),
		utils.NewTestCase("aaa", true),
		utils.NewTestCase("jchzalrnumimnmhp", false),
		utils.NewTestCase("haegwjzuvuyypxyu", false),
		utils.NewTestCase("dvszwmarrgswjxmb", false),
	}

	for _, c := range cases {
		got := d05_p1(c.Input)
		if got != c.Expected {
			t.Errorf("ERROR %s should result in %s, but resulted in %s instead\n", c.Input, strconv.FormatBool(c.Expected.(bool)), strconv.FormatBool(got))
		} else {
			t.Logf("SUCCESS %s resulted in %s correctly\n", c.Input, strconv.FormatBool(got))
		}
	}
}

func TestDay05_2(t *testing.T) {
	cases := []utils.TestCase{
		utils.NewTestCase("qjhvhtzxzqqjkmpb", true),
		utils.NewTestCase("xxyxx", true),
		utils.NewTestCase("uurcxstgmygtbstg", false),
		utils.NewTestCase("ieodomkazucvgmuy", false),
	}

	for _, c := range cases {
		got := d05_p2(c.Input)
		if got != c.Expected {
			t.Errorf("ERROR %s should result in %s, but resulted in %s instead\n", c.Input, strconv.FormatBool(c.Expected.(bool)), strconv.FormatBool(got))
		} else {
			t.Logf("SUCCESS %s resulted in %s correctly\n", c.Input, strconv.FormatBool(got))
		}
	}
}
