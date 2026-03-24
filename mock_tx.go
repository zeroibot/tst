package tst

import "database/sql"

type Tx struct {
	err    error
	result sql.Result
}

func NewTx() *Tx {
	return new(Tx)
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

func (t *Tx) SetError(err error) {
	t.err = err
}
