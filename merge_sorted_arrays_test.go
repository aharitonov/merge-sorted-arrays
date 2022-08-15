package main_test

import (
	"fmt"
	"testing"
	. "example.org/main"
)

type numer = Numeric

type testCase struct {
	inA []numer
	inB []numer
	out []numer
	isPositive bool
}

func newCase(inA, inB, out []numer, valid_optional ...bool) testCase {
	valid := true
	if len(valid_optional) > 0 {
		valid = valid_optional[0]
	}
	return testCase{inA, inB, out, valid}
}

func (me testCase) Name() string {
	if (me.isPositive) {
		return fmt.Sprintf("F(%d) should be equal %d", me.inA, me.out)
	} else {
		return fmt.Sprintf("F(%d) should NOT be equal %d", me.inA, me.out)
	}
}

func (me testCase) Run(t *testing.T) {
	result := MergeSortedArrays(me.inA, me.inB)
	if me.isPositive && !isArrayEquals(result, me.out) || !me.isPositive && isArrayEquals(result, me.out) {
		t.Errorf("F(%d, %d): expected %d, actual %v", me.inA, me.inB, me.out, result)
	}
}

func isArrayEquals(a []numer, b []numer) bool {
    if len(a) != len(b) {
        return false
    }
    for i, v := range a {
        if v != b[i] {
            return false
        }
    }
    return true
}

func TestMain(t *testing.T) {
	for _, tc := range []testCase{
		newCase([]numer{1,2,3}, []numer{2,3,4}, []numer{1,2,2,3,3,4}),
		newCase([]numer{1,2,3}, []numer{2,3,4}, []numer{1,2,2,3,3,4,5}, false),
	} { t.Run(tc.Name(), tc.Run) }
}
