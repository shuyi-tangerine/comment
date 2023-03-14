package top

import "context"

type CommentService interface {
	// Post 发评
	Post(ctx context.Context, req *Comment) (postedComment *Comment, err error)
}
