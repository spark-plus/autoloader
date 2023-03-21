package databases

import (
	"context"

	"xorm.io/xorm"
)

type Tx interface {
	DoTransaction(ctx context.Context, fn func(session *xorm.Session) error) error
	NewSession() *xorm.Session
}

type transactionManager struct {
	engine *xorm.Engine
}

func NewTransactionManager(engine *xorm.Engine) Tx {
	return &transactionManager{engine: engine}
}

func (tm *transactionManager) DoTransaction(ctx context.Context, fn func(session *xorm.Session) error) error {
	session := tm.engine.NewSession()
	defer session.Close()

	if err := session.Begin(); err != nil {
		return err
	}

	if err := fn(session); err != nil {
		session.Rollback()
		return err
	}

	return session.Commit()
}

func (tm *transactionManager) NewSession() *xorm.Session {
	session := tm.engine.NewSession()
	return session
}
