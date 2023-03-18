package top

import (
	"context"
	"encoding/json"
	"github.com/shuyi-tangerine/comment/gen-go/tangerine/comment"
	"time"
)

// Comment  评论
type Comment struct {
	ID             int64     `json:"id"`               //  自增主键
	CommentID      int64     `json:"comment_id"`       //  评论id
	UserID         int64     `json:"user_id"`          //  用户id
	GroupID        int64     `json:"group-id"`         //  所属分组ID，比如视频、文章等
	AppID          int32     `json:"app_id"`           //  应用id
	Text           string    `json:"text"`             //  文本
	Images         *string   `json:"images"`           //  图片
	ReplyCommentID *int64    `json:"reply_comment_id"` // 回复的评论
	CreatedAt      time.Time `json:"created_at"`       //  创建时间，也就是通知发送时间
	UpdatedAt      time.Time `json:"updated_at"`       //  最后更新时间
	Status         int8      `json:"status"`           //  状态，1-正常，2-已删除
	Extra          *string   `json:"extra"`            //  一些额外的信息
}

func NewCommentByPostRequest(req *comment.PostRequest) (c *Comment, err error) {
	c = &Comment{
		UserID:         req.UserID,
		GroupID:        req.GroupID,
		AppID:          req.AppID,
		Text:           req.Text,
		ReplyCommentID: req.ReplyCommentID,
		Status:         int8(comment.CommentStatus_Normal),
		Extra:          req.Extra,
	}

	if req.CommentID != nil {
		c.CommentID = *req.CommentID
	}

	if len(req.Images) > 0 {
		bts, err := json.Marshal(req.Images)
		if err != nil {
			return nil, err
		}
		images := string(bts)
		c.Images = &images
	}

	return
}

func (m *Comment) ToCommentData(ctx context.Context) (commentData *comment.CommentData, err error) {
	commentData = &comment.CommentData{
		ID:             m.ID,
		CommentID:      m.CommentID,
		UserID:         m.UserID,
		GroupID:        m.GroupID,
		AppID:          m.AppID,
		Text:           m.Text,
		ReplyCommentID: m.ReplyCommentID,
		CreatedAt:      m.CreatedAt.Unix(),
		UpdatedAt:      m.UpdatedAt.Unix(),
		Extra:          m.Extra,
	}

	images, err := m.ToListImages(ctx)
	if err != nil {
		return
	}

	commentData.Images = images
	return
}

func (m *Comment) ToListImages(ctx context.Context) (images []string, err error) {
	if m.Images == nil {
		return
	}
	err = json.Unmarshal([]byte(*m.Images), &images)
	return
}

func (m *Comment) Clone(ctx context.Context) (c *Comment) {
	return &Comment{
		ID:             m.ID,
		CommentID:      m.CommentID,
		UserID:         m.UserID,
		GroupID:        m.GroupID,
		AppID:          m.AppID,
		Text:           m.Text,
		Images:         m.Images,
		ReplyCommentID: m.ReplyCommentID,
		CreatedAt:      m.CreatedAt,
		UpdatedAt:      m.UpdatedAt,
		Status:         m.Status,
		Extra:          m.Extra,
	}
}
