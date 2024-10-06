package wslice

import (
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{input: []int{1, 2, 3, 4, 5}, expected: []int{5, 4, 3, 2, 1}},
		{input: []int{1, 2, 3, 4}, expected: []int{4, 3, 2, 1}},
		{input: []int{1}, expected: []int{1}},
		{input: []int{}, expected: []int{}},
	}

	for _, test := range tests {
		Reverse(test.input)
		if !reflect.DeepEqual(test.input, test.expected) {
			t.Errorf("Reverse() = %v, want %v", test.input, test.expected)
		}
	}
}

func TestContains(t *testing.T) {
	tests := []struct {
		slice    []int
		element  int
		expected bool
	}{
		{slice: []int{1, 2, 3, 4, 5}, element: 3, expected: true},
		{slice: []int{1, 2, 3, 4, 5}, element: 6, expected: false},
		{slice: []int{}, element: 1, expected: false},
	}

	for _, test := range tests {
		result := Contains(test.slice, test.element)
		if result != test.expected {
			t.Errorf("Contains() = %v, want %v", result, test.expected)
		}
	}
}

func TestSort(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{input: []int{5, 3, 4, 1, 2}, expected: []int{1, 2, 3, 4, 5}},
		{input: []int{4, 3, 2, 1}, expected: []int{1, 2, 3, 4}},
		{input: []int{1}, expected: []int{1}},
		{input: []int{}, expected: []int{}},
	}

	for _, test := range tests {
		Sort(test.input)
		if !reflect.DeepEqual(test.input, test.expected) {
			t.Errorf("HeapSort() = %v, want %v", test.input, test.expected)
		}
	}
}
