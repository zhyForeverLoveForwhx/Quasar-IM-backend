package db

import (
	"context"
	"database/sql"
	"fmt"
)

// Store 接口
type Store interface {
	Querier
	//事务添加在下面
}

// Store 提供了所有查询数据库的函数和事务
type SQLStore struct {
	*Queries
	db *sql.DB
}

// NewStore creates a new Store
func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}

//事务执行函数,不对外暴露
func (store *SQLStore) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	//执行事务
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err:%v, rb err:%v", err, rbErr)
		}
	}

	return tx.Commit()
}

//TODO: 下面添加事务
