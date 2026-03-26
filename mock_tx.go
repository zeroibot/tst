package tst

import "database/sql"

type Tx struct {
	result sql.Result
	err    error
}

func NewTx() *Tx {
	return new(Tx)
}

func NewTxFrom(result sql.Result, err error) *Tx {
	return new(Tx{result: result, err: err})
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

func (t *Tx) SetResult(result sql.Result) {
	t.result = result
}
