package tst

import "database/sql"

type Conn[T any] struct {
	items []T
	err   error
	row   *Row
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

func (c *Conn[T]) QueryRow(query string, args ...any) *Row {
	return c.row
}

func (c *Conn[T]) Query(query string, args ...any) (*Rows, error) {
	// TODO: Implement
	return nil, c.err
}

func (c *Conn[T]) SetError(err error) {
	c.err = err
}

func (c *Conn[T]) SetRow(row *Row) {
	c.row = row
}
