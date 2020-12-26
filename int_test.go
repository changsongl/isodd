package isodd

import (
	"math"
	"testing"
)

func TestInt(t *testing.T) {
	type testcase struct {
		num   int
		isOdd bool
	}

	testCases := []testcase{
		{1, true},
		{-2, false},
		{0, false},
		{2, false},
	}

	for _, test := range testCases {
		isOdd := Int(test.num)

		if isOdd != test.isOdd {
			t.Fatalf("[TestInt] %d failed", test.num)
		}
	}
}

func TestInt8(t *testing.T) {
	type testcase struct {
		num   int8
		isOdd bool
	}

	testCases := []testcase{
		{1, true},
		{-2, false},
		{0, false},
		{2, false},
		{math.MaxInt8, true},  // 127
		{math.MinInt8, false}, // -128
	}

	for _, test := range testCases {
		isOdd := Int8(test.num)

		if isOdd != test.isOdd {
			t.Fatalf("[TestInt8] %d failed", test.num)
		}
	}
}

func TestInt16(t *testing.T) {
	type testcase struct {
		num   int16
		isOdd bool
	}

	testCases := []testcase{
		{1, true},
		{-2, false},
		{0, false},
		{2, false},
		{math.MaxInt16, true},
		{math.MinInt16, false},
	}

	for _, test := range testCases {
		isOdd := Int16(test.num)

		if isOdd != test.isOdd {
			t.Fatalf("[TestInt16] %d failed", test.num)
		}
	}
}

func TestInt32(t *testing.T) {
	type testcase struct {
		num   int32
		isOdd bool
	}

	testCases := []testcase{
		{1, true},
		{-2, false},
		{0, false},
		{2, false},
		{math.MaxInt32, true},
		{math.MinInt32, false},
	}

	for _, test := range testCases {
		isOdd := Int32(test.num)

		if isOdd != test.isOdd {
			t.Fatalf("[TestInt32] %d failed", test.num)
		}
	}
}

func TestInt64(t *testing.T) {
	type testcase struct {
		num   int64
		isOdd bool
	}

	testCases := []testcase{
		{1, true},
		{-2, false},
		{0, false},
		{2, false},
		{math.MaxInt64, true},
		{math.MinInt64, false},
	}

	for _, test := range testCases {
		isOdd := Int64(test.num)

		if isOdd != test.isOdd {
			t.Fatalf("[TestInt64] %d failed", test.num)
		}
	}
}

func TestUInt(t *testing.T) {
	type testcase struct {
		num   uint
		isOdd bool
	}

	testCases := []testcase{
		{1, true},
		{0, false},
		{2, false},
	}

	for _, test := range testCases {
		isOdd := Uint(test.num)

		if isOdd != test.isOdd {
			t.Fatalf("[TestUInt] %d failed", test.num)
		}
	}
}

func TestUint8(t *testing.T) {
	type testcase struct {
		num   uint8
		isOdd bool
	}

	testCases := []testcase{
		{1, true},
		{0, false},
		{2, false},
		{math.MaxUint8, true},
	}

	for _, test := range testCases {
		isOdd := Uint8(test.num)

		if isOdd != test.isOdd {
			t.Fatalf("[TestUint8] %d failed", test.num)
		}
	}
}

func TestUint16(t *testing.T) {
	type testcase struct {
		num   uint16
		isOdd bool
	}

	testCases := []testcase{
		{1, true},
		{0, false},
		{2, false},
		{math.MaxUint16, true},
	}

	for _, test := range testCases {
		isOdd := Uint16(test.num)

		if isOdd != test.isOdd {
			t.Fatalf("[TestUint16] %d failed", test.num)
		}
	}
}

func TestUint32(t *testing.T) {
	type testcase struct {
		num   uint32
		isOdd bool
	}

	testCases := []testcase{
		{1, true},
		{0, false},
		{2, false},
		{math.MaxUint32, true},
	}

	for _, test := range testCases {
		isOdd := Uint32(test.num)

		if isOdd != test.isOdd {
			t.Fatalf("[TestUint32] %d failed", test.num)
		}
	}
}

func TestUint64(t *testing.T) {
	type testcase struct {
		num   uint64
		isOdd bool
	}

	testCases := []testcase{
		{1, true},
		{0, false},
		{2, false},
		{math.MaxUint64, true},
	}

	for _, test := range testCases {
		isOdd := Uint64(test.num)

		if isOdd != test.isOdd {
			t.Fatalf("[TestUint64] %d failed", test.num)
		}
	}
}
