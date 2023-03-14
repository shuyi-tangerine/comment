package top

type CommentStorage interface {
	// InsertOne 插入一条数据，并返回存储中的数据
	InsertOne(data *Comment) (id int64, err error)
	// SelectOne 只返回一条数据，一般用于主键、唯一键查询
	SelectOne(req *OperateCommentRequest) (storageComment *Comment, err error)
	// Select 返回列表
	Select(req *OperateCommentRequest) (storageComments []*Comment, err error)
	// UpdateOne 更新，按照主键、唯一键更新
	UpdateOne(req *OperateCommentRequest) (err error)
	// DeleteOne 删除，一般为逻辑删除,按照主键、唯一键删除
	DeleteOne(req *OperateCommentRequest) (err error)
}

type OperateCommentRequest struct {
	// TODO
}
