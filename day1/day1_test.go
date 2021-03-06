package main

import "testing"

func TestCount(t *testing.T) {
	cases := []struct {
		in       string
		expected int
	}{
		{"(", 1},
		{"()", 0},
		{"(((", 3},
		{"(()(()(", 3},
		{"))(((((", 3},
		{"())", -1},
		{"))(", -1},
		{")))", -3},
		{")())())", -3},
	}
	for _, c := range cases {
		result := count(c.in)
		if result != c.expected {
			t.Errorf("Count(%q) == %d, expected %d", c.in, result, c.expected)
		}
	}
}

func TestEntryPoint(t *testing.T) {
	cases := []struct {
		input    string
		expected int
	}{
		{"", 0},
		{")", 1},
		{"()())", 5},
	}
	for _, c := range cases {
		basement := -1
		_, result := findFloor(c.input, 0, basement)
		if result != c.expected {
			t.Errorf("Floor(%d) reached at character %d, not %d", basement, result, c.expected)
		}
	}
}
