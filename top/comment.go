package top

import "time"

// Comment  评论
type Comment struct {
	ID        int64     `json:"id"`         //  自增主键
	CommentID int64     `json:"comment_id"` //  评论id
	UserID    int64     `json:"user_id"`    //  用户id
	GroupID   int64     `json:"group-id"`   //  所属分组ID，比如视频、文章等
	AppID     int32     `json:"app_id"`     //  应用id
	Text      string    `json:"text"`       //  文本
	Images    string    `json:"images"`     //  图片
	CreatedAt time.Time `json:"created_at"` //  创建时间，也就是通知发送时间
	UpdatedAt time.Time `json:"updated_at"` //  最后更新时间
	Status    int8      `json:"status"`     //  状态，1-正常，2-已删除
	Extra     string    `json:"extra"`      //  一些额外的信息
}
