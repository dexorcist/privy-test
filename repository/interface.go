package repository

import (
	"context"
	"database/sql"
	"fmt"
	"privy-test/model"
	"privy-test/param/cake"
)

// ErrStmtNotFound error if retrieved prepared statement not registered in repo map.
var ErrStmtNotFound = fmt.Errorf("repo: prepared statement not found")

// Statement Interface to get registered prepared statement.
type Statement interface {
	Statement(ctx context.Context, name string) (*sql.Stmt, error)
}

type CakeRepository interface {
	Statement
	Create(ctx context.Context, cakeModel *model.Cake) (*model.Cake, error)
	Delete(ctx context.Context, cakeID int64) error
	Update(ctx context.Context, cakeID int64, cakeModel *model.Cake) (*model.Cake, error)
	GetDetail(ctx context.Context, cakeID int64) (*model.Cake, error)
	GetList(ctx context.Context, request *cake.FindAllRequest) ([]model.Cake, error)
}
