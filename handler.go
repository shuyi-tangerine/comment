package main

import (
	"context"
	"github.com/shuyi-tangerine/comment/gen-go/base"
	"github.com/shuyi-tangerine/comment/gen-go/tangerine/comment"
	"github.com/shuyi-tangerine/comment/top"
	"time"
)

type CommentHandler struct {
	commentService top.CommentService
}

func NewCommentHandler(commentService top.CommentService) comment.CommentHandler {
	return &CommentHandler{
		commentService: commentService,
	}
}

// GenCommentID
// TODO 改成分布式ID生成算法
func (m *CommentHandler) GenCommentID(ctx context.Context, req *comment.GenCommentIDRequest) (res *comment.GenCommentIDResponse, err error) {
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

func (m *CommentHandler) Post(ctx context.Context, req *comment.PostRequest) (res *comment.PostResponse, err error) {
	res = &comment.PostResponse{
		Base: &base.RPCResponse{
			Code:    0,
			Message: "",
		},
	}

	defer func() {
		if err != nil {
			res = &comment.PostResponse{
				Base: &base.RPCResponse{
					Code:    -1,
					Message: err.Error(),
				},
			}
			err = nil
		}
	}()

	postReq, err := top.NewCommentByPostRequest(req)
	if err != nil {
		return
	}

	postRes, err := m.commentService.Post(ctx, postReq)
	if err != nil {
		return
	}

	commentData, err := postRes.ToCommentData(ctx)
	if err != nil {
		return
	}

	res.CommentData = commentData
	return
}

func (m *CommentHandler) List(ctx context.Context, req *comment.ListRequest) (res *comment.ListResponse, err error) {
	panic("not support")
}

func (m *CommentHandler) Delete(ctx context.Context, req *comment.DeleteRequest) (res *comment.DeleteResponse, err error) {
	panic("not support")
}
