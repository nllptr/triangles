package triangles

import (
	"testing"
)

var validCases = []struct {
	input    []int
	expected bool
}{
	{[]int{1, 1, 1}, true},  // equilateral
	{[]int{4, 4, 3}, true},  // isosceles
	{[]int{4, 3, 5}, true},  // scalene
	{[]int{3, 3, 7}, false}, // invalid
	{[]int{3, 3, 6}, false}, // degenerate
}

func TestValid(t *testing.T) {
	for i, c := range validCases {
		got := New(c.input[0], c.input[1], c.input[2]).Valid()
		if got != c.expected {
			t.Errorf("Case %d: Expected %t, got %t", i, c.expected, got)
		}
	}
}

var typeCases = []struct {
	input    []int
	expected TriangleType
}{
	{[]int{1, 1, 1}, EQUILATERAL}, // equilateral
	{[]int{4, 4, 3}, ISOSCELES},   // isosceles
	{[]int{4, 3, 5}, SCALENE},     // scalene
	{[]int{3, 3, 7}, INVALID},     // invalid
	{[]int{3, 3, 6}, INVALID},     // degenerate
}

func TestType(t *testing.T) {
	for i, c := range typeCases {
		got := New(c.input[0], c.input[1], c.input[2]).Type()
		if got != c.expected {
			t.Errorf("Case %d: Expected %v, got %v", i, c.expected, got)
		}
	}
}
