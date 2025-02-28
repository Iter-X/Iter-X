package repository

import "context"

// Transaction 事务封装接口
type Transaction interface {
	// Exec 按事务执行
	Exec(context.Context, func(ctx context.Context) error) error
}
