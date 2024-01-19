package main

import "testing"

type testCase struct {
	name     string
	input    []int
	expected int
	errMsg   string
}

func ShouldSum(t *testing.T) {
	testCases := []testCase{
		{
			name:     "Sum of 1, 2, 3, 4, 5",
			input:    []int{1, 2, 3, 4, 5},
			expected: 15,
		},
		{
			name:     "Sum of 1, 9",
			input:    []int{1, 9},
			expected: 10,
		},
		{
			name:     "Sum of 10, 20, 30, 40, 50",
			input:    []int{10, 20, 30, 40, 50},
			expected: 150,
		},
		{
			name:     "Sum of 10, -1",
			input:    []int{10, -1},
			expected: 9,
		},
		{
			name:     "Sum of 10",
			input:    []int{10},
			expected: 10,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Sum(tc.input...)
			if result != tc.expected {
				t.Errorf("Expected %d, got %d", tc.expected, result)
			}
		})
	}
}

func ShouldSub(t *testing.T) {
	testCases := []testCase{
		{
			name:     "Sub of 90, 15, 3",
			input:    []int{90, 15, 3},
			expected: 72,
		},
		{
			name:     "Sub of 10, 9",
			input:    []int{10, -9},
			expected: 19,
		},
		{
			name:     "Sub of -10, -30",
			input:    []int{-10, -30},
			expected: 20,
		},
		{
			name:     "Sub of -10, 30",
			input:    []int{-10, 30},
			expected: -40,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Sub(tc.input...)
			if result != tc.expected {
				t.Errorf("Expected %d, got %d", tc.expected, result)
			}
		})
	}
}

func ShouldMultiply(t *testing.T) {
	testCases := []testCase{
		{
			name:     "Multiply of 9, 4",
			input:    []int{9, 4},
			expected: 36,
		},
		{
			name:     "Multiply of 10, 0",
			input:    []int{10, 0},
			expected: 0,
		},
		{
			name:     "Multiply of 10, -1",
			input:    []int{10, -1},
			expected: -10,
		},
		{
			name:     "Multiply of 10, -3, 5, -2",
			input:    []int{10, -3, 5, -2},
			expected: 300,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Multiply(tc.input...)
			if result != tc.expected {
				t.Errorf("Expected %d, got %d", tc.expected, result)
			}
		})
	}
}

func ShouldDivide(t *testing.T) {
	testCases := []testCase{
		{
			name:     "Divide of 81, 9",
			input:    []int{81, 9},
			expected: 9,
			errMsg:   "",
		},
		{
			name:     "Divide of 10, 0",
			input:    []int{10, 0},
			expected: 0,
			errMsg:   "Cannot divide by zero",
		},
		{
			name:     "Divide of 0, 10",
			input:    []int{0, 10},
			expected: 0,
			errMsg:   "",
		},
		{
			name:     "Divide of 10, -2",
			input:    []int{10, -2},
			expected: -5,
			errMsg:   "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := Divide(tc.input...)
			if err != nil {
				if tc.errMsg != err.Error() {
					t.Errorf("Expected %s, got %s", tc.errMsg, err.Error())
				}
			} else {
				if result != tc.expected {
					t.Errorf("Expected %d, got %d", tc.expected, result)
				}
			}
		})
	}
}
