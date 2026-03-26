package tst

import (
	"database/sql"
	"errors"
	"slices"
)

var (
	errNotFound      = errors.New("not found")
	errMissingParams = errors.New("missing params")
)

type Conn[T any] struct {
	items  []T
	err    error
	testFn func(T) bool
	rowFn  func([]T) ([]any, error)
	rowsFn func(T) ([]any, error)
}

func NewConn[T any](items ...T) *Conn[T] {
	return new(Conn[T]{items: items, err: nil})
}

func (c *Conn[T]) Begin() (*Tx, error) {
	// TODO: Update for Tx
	return nil, c.err
}

func (c *Conn[T]) Exec(query string, args ...any) (sql.Result, error) {
	// TODO: Implement
	return nil, c.err
}

func (c *Conn[T]) Query(query string, args ...any) (*Rows, error) {
	if c.testFn == nil || c.rowsFn == nil || c.err != nil {
		err := errMissingParams
		if c.err != nil {
			err = c.err
		}
		return NewRows(), err
	}
	validItems := make([]T, 0, len(c.items))
	for _, item := range c.items {
		if c.testFn(item) {
			validItems = append(validItems, item)
		}
	}
	rowValues := make([][]any, 0, len(validItems))
	for _, item := range validItems {
		values, err := c.rowsFn(item)
		if err != nil {
			return NewRows(), err
		}
		rowValues = append(rowValues, values)
	}
	return NewRows(rowValues...), nil
}

func (c *Conn[T]) QueryRow(query string, args ...any) *Row {
	if c.testFn == nil || c.rowFn == nil || c.err != nil {
		return NewRow()
	}
	validItems := make([]T, 0, len(c.items))
	for _, item := range c.items {
		if c.testFn(item) {
			validItems = append(validItems, item)
		}
	}
	values, err := c.rowFn(validItems)
	if err != nil {
		return NewRow()
	}
	return NewRow(values...)
}

func (c *Conn[T]) SetError(err error) {
	c.err = err
}

func (c *Conn[T]) PrepRow(testFn func(T) bool, rowFn func([]T) ([]any, error)) func() {
	return func() {
		c.SetError(nil)
		c.testFn = testFn
		c.rowFn = rowFn
	}
}

func (c *Conn[T]) PrepOne(testFn func(T) bool, rowFn func(T) []any) func() {
	return func() {
		c.SetError(nil)
		c.testFn = testFn
		c.rowFn = func(items []T) ([]any, error) {
			if len(items) == 0 {
				return nil, errNotFound
			}
			return rowFn(items[0]), nil
		}
	}
}

func (c *Conn[T]) PrepSortOne(testFn func(T) bool, sortFn func(T, T) int, rowFn func(T) []any) func() {
	return func() {
		c.SetError(nil)
		c.testFn = testFn
		c.rowFn = func(items []T) ([]any, error) {
			if len(items) == 0 {
				return nil, errNotFound
			}
			slices.SortFunc(items, sortFn)
			return rowFn(items[0]), nil
		}
	}
}

func (c *Conn[T]) PrepRows(testFn func(T) bool, rowsFn func(T) ([]any, error)) func() {
	return func() {
		c.SetError(nil)
		c.testFn = testFn
		c.rowsFn = rowsFn
	}
}
