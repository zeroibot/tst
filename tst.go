// Package tst contains unit testing functions and TestCase structs
package tst

import "testing"

type P1W1[I, R any] struct {
	P1 I
	W1 R
}

type P2W1[I1, I2, R any] struct {
	P1 I1
	P2 I2
	W1 R
}

type P2W2[I1, I2, R1, R2 any] struct {
	P1 I1
	P2 I2
	W1 R1
	W2 R2
}

type P3W1[I1, I2, I3, R any] struct {
	P1 I1
	P2 I2
	P3 I3
	W1 R
}

// AllCompare1 tests 1 pair of [actual, want] from the given generic test cases
func AllCompare1[T, R any](t *testing.T, testCases []T, name string, testFn func(T) (R, R), assert assertFn[R]) {
	for _, x := range testCases {
		actual, want := testFn(x)
		assert(t, name, actual, want)
	}
}

// AllP1W1 tests all P1W1 test cases
func AllP1W1[I, R any](t *testing.T, testCases []P1W1[I, R], name string, testFn func(I) R, assert assertFn[R]) {
	for _, x := range testCases {
		actual := testFn(x.P1)
		assert(t, name, actual, x.W1)
	}
}

// AllP2W1 tests all P2W1 test cases
func AllP2W1[I1, I2, R any](t *testing.T, testCases []P2W1[I1, I2, R], name string, testFn func(I1, I2) R, assert assertFn[R]) {
	for _, x := range testCases {
		actual := testFn(x.P1, x.P2)
		assert(t, name, actual, x.W1)
	}
}

// AllP2W2 tests all P2W2 test cases
func AllP2W2[I1, I2, R1, R2 any](t *testing.T, testCases []P2W2[I1, I2, R1, R2], name string, testFn func(I1, I2) (R1, R2), assert1 assertFn[R1], assert2 assertFn[R2]) {
	for _, x := range testCases {
		actual1, actual2 := testFn(x.P1, x.P2)
		assert1(t, name, actual1, x.W1)
		assert2(t, name, actual2, x.W2)
	}
}

// AllP3W1 tests all P3W1 test cases
func AllP3W1[I1, I2, I3, R any](t *testing.T, testCases []P3W1[I1, I2, I3, R], name string, testFn func(I1, I2, I3) R, assert assertFn[R]) {
	for _, x := range testCases {
		actual := testFn(x.P1, x.P2, x.P3)
		assert(t, name, actual, x.W1)
	}
}
