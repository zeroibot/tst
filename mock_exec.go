package tst

import "database/sql"

type Result struct {
	rowsAffected int64
	lastInsertID int64
	err          error
}

func NewResult(rowsAffected, lastInsertID int, err error) *Result {
	return new(Result{rowsAffected: int64(rowsAffected), lastInsertID: int64(lastInsertID), err: err})
}

func (r *Result) LastInsertId() (int64, error) {
	return r.lastInsertID, r.err
}

func (r *Result) RowsAffected() (int64, error) {
	return r.rowsAffected, r.err
}

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
