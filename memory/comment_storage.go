package memory

import (
	"context"
	"fmt"
	"github.com/shuyi-tangerine/comment/top"
	"time"
)

type CommentStorage struct {
	comments []*top.Comment
}

func NewCommentStorage() top.CommentStorage {
	return &CommentStorage{}
}

func (m *CommentStorage) InsertOne(ctx context.Context, data *top.Comment) (id int64, err error) {
	for _, v := range m.comments {
		if v.CommentID == data.CommentID {
			return 0, fmt.Errorf("comment_id dulicated")
		}
	}

	newComment := data.Clone(ctx)
	newComment.ID = int64(len(m.comments)) + 1
	newComment.CreatedAt = time.Now()
	newComment.UpdatedAt = newComment.CreatedAt

	m.comments = append(m.comments, newComment)
	return newComment.ID, nil
}

func (m *CommentStorage) SelectOne(ctx context.Context, req *top.OperateCommentRequest) (storageComment *top.Comment, err error) {
	for _, v := range m.comments {
		if req.CommentID > 0 && v.CommentID == req.CommentID || req.ID > 0 && v.ID == req.ID {
			return v.Clone(ctx), nil
		}
	}
	return nil, nil
}

func (m *CommentStorage) Select(ctx context.Context, req *top.OperateCommentRequest) (storageComments []*top.Comment, err error) {
	//TODO implement me
	panic("implement me")
}

func (m *CommentStorage) UpdateOne(ctx context.Context, req *top.OperateCommentRequest) (err error) {
	//TODO implement me
	panic("implement me")
}

func (m *CommentStorage) DeleteOne(ctx context.Context, req *top.OperateCommentRequest) (err error) {
	//TODO implement me
	panic("implement me")
}
