package tst

import "database/sql"

type Conn struct {
	row    *Row
	rows   *Rows
	err    error
	result sql.Result
}

type Tx struct {
	err    error
	result sql.Result
}

func NewConn() *Conn {
	return new(Conn)
}

func NewTx() *Tx {
	return new(Tx)
}

func (c *Conn) QueryRow(query string, args ...any) *Row {
	return c.row
}

func (c *Conn) SetRow(row *Row) {
	c.row = row
}

func (c *Conn) SetError(err error) {
	c.err = err
}

func (c *Conn) SetRows(rows *Rows) {
	c.rows = rows
}

func (c *Conn) Query(query string, args ...any) (*Rows, error) {
	return c.rows, c.err
}

func (c *Conn) SetResult(result sql.Result) {
	c.result = result
}

func (c *Conn) Exec(query string, args ...any) (sql.Result, error) {
	return c.result, c.err
}

func (c *Conn) Begin() (*Tx, error) {
	// TODO: Update for Tx
	return nil, c.err
}

func (t *Tx) SetError(err error) {
	t.err = err
}

func (t *Tx) SetResult(result sql.Result) {
	t.result = result
}

func (t *Tx) Exec(query string, args ...any) (sql.Result, error) {
	return t.result, t.err
}

func (t *Tx) Commit() error {
	return t.err
}

func (t *Tx) Rollback() error {
	return t.err
}
