package main

import (
	"context"
	"github.com/shuyi-tangerine/comment/gen-go/base"
	"github.com/shuyi-tangerine/comment/gen-go/tangerine/comment"
	"time"
)

type CommentHandler struct {
}

func NewCommentHandler() comment.CommentHandler {
	return &CommentHandler{}
}

// GenCommentID
// TODO 改成分布式ID生成算法
func (m *CommentHandler) GenCommentID(ctx context.Context, req *comment.GenCommentIDRequest) (res *comment.GenCommentIDResponse, _err error) {
	res = &comment.GenCommentIDResponse{
		CommentIds: nil,
		Base: &base.RPCResponse{
			Code:    0,
			Message: "ok",
		},
	}

	amount := 1
	if req.Amount != nil && *req.Amount > 0 {
		amount = int(*req.Amount)
	}

	baseID := time.Now().UnixMilli()
	res.CommentIds = append(res.CommentIds, baseID)
	for i := 1; i < amount; i++ {
		res.CommentIds = append(res.CommentIds, baseID+int64(i))
	}
	return
}

func (m *CommentHandler) Post(ctx context.Context, req *comment.PostRequest) (res *comment.PostResponse, _err error) {
	panic("not support")
}

func (m *CommentHandler) List(ctx context.Context, req *comment.ListRequest) (res *comment.ListResponse, _err error) {
	panic("not support")
}

func (m *CommentHandler) Delete(ctx context.Context, req *comment.DeleteRequest) (res *comment.DeleteResponse, _err error) {
	panic("not support")
}
