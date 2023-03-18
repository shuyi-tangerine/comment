include "base.thrift"

namespace go tangerine.comment

enum CommentStatus {
    Normal = 1;
    Deleted = 2;
}

struct GenCommentIDRequest {
    1: optional i8 amount   // 需要的数量，默认返回 1 个
    255: optional base.RPCRequest Base
}

struct GenCommentIDResponse {
    1: required list<i64> comment_ids
    255: required base.RPCResponse Base
}

struct PostRequest {
    1: optional i64 comment_id  // null、0 的时候系统会生成一个，一般建议调用接口生成
    2: required i64 user_id
    3: required i64 group_id
    4: required i32 app_id
    5: required string text
    6: optional list<string> images
    7: optional i64 reply_comment_id
    8: optional string extra  // 要求 JSON string
    255: optional base.RPCRequest Base
}

struct PostResponse {
    1: optional CommentData comment_data
    255: required base.RPCResponse Base
}

struct ListRequest {
    1: optional i64 comment_id
    2: optional i64 user_id
    3: optional i64 group_id
    4: optional i32 app_id
    5: optional TimeRange created_at
    6: required i64 offset
    7: required i32 limit
    255: optional base.RPCRequest Base
}

struct TimeRange {
    1: required i64 s  // start
    2: required i64 e  // end
}

struct ListResponse {
    1: required i64 offset
    2: required i32 total
    3: optional list<CommentData> comments_data
    255: required base.RPCResponse Base
}

struct CommentData {
    1: required i64 id
    2: required i64 comment_id
    3: required i64 user_id
    4: required i64 group_id
    5: required i32 app_id
    6: required string text
    7: optional list<string> images
    8: optional i64 reply_comment_id
    9: required i64 created_at
    10: required i64 updated_at
    11: optional string extra  // JSON string
}

struct DeleteRequest {
    1: required i64 comment_id
    255: optional base.RPCRequest Base
}

struct DeleteResponse {
    255: required base.RPCResponse Base
}

service CommentHandler {
    // 生成 comment_id
    GenCommentIDResponse GenCommentID(1:GenCommentIDRequest req)
    // 发送评论
    PostResponse Post(1:PostRequest req)
    // 获取评论列表
    ListResponse List(1:ListRequest req)
    // 删除评论
    DeleteResponse Delete(1:DeleteRequest req)
}