package tst

import (
	"maps"
	"reflect"
	"slices"
	"testing"
)

type assertFn[T any] = func(*testing.T, string, T, T)

// Convert maps the list of items to a new list using the conversion function
func Convert[T any](items []T, convert func(T) T) []T {
	items2 := make([]T, len(items))
	for i, item := range items {
		items2[i] = convert(item)
	}
	return items2
}

// AssertTrue asserts that the given condition is true
func AssertTrue(t *testing.T, name string, condition bool) {
	if condition != true {
		t.Errorf("%s = condition failed", name)
	}
}

// AssertFalse asserts that the given condition is false
func AssertFalse(t *testing.T, name string, condition bool) {
	if condition != false {
		t.Errorf("%s = condition failed", name)
	}
}

// AssertDeepEqual asserts that the two `any` items are deeply equal
func AssertDeepEqual(t *testing.T, name string, a, b any) {
	if reflect.DeepEqual(a, b) == false {
		t.Errorf("%s = %v, want %v", name, a, b)
	}
}

// AssertDeepEqualAnd asserts that the two `any` items are deeply equal and the boolean flags are equal
func AssertDeepEqualAnd(t *testing.T, name string, a, b any, flag1, flag2 bool) {
	if flag1 != flag2 || reflect.DeepEqual(a, b) == false {
		t.Errorf("%s = %v, %t, want %v, %t", name, a, flag1, b, flag2)
	}
}

// AssertDeepEqualError asserts that the two `any` items are deeply equal and the error follows notNil flag
func AssertDeepEqualError(t *testing.T, name string, a, b any, err error, notNil bool) {
	if (err != nil) != notNil || reflect.DeepEqual(a, b) == false {
		t.Errorf("%s = %v, %v, want %v, %v", name, a, err, b, notNilString(notNil))
	}
}

// AssertEqual asserts that the two given values are equal
func AssertEqual[T comparable](t *testing.T, name string, a, b T) {
	if a != b {
		t.Errorf("%s = %v, want %v", name, a, b)
	}
}

// AssertEqual2 asserts that the two pairs of values are all equal
func AssertEqual2[T1, T2 comparable](t *testing.T, name string, a1, b1 T1, a2, b2 T2) {
	if a1 != b1 || a2 != b2 {
		t.Errorf("%s = %v, %v, want %v, %v", name, a1, a2, b1, b2)
	}
}

// AssertEqual3 asserts that the three pairs of values are all equal
func AssertEqual3[T1, T2, T3 comparable](t *testing.T, name string, a1, b1 T1, a2, b2 T2, a3, b3 T3) {
	if a1 != b1 || a2 != b2 || a3 != b3 {
		t.Errorf("%s = %v, %v, %v, want %v, %v, %v", name, a1, a2, a3, b1, b2, b3)
	}
}

// AssertEqual4 asserts that the four pairs of values are all equal
func AssertEqual4[T1, T2, T3, T4 comparable](t *testing.T, name string, a1, b1 T1, a2, b2 T2, a3, b3 T3, a4, b4 T4) {
	if a1 != b1 || a2 != b2 || a3 != b3 || a4 != b4 {
		t.Errorf("%s = %v, %v, %v, %v,  want %v, %v,%v, %v", name, a1, a2, a3, a4, b1, b2, b3, b4)
	}
}

// AssertEqualAnd asserts that the two given values are equal and the boolean flags are equal
func AssertEqualAnd[T comparable](t *testing.T, name string, a, b T, flag1, flag2 bool) {
	if flag1 != flag2 || a != b {
		t.Errorf("%s = %v, %t, want %v, %t", name, a, flag1, b, flag2)
	}
}

// AssertEqualError asserts that the two given values are equal and the error follows notNil flag
func AssertEqualError[T comparable](t *testing.T, name string, a, b T, err error, notNil bool) {
	if (err != nil) != notNil || a != b {
		t.Errorf("%s = %v, %v, want %v, %v", name, a, err, b, notNilString(notNil))
	}
}

