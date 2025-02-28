package repository

import "context"

// TransactionRepo 事务封装接口
type TransactionRepo interface {
	// Exec 按事务执行
	Exec(context.Context, func(ctx context.Context) error) error
}
