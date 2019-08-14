package main

import (
	"testing"
)

func TestCalculator(t *testing.T) {
	params := []struct {
		name     string
		expr     string
		expected int
	}{
		{"Example 1", "4 2 5 * + 1 3 2 * + /", 2},
		{"Example 2", "4 7 * 5 2 * *", 280},
		{"Example 3", "90 3 / 30 5 / /", 5},
		{"Example 4", "99 11 + 8 7 + +", 125},
		{"Example 5", "15 7 1 1 + - / 3 * 2 1 1 + + -", 5},
	}

	for _, p := range params {
		t.Run(p.name, func(t *testing.T) {
			res := calc(p.expr)
			if res != p.expected {
				t.Errorf("result is %v, but expected %v", res, p.expected)
			}
		})
	}
}
