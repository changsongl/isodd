package isodd

import (
	"testing"
)

func TestFloat32(t *testing.T) {
	type testcase struct {
		num   float32
		isOdd bool
	}

	testCases := []testcase{
		{1.232458, true},
		{0.324, false},
		{2.94038598, false},
		{-200000.94038598, false},
		{-999999.232458, true},
	}

	for _, test := range testCases {
		isOdd := Float32(test.num)

		if isOdd != test.isOdd {
			t.Fatalf("[TestFloat32] %f failed", test.num)
		}
	}
}

func TestFloat64(t *testing.T) {
	type testcase struct {
		num   float64
		isOdd bool
	}

	testCases := []testcase{
		{1.232458, true},
		{0.324, false},
		{2.94038598, false},
		{-200000.94038598, false},
		{-999999.232458, true},
	}

	for _, test := range testCases {
		isOdd := Float64(test.num)

		if isOdd != test.isOdd {
			t.Fatalf("[TestFloat64] %f failed", test.num)
		}
	}
}
