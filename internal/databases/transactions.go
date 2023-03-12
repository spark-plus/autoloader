package databases

import "xorm.io/xorm"

type Tx interface {
	Commit() error
	Rollback() error
}

type ConnectionTx struct {
	session *xorm.Session
}

func (t *ConnectionTx) Commit() error {
	return t.session.Commit()
}

func (t *ConnectionTx) Rollback() error {
	return t.session.Rollback()
}
