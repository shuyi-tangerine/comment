package main

import (
	"context"
	"fmt"
	"github.com/shuyi-tangerine/comment/gen-go/base"
	"github.com/shuyi-tangerine/comment/gen-go/tangerine/comment"
	"testing"
)

func TestCommentHandler_Post(t *testing.T) {
	handler := NewCommentHandler()
	res, err := handler.Post(context.Background(), &comment.PostRequest{
		CommentID:      nil,
		UserID:         2023,
		GroupID:        2023,
		AppID:          100,
		Text:           "你好，我是第一条测试评论！",
		Images:         nil,
		ReplyCommentID: nil,
		Extra:          nil,
		Base:           &base.RPCRequest{},
	})
	fmt.Println(res, err)
}
