package isodd

import "testing"

func TestString(t *testing.T) {
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
		isOdd, err := String(test.num)

		if isOdd != test.isOdd || err != test.err {
			t.Fatalf("[TestString] {string: %s, err: %v} failed", test.num, err)
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
		isOdd, err := String(test.num)

		if isOdd != false || err == nil {
			t.Fatalf("[TestString] {string: %s, err: %v} failed", test.num, err)
		}
	}
}
