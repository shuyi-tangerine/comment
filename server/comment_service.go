package server

import (
	"context"
	"github.com/shuyi-tangerine/comment/top"
)

type CommentService struct {
	commentStorage top.CommentStorage
	idGenerator    top.IDGenerator
}

func NewCommentService(commentStorage top.CommentStorage, idGenerator top.IDGenerator) top.CommentService {
	return &CommentService{
		commentStorage: commentStorage,
		idGenerator:    idGenerator,
	}
}

func (m *CommentService) Post(ctx context.Context, req *top.Comment) (postedComment *top.Comment, err error) {
	postReq := req.Clone(ctx)

	// 补全 comment_id
	if postReq.CommentID <= 0 {
		ids, err := m.idGenerator.Gen(ctx, 1)
		if err != nil {
			return nil, err
		}
		postReq.CommentID = ids[0]
	}

	id, err := m.commentStorage.InsertOne(ctx, postReq)
	if err != nil {
		return
	}

	postedComment, err = m.commentStorage.SelectOne(ctx, &top.OperateCommentRequest{
		ID:        id,
		CommentID: postReq.CommentID,
	})
	return
}
