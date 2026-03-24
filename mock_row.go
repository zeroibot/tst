package tst

import (
	"fmt"
	"reflect"
)

type Row struct {
	items []any
}

func NewRow(items ...any) *Row {
	return new(Row{items: items})
}

func (r *Row) Scan(fieldRefs ...any) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic encountered: %v", r)
		}
	}()
	if len(fieldRefs) != len(r.items) {
		return fmt.Errorf("expected %d fieldRefs, got %d", len(r.items), len(fieldRefs))
	}
	for i, fieldRef := range fieldRefs {
		fieldValue := reflect.ValueOf(fieldRef).Elem()
		fieldValue.Set(reflect.ValueOf(r.items[i]))
	}
	return err
}
