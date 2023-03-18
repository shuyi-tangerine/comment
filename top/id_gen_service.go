package top

import "context"

type IDGenerator interface {
	// Gen 生成 amount 个 ID
	Gen(ctx context.Context, amount int8) (ids []int64, err error)
}
