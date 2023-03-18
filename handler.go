package main

import (
	"context"
	"github.com/shuyi-tangerine/comment/gen-go/base"
	"github.com/shuyi-tangerine/comment/gen-go/tangerine/comment"
	"github.com/shuyi-tangerine/comment/idgen"
	"github.com/shuyi-tangerine/comment/memory"
	"github.com/shuyi-tangerine/comment/server"
	"github.com/shuyi-tangerine/comment/top"
)

type CommentHandler struct {
	commentService top.CommentService
	idGenerator    top.IDGenerator
}

func NewCommentHandler() comment.CommentHandler {

	idGenerator := idgen.NewTimestamp()
	commentStorage := memory.NewCommentStorage()
	commentService := server.NewCommentService(commentStorage, idGenerator)

	return &CommentHandler{
		commentService: commentService,
		idGenerator:    idGenerator,
	}
}

func (m *CommentHandler) GenCommentID(ctx context.Context, req *comment.GenCommentIDRequest) (res *comment.GenCommentIDResponse, err error) {
	res = &comment.GenCommentIDResponse{
		CommentIds: nil,
		Base: &base.RPCResponse{
			Code:    0,
			Message: "ok",
		},
	}

	defer func() {
		if err != nil {
			res = &comment.GenCommentIDResponse{
				Base: &base.RPCResponse{
					Code:    -1,
					Message: err.Error(),
				},
			}
			err = nil
		}
	}()

	amount := int8(1)
	if req.Amount != nil && *req.Amount > 0 {
		amount = *req.Amount
	}

	ids, err := m.idGenerator.Gen(ctx, amount)
	if err != nil {
		return
	}

	res.CommentIds = ids
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
