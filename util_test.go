package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmod(t *testing.T) {
	testcases := []struct {
		name   string
		input  int
		mod    int
		result int
	}{
		{
			name:   "standart",
			input:  10,
			mod:    100,
			result: 10,
		},
		{
			name:   "negative input",
			input:  -10,
			mod:    100,
			result: 90,
		},
		{
			name:   "greater input",
			input:  110,
			mod:    100,
			result: 10,
		},
		{
			name:   "negative greater input",
			input:  -110,
			mod:    100,
			result: 90,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.result, emod(tc.input, tc.mod))
		})
	}
}

func TestClamp(t *testing.T) {
	testcases := []struct {
		name     string
		input    float64
		lower    float64
		higher   float64
		expected float64
	}{
		{
			name:     "standart",
			input:    2.0,
			lower:    0.0,
			higher:   10.0,
			expected: 2.0,
		},
		{
			name:     "low than lower",
			input:    -2.0,
			lower:    0.0,
			higher:   10.0,
			expected: 0.0,
		},
		{
			name:     "higher than high",
			input:    20.0,
			lower:    0.0,
			higher:   10.0,
			expected: 10.0,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			clamp(&tc.input, tc.lower, tc.higher)
			assert.Equal(t, tc.expected, tc.input)
		})
	}

}
