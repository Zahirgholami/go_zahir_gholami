package main

import (
	"testing"
)

func TestRectangleArea(t *testing.T) {
	tests := []struct {
		name     string
		length   int
		width    int
		expected string
	}{
		{"even area", 2, 4, "even rectangle"},
		{"odd area", 3, 5, "odd rectangle"},
		{"zero area", 0, 0, "even rectangle"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := RectangleArea(test.length, test.width)
			if actual != test.expected {
				t.Errorf("RectangleArea(%d, %d) = %s, want %s", test.length, test.width, actual, test.expected)
			}
		})
	}
}

func TestRectanglePerimeter(t *testing.T) {
	tests := []struct {
		name     string
		length   int
		width    int
		expected bool
	}{
		{"perimeter divisible by 5", 5, 5, true},
		{"perimeter not divisible by 5", 3, 4, false},
		{"zero perimeter", 0, 0, false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := RectanglePerimeter(test.length, test.width)
			if actual != test.expected {
				t.Errorf("RectanglePerimeter(%d, %d) = %t, want %t", test.length, test.width, actual, test.expected)
			}
		})
	}
}

func TestSquareArea(t *testing.T) {
	tests := []struct {
		name     string
		s        int
		expected string
	}{
		{"even area", 2, "even square"},
		{"odd area", 3, "odd square"},
		{"zero area", 0, "even square"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := SquareArea(test.s)
			if actual != test.expected {
				t.Errorf("SquareArea(%d) = %s, want %s", test.s, actual, test.expected)
			}
		})
	}
}

func TestSquarePerimeter(t *testing.T) {
	tests := []struct {
		name     string
		s        int
		expected bool
	}{
		{"perimeter greater than or equal to 40", 10, true},
		{"perimeter less than 