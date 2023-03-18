package top

import (
	"context"
)

type CommentStorage interface {
	// InsertOne 插入一条数据，并返回存储中的数据
	InsertOne(ctx context.Context, data *Comment) (id int64, err error)
	// SelectOne 只返回一条数据，一般用于主键、唯一键查询
	SelectOne(ctx context.Context, req *OperateCommentRequest) (storageComment *Comment, err error)
	// Select 返回列表
	Select(ctx context.Context, req *OperateCommentRequest) (storageComments []*Comment, err error)
	// UpdateOne 更新，按照主键、唯一键更新
	UpdateOne(ctx context.Context, req *OperateCommentRequest) (err error)
	// DeleteOne 删除，一般为逻辑删除,按照主键、唯一键删除
	DeleteOne(ctx context.Context, req *OperateCommentRequest) (err error)
}

type OperateCommentRequest struct {
	ID        int64 `json:"id"`         //  自增主键
	CommentID int64 `json:"comment_id"` //  评论id
	UserID    int64 `json:"user_id"`    //  用户id
	GroupID   int64 `json:"group-id"`   //  所属分组ID，比如视频、文章等
	AppID     int32 `json:"app_id"`     //  应用id
}
