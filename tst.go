// Package tst contains unit testing functions and TestCase structs
package tst

import (
	"fmt"
	"testing"
)

type HasPostTest interface {
	PostTest() bool
}

// P1W1 Test case with 1 input, 1 output
type P1W1[I, R any] struct {
	P1 I
	W1 R
}

// P1W2 Test case with 1 input, 2 outputs
type P1W2[I, R1, R2 any] struct {
	P1 I
	W1 R1
	W2 R2
}

// P2W1 Test case with 2 inputs, 1 output
type P2W1[I1, I2, R any] struct {
	P1 I1
	P2 I2
	W1 R
}

// P2W2 Test case with 2 inputs, 2 outputs
type P2W2[I1, I2, R1, R2 any] struct {
	P1 I1
	P2 I2
	W1 R1
	W2 R2
}

// P2W2Pre Test Case with 2 inputs, 2 outputs, and a `prepare` step
type P2W2Pre[I1, I2, R1, R2 any] struct {
	Prep func()
	P1   I1
	P2   I2
	W1   R1
	W2   R2
}

// P2W3Pre Test Case with 2 inputs, 3 outputs, and a `prepare` step
type P2W3Pre[I1, I2, R1, R2, R3 any] struct {
	Prep func()
	P1   I1
	P2   I2
	W1   R1
	W2   R2
	W3   R3
}

// P3W1 Test case with 3 inputs, 1 output
type P3W1[I1, I2, I3, R any] struct {
	P1 I1
	P2 I2
	P3 I3
	W1 R
}

// P3W2 Test case with 3 inputs, 2 outputs
type P3W2[I1, I2, I3, R1, R2 any] struct {
	P1 I1
	P2 I2
	P3 I3
	W1 R1
	W2 R2
}

// P3W1Post Test case with 3 inputs, 1 output, and a post-test
type P3W1Post[I1, I2, I3, R any] struct {
	P1   I1
	P2   I2
	P3   I3
	W1   R
	Test func() bool
}

// AllCompare1 tests 1 pair of [actual, want] from the given generic test cases
func AllCompare1[T, R any](t *testing.T, testCases []T, name string, testFn func(T) (R, R), assert assertFn[R]) {
	for i, x := range testCases {
		label := fmt.Sprintf("%s:%d", name, i)
		actual, want := testFn(x)
		assert(t, label, actual, want)
	}
}

// AllActionPost performs the action for all generic test cases and checks the post-test
func AllActionPost[T HasPostTest](t *testing.T, testCases []T, name string, actionFn func(T)) {
	for i, x := range testCases {
		label := fmt.Sprintf("%s:%d", name, i)
		actionFn(x)
		assertTest(t, label, x.PostTest)
	}
}

// All checks if all pairs are equal using the assert function
func All[T any](t *testing.T, pairs [][2]T, name string, assert assertFn[T]) {
	for i, pair := range pairs {
		label := fmt.Sprintf("%s:%d", name, i)
		a, b := pair[0], pair[1]
		assert(t, label, a, b)
	}
}

// AllP1W1 tests all P1W1 test cases
func AllP1W1[I, R any](t *testing.T, testCases []P1W1[I, R], name string, testFn func(I) R, assert assertFn[R]) {
	for i, x := range testCases {
		label := fmt.Sprintf("%s:%d", name, i)
		actual := testFn(x.P1)
		assert(t, label, actual, x.W1)
	}
}

// AllP1W2 tests all P1W2 test cases
func AllP1W2[I, R1, R2 any](t *testing.T, testCases []P1W2[I, R1, R2], name string, testFn func(I) (R1, R2), assert1 assertFn[R1], assert2 assertFn[R2]) {
	for i, x := range testCases {
		label := fmt.Sprintf("%s:%d", name, i)
		actual1, actual2 := testFn(x.P1)
		assert1(t, label, actual1, x.W1)
		assert2(t, label, actual2, x.W2)
	}
}

