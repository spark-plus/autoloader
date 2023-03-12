package databases

import (
	"context"

	"xorm.io/xorm"
)

type DBSession struct {
	engine *xorm.Engine
}

func (r *DBSession) BeginTx(ctx context.Context) (Tx, error) {
	session := r.engine.NewSession()
	err := session.Begin()
	if err != nil {
		return nil, err
	}
	return &ConnectionTx{session}, nil
}

func (r *DBSession) GetSession(tx Tx) *xorm.Session {
	return tx.(*ConnectionTx).session
}

func (r *DBSession) NewSession(ctx context.Context) *xorm.Session {
	return r.engine.NewSession()
}
