package isodd

import (
	"math"
	"testing"
)

func TestInterfaceNoMatch(t *testing.T) {
	var (
		i   *int   = nil
		i8  *int8  = nil
		i16 *int16 = nil
	)

	testCases := []interface{}{i, i8, i16}
	for _, num := range testCases {
		isOdd, err := Interface(num)
		if isOdd != false || err != ErrorInterfaceNoMatch {
			t.Fatalf("[TestInterfaceNoMatch] %v, %v failed", num, err)
		}
	}
}

func TestInterfaceInt(t *testing.T) {
	type testcase struct {
		num   int
		isOdd bool
	}

	testCases := []testcase{
		{num: 1, isOdd: true},
		{num: -2, isOdd: false},
		{num: 0, isOdd: false},
		{num: 2, isOdd: false},
	}

	for _, test := range testCases {
		isOdd, err := Interface(test.num)

		if isOdd != test.isOdd || err != nil {
			t.Fatalf("[TestInterfaceInt] {string: %d, err: %v} failed", test.num, err)
		}
	}
}

func TestInterfaceInt8(t *testing.T) {
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
		isOdd, err := Interface(test.num)

		if isOdd != test.isOdd || err != nil {
			t.Fatalf("[TestInterfaceInt8] {string: %d, err: %v} failed", test.num, err)
		}
	}
}

func TestInterfaceInt16(t *testing.T) {
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
		isOdd, err := Interface(test.num)

		if isOdd != test.isOdd || err != nil {
			t.Fatalf("[TestInterfaceInt16] {string: %d, err: %v} failed", test.num, err)
		}
	}
}

func TestInterfaceInt32(t *testing.T) {
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
		isOdd, err := Interface(test.num)

		if isOdd != test.isOdd || err != nil {
			t.Fatalf("[TestInterfaceInt32] {string: %d, err: %v} failed", test.num, err)
		}
	}
}

func TestInterfaceInt64(t *testing.T) {
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
		isOdd, err := Interface(test.num)

		if isOdd != test.isOdd || err != nil {
			t.Fatalf("[TestInterfaceInt64] {string: %d, err: %v} failed", test.num, err)
		}
	}
}

func TestInterfaceUint(t *testing.T) {
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
		isOdd, err := Interface(test.num)

		if isOdd != test.isOdd || err != nil {
			t.Fatalf("[TestInterfaceUint] {string: %d, err: %v} failed", test.num, err)
		}
	}
}

func TestInterfaceUint8(t *testing.T) {
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
		isOdd, err := Interface(test.num)

		if isOdd != test.isOdd || err != nil {
			t.Fatalf("[TestInterfaceUint8] {string: %d, err: %v} failed", test.num, err)
		}
	}
}

func TestInterfaceUint16(t *testing.T) {
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
		isOdd, err := Interface(test.num)

		if isOdd != test.isOdd || err != nil {
			t.Fatalf("[TestInterfaceUint16] {string: %d, err: %v} failed", test.num, err)
		}
	}
}

func TestInterfaceUint32(t *testing.T) {
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
		isOdd, err := Interface(test.num)

		if isOdd != test.isOdd || err != nil {
			t.Fatalf("[TestInterfaceUint32] {string: %d, err: %v} failed", test.num, err)
		}
	}
}

func TestInterfaceUint64(t *testing.T) {
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
		isOdd, err := Interface(test.num)

		if isOdd != test.isOdd || err != nil {
			t.Fatalf("[TestInterfaceUint64] {string: %d, err: %v} failed", test.num, err)
		}
	}
}

func TestInterfaceString(t *testing.T) {
	type testcase struct {
		num   string
		isOdd bool
		err   error
	}

	testCases := []testcase{
		{"1", true, nil},
		{"0", false, nil},
		{"2", false, nil},
		{"-200000", false, nil},
		{"-999999", true, nil},
	}

	for _, test := range testCases {
		isOdd, err := Interface(test.num)

		if isOdd != test.isOdd || err != test.err {
			t.Fatalf("[TestInterfaceString] {string: %s, err: %v} failed", test.num, err)
		}
	}

	testCasesErr := []testcase{
		{num: "1.232"},
		{num: "asdasds"},
		{num: "1-"},
		{num: "xx1"},
		{num: " 1232"},
	}

	for _, test := range testCasesErr {
		isOdd, err := Interface(test.num)

		if isOdd != false || err == nil {
			t.Fatalf("[TestInterfaceString] {string: %s, err: %v} failed", test.num, err)
		}
	}
}

func TestInterfaceFloat32(t *testing.T) {
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
		isOdd, err := Interface(test.num)

		if isOdd != test.isOdd || err != nil {
			t.Fatalf("[TestInterfaceFloat32] %f failed", test.num)
		}
	}
}

func TestInterfaceFloat64(t *testing.T) {
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
		isOdd, err := Interface(test.num)

		if isOdd != test.isOdd || err != nil {
			t.Fatalf("[TestInterfaceFloat64] %f failed", test.num)
		}
	}
}