// AssertEqualAny checks if two `any` items are equal
func AssertEqualAny(t *testing.T, name string, a, b any) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("%s panicked", name)
		}
	}()
	if a != b {
		t.Errorf("%s = %v, want %v", name, a, b)
	}
}

// AssertEqualAnyAnd checks if two `any` items are equal and the boolean flags are equal
func AssertEqualAnyAnd(t *testing.T, name string, a, b any, flag1, flag2 bool) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("%s panicked", name)
		}
	}()
	if flag1 != flag2 || a != b {
		t.Errorf("%s = %v, %t, want %v, %t", name, a, flag1, b, flag2)
	}
}

// AssertEqualAnyError checks if two `any` items are equal and the error follows notNil flag
func AssertEqualAnyError(t *testing.T, name string, a, b any, err error, notNil bool) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("%s panicked", name)
		}
	}()
	if (err != nil) != notNil || a != b {
		t.Errorf("%s = %v, %v, want %v, %v", name, a, err, b, notNilString(notNil))
	}
}

// AssertListEqual asserts that the two given lists are equal
func AssertListEqual[S ~[]T, T comparable](t *testing.T, name string, a, b S) {
	if slices.Equal(a, b) == false {
		t.Errorf("%s = %v, want %v", name, a, b)
	}
}

// AssertListEqualAnd asserts that the two given lists are equal and the boolean flags are equal
func AssertListEqualAnd[S ~[]T, T comparable](t *testing.T, name string, a, b S, flag1, flag2 bool) {
	if flag1 != flag2 || slices.Equal(a, b) == false {
		t.Errorf("%s = %v, %t, want %v, %t", name, a, flag1, b, flag2)
	}
}

// AssertListEqualError asserts that the two given lists are equal and the error follows notNil flag
func AssertListEqualError[S ~[]T, T comparable](t *testing.T, name string, a, b S, err error, notNil bool) {
	if (err != nil) != notNil || slices.Equal(a, b) == false {
		t.Errorf("%s = %v, %v, want %v, %v", name, a, err, b, notNilString(notNil))
	}
}

// AssertMapEqual asserts that the two given maps are equal
func AssertMapEqual[M1, M2 ~map[K]V, K, V comparable](t *testing.T, name string, a M1, b M2) {
	if maps.Equal(a, b) == false {
		t.Errorf("%s = %v, want %v", name, a, b)
	}
}

// AssertMapEqualAnd asserts that the two given maps are equal and the boolean flags are equal
func AssertMapEqualAnd[M1, M2 ~map[K]V, K, V comparable](t *testing.T, name string, a M1, b M2, flag1, flag2 bool) {
	if flag1 != flag2 || maps.Equal(a, b) == false {
		t.Errorf("%s = %v, %t, want %v, %t", name, a, flag1, b, flag2)
	}
}

// AssertMapEqualError asserts that the two given maps are equal and the error follows notNil flag
func AssertMapEqualError[M1, M2 ~map[K]V, K, V comparable](t *testing.T, name string, a M1, b M2, err error, notNil bool) {
	if (err != nil) != notNil || maps.Equal(a, b) == false {
		t.Errorf("%s = %v, %v, want %v, %v", name, a, err, b, notNilString(notNil))
	}
}

// AssertPanic asserts that the end of the function will panic
// Usage: defer AssertPanic(t, name)
func AssertPanic(t *testing.T, name string) {
	if err := recover(); err == nil {
		t.Errorf("%s did not panic", name)
	}
}

// assertTest calls the test function
func assertTest(t *testing.T, name string, test func() bool) {
	if test != nil && test() == false {
		t.Errorf("%s post test failed", name)
	}
}

// Returns a string for notNil boolean
func notNilString(notNil bool) string {
	if notNil {
		return "<error>"
	}
	return "<nil>"
}
