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
	_ = session.Begin()
	return &ConnectionTx{session}, nil
}