// AllP2W1 tests all P2W1 test cases
func AllP2W1[I1, I2, R any](t *testing.T, testCases []P2W1[I1, I2, R], name string, testFn func(I1, I2) R, assert assertFn[R]) {
	for i, x := range testCases {
		label := fmt.Sprintf("%s:%d", name, i)
		actual := testFn(x.P1, x.P2)
		assert(t, label, actual, x.W1)
	}
}

// AllP2W2 tests all P2W2 test cases
func AllP2W2[I1, I2, R1, R2 any](t *testing.T, testCases []P2W2[I1, I2, R1, R2], name string, testFn func(I1, I2) (R1, R2), assert1 assertFn[R1], assert2 assertFn[R2]) {
	for i, x := range testCases {
		label := fmt.Sprintf("%s:%d", name, i)
		actual1, actual2 := testFn(x.P1, x.P2)
		assert1(t, label, actual1, x.W1)
		assert2(t, label, actual2, x.W2)
	}
}

// AllP2W2Pre tests all P2W2Pre test cases
func AllP2W2Pre[I1, I2, R1, R2 any](t *testing.T, testCases []P2W2Pre[I1, I2, R1, R2], name string, testFn func(I1, I2) (R1, R2), assert1 assertFn[R1], assert2 assertFn[R2]) {
	for i, x := range testCases {
		label := fmt.Sprintf("%s:%d", name, i)
		if x.Prep != nil {
			x.Prep()
		}
		actual1, actual2 := testFn(x.P1, x.P2)
		assert1(t, label, actual1, x.W1)
		assert2(t, label, actual2, x.W2)
	}
}

// AllP2W3Pre tests all P2W3Pre test cases
func AllP2W3Pre[I1, I2, R1, R2, R3 any](t *testing.T, testCases []P2W3Pre[I1, I2, R1, R2, R3], name string, testFn func(I1, I2) (R1, R2, R3), assert1 assertFn[R1], assert2 assertFn[R2], assert3 assertFn[R3]) {
	for i, x := range testCases {
		label := fmt.Sprintf("%s:%d", name, i)
		if x.Prep != nil {
			x.Prep()
		}
		actual1, actual2, actual3 := testFn(x.P1, x.P2)
		assert1(t, label, actual1, x.W1)
		assert2(t, label, actual2, x.W2)
		assert3(t, label, actual3, x.W3)
	}
}

// AllP3W1 tests all P3W1 test cases
func AllP3W1[I1, I2, I3, R any](t *testing.T, testCases []P3W1[I1, I2, I3, R], name string, testFn func(I1, I2, I3) R, assert assertFn[R]) {
	for i, x := range testCases {
		label := fmt.Sprintf("%s:%d", name, i)
		actual := testFn(x.P1, x.P2, x.P3)
		assert(t, label, actual, x.W1)
	}
}

// AllP3W1Post tests all P3W1Post test cases and checks the post-test function
func AllP3W1Post[I1, I2, I3, R any](t *testing.T, testCases []P3W1Post[I1, I2, I3, R], name string, testFn func(I1, I2, I3) R, assert assertFn[R]) {
	for i, x := range testCases {
		label := fmt.Sprintf("%s:%d", name, i)
		actual := testFn(x.P1, x.P2, x.P3)
		assert(t, label, actual, x.W1)
		assertTest(t, label, x.Test)
	}
}

// AllP3W2 tests all P3W2 test cases
func AllP3W2[I1, I2, I3, R1, R2 any](t *testing.T, testCases []P3W2[I1, I2, I3, R1, R2], name string, testFn func(I1, I2, I3) (R1, R2), assert1 assertFn[R1], assert2 assertFn[R2]) {
	for i, x := range testCases {
		label := fmt.Sprintf("%s:%d", name, i)
		actual1, actual2 := testFn(x.P1, x.P2, x.P3)
		assert1(t, label, actual1, x.W1)
		assert2(t, label, actual2, x.W2)
	}
}
