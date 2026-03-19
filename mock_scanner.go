package tst

import (
	"fmt"
	"reflect"
)

type MockScanner struct {
	items []any
}

func NewMockScanner(items ...any) MockScanner {
	return MockScanner{items: items}
}

func (m MockScanner) Scan(fieldRefs ...any) (err error) {
	defer func() {
		if err := recover(); err != nil {
			err = fmt.Errorf("panic encountered: %v", err)
		}
	}()
	if len(fieldRefs) != len(m.items) {
		return fmt.Errorf("expected %d fieldRefs, got %d", len(m.items), len(fieldRefs))
	}
	for i, fieldRef := range fieldRefs {
		fieldValue := reflect.ValueOf(fieldRef).Elem()
		fieldValue.Set(reflect.ValueOf(m.items[i]))
	}
	return err
}
