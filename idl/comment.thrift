include "base.thrift"

namespace go tangerine.comment

struct GenCommentIDRequest {
    1: optional i8 amount = 1   // 需要的数量，默认返回 1 个
    255: optional base.RPCRequest Base
}

struct GenCommentIDResponse {
    2: required list<i64> comment_ids
    255: required base.RPCResponse Base
}

struct PostRequest {
    1: required i32 app_id
    2: required i64 user_id
    3: required string text
    4: optional list<images> images
    5: optional int64 comment_id = 0  // null、0 的时候系统会生成一个，一般建议调用接口生成
    255: optional base.RPCRequest Base
}

struct PostResponse {
    255: required base.RPCResponse Base
}

struct GetRequest {
    255: optional base.RPCRequest Base
}

struct GetResponse {
    255: required base.RPCResponse Base
}

struct DeleteRequest {
    255: optional base.RPCRequest Base
}

struct DeleteResponse {
    255: required base.RPCResponse Base
}

service CommentService {
    // 生成 comment_id
    GenCommentIDResponse GenCommentID(1:GenCommentIDRequest req)
    // 发送评论
    PostResponse Post(1:PostRequest req)
    // 获取评论
    GetResponse Get(1:GetRequest req)
    // 删除评论
    DeleteResponse Delete(1:DeleteRequest req)
}