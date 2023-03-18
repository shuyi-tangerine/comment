package idgen

import (
	"context"
	"time"
)

// Timestamp 单机时间戳生成
type Timestamp struct {
}

func NewTimestamp() *Timestamp {
	return &Timestamp{}
}

func (m *Timestamp) Gen(ctx context.Context, amount int8) (ids []int64, err error) {
	baseID := time.Now().UnixMilli()
	ids = append(ids, baseID)
	for i := 1; i < int(amount); i++ {
		ids = append(ids, baseID+int64(i))
	}
	return
}
