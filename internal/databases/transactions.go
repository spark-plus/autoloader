package databases

import "xorm.io/xorm"

type Tx interface {
	Commit() error
	Rollback() error
}

type ConnectionTx struct {
	Session *xorm.Session
}

func (t *ConnectionTx) Commit() error {
	return t.Session.Commit()
}

func (t *ConnectionTx) Rollback() error {
	return t.Session.Rollback()
}
